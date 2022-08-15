package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/JenswBE/go-commerce/admin"
	"github.com/JenswBE/go-commerce/admin/auth"
	contentHandler "github.com/JenswBE/go-commerce/api/handler/content"
	productHandler "github.com/JenswBE/go-commerce/api/handler/product"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/config"
	"github.com/JenswBE/go-commerce/repositories/contentpg"
	"github.com/JenswBE/go-commerce/repositories/localstorage"
	"github.com/JenswBE/go-commerce/repositories/productpg"
	"github.com/JenswBE/go-commerce/usecases/content"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/JenswBE/go-commerce/utils/logging"
	"github.com/JenswBE/go-commerce/utils/sanitizer"
	"github.com/JenswBE/go-commerce/utils/shortid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"golang.org/x/exp/maps"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Setup logging
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = log.Output(output)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	gin.SetMode(gin.ReleaseMode)

	// Parse config
	svcConfig, err := config.ParseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}

	// Setup Debug logging if enabled
	if svcConfig.Server.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Logger.Hook(logging.CallerInfoHook{})
		gin.SetMode(gin.DebugMode)
		log.Debug().Msg("Debug logging enabled")
	}

	// DB
	contentDSN := config.BuildDSN(svcConfig.Database.Content, svcConfig.Database.Default)
	contentDB, err := gorm.Open(postgres.Open(contentDSN), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to content DB")
	}
	productDSN := config.BuildDSN(svcConfig.Database.Product, svcConfig.Database.Default)
	productDB, err := gorm.Open(postgres.Open(productDSN), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to product DB")
	}

	// Services
	contentDatabase, err := contentpg.NewContentPostgres(contentDB)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate DB and create content repository")
	}
	productDatabase, err := productpg.NewProductPostgres(productDB)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate DB and create products repository")
	}
	imageStorage, err := getStorageRepo(svcConfig.Storage.Images)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create image storage repository")
	}
	allowedImageConfigs, err := config.ParseAllowedImageConfigs(svcConfig.ImageProxy.AllowedConfigs)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse allowed image configs")
	}
	allowedImageConfigs = append(allowedImageConfigs, maps.Values(admin.ProductImageConfigs)...)
	imageProxyService, err := imageproxy.NewImgProxyService(svcConfig.ImageProxy.BaseURL, svcConfig.ImageProxy.Key, svcConfig.ImageProxy.Salt, allowedImageConfigs)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create image proxy service")
	}
	contentService, err := content.NewService(contentDatabase, svcConfig.Features.Content.List.ToEntity())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create content service")
	}
	productService := product.NewService(productDatabase, imageProxyService, imageStorage)
	shortIDService := shortid.NewBase58Service()
	sanitizerService := sanitizer.NewBluemondayService()
	presenter := presenter.New(shortIDService, sanitizerService)

	// Setup Gin
	router := gin.Default()
	router.RedirectTrailingSlash = true
	err = router.SetTrustedProxies(svcConfig.Server.TrustedProxies)
	if err != nil {
		log.Fatal().Err(err).Strs("trusted_proxies", svcConfig.Server.TrustedProxies).Msg("Failed to set trusted proxies")
	}
	router.GET("/", func(c *gin.Context) { c.Redirect(http.StatusTemporaryRedirect, admin.PrefixAdmin) })

	// Setup API routes
	apiGroup := router.Group("/api")
	apiGroup.StaticFile("", "docs/index.html")
	apiGroup.StaticFile("/", "docs/index.html")
	apiGroup.StaticFile("/index.html", "docs/index.html")
	apiGroup.StaticFile("/openapi.yml", "docs/openapi.yml")
	contentHandler.NewContentHandler(presenter, contentService).RegisterRoutes(apiGroup)
	productHandler.NewProductHandler(presenter, productService).RegisterRoutes(apiGroup)

	// Setup admin authentication
	var oidcClient *auth.OIDCClient
	switch svcConfig.Authentication.Type {
	case config.AuthTypeNone:
		log.Warn().Msg("Authentication disabled because of Authentication type NONE. This should only be used for testing or with an external service to enforce authentication!")
	case config.AuthTypeOIDC:
		oidcClient, err = auth.NewOIDCClient(svcConfig.Authentication.OIDC.IssuerURL, svcConfig.Authentication.OIDC.ClientID, svcConfig.Authentication.OIDC.ClientSecret)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create OIDC verifier")
		}
	default:
		log.Fatal().Msg("Invalid authentication type specified") // Should be captured by exhaustive, but too important to not shield.
	}

	// Setup admin GUI
	adminHandler := admin.NewHandler(svcConfig.Features, productService, contentService, oidcClient, svcConfig.Authentication.SessionAuthKey, svcConfig.Authentication.SessionEncKey)
	adminHandler.RegisterRoutes(router)
	router.HTMLRender = adminHandler.NewRenderer()

	// Start Gin
	port := strconv.Itoa(svcConfig.Server.Port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal().Err(err).Int("port", svcConfig.Server.Port).Msg("Failed to start Gin server")
	}
}

func getStorageRepo(svcConfig config.Storage) (product.StorageRepository, error) {
	switch svcConfig.Type {
	case config.StorageTypeFS:
		return localstorage.NewLocalStorage(svcConfig.Path)
	default:
		return nil, fmt.Errorf(`unknown storage type "%s"`, svcConfig.Type)
	}
}
