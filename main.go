package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/JenswBE/go-commerce/admin"
	configHandler "github.com/JenswBE/go-commerce/api/handler/config"
	contentHandler "github.com/JenswBE/go-commerce/api/handler/content"
	productHandler "github.com/JenswBE/go-commerce/api/handler/product"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/config"
	"github.com/JenswBE/go-commerce/repositories/contentpg"
	"github.com/JenswBE/go-commerce/repositories/localstorage"
	"github.com/JenswBE/go-commerce/repositories/productpg"
	"github.com/JenswBE/go-commerce/usecases/content"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/JenswBE/go-commerce/utils/auth"
	"github.com/JenswBE/go-commerce/utils/generics"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/JenswBE/go-commerce/utils/sanitizer"
	"github.com/JenswBE/go-commerce/utils/shortid"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Setup logging
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

	// Setup API routes
	apiGroup := router.Group("/api")
	apiGroup.StaticFile("", "docs/index.html")
	apiGroup.StaticFile("/", "docs/index.html")
	apiGroup.StaticFile("/index.html", "docs/index.html")
	apiGroup.StaticFile("/oauth2-redirect.html", "docs/oauth2-redirect.html")
	apiGroup.StaticFile("/openapi.yml", "docs/openapi.yml")

	// Setup handlers
	configHandler := configHandler.NewConfigHandler(presenter, *svcConfig)
	contentHandler := contentHandler.NewContentHandler(presenter, contentService)
	productHandler := productHandler.NewProductHandler(presenter, productService)

	// API public routes
	apiPublic := apiGroup.Group("/")
	configHandler.RegisterPublicRoutes(apiPublic)
	contentHandler.RegisterPublicRoutes(apiPublic)
	productHandler.RegisterPublicRoutes(apiPublic)

	// Setup admin authentication
	var authVerifier auth.Verifier
	switch svcConfig.Authentication.Type {
	case config.AuthTypeBasicAuth:
		log.Warn().Msg("Using authentication type BASIC_AUTH. This should only be used for E2E testing!")
		authVerifier = auth.NewBasicVerifier(svcConfig.Authentication.BasicAuth.Username, svcConfig.Authentication.BasicAuth.Password)
	case config.AuthTypeOIDC:
		authVerifier, err = auth.NewOIDCVerifier(svcConfig.Authentication.OIDC.IssuerURL, svcConfig.Authentication.OIDC.ClientID)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create OIDC middleware")
		}
	default:
		log.Fatal().Msg("Invalid authentication type specified") // Should be captured by exhaustive, but too important to not shield.
	}

	// API admin routes
	apiAdmin := apiGroup.Group("/")
	apiAdmin.Use(auth.NewAuthMiddleware(authVerifier).EnforceRoles([]string{auth.RoleAdmin}))
	contentHandler.RegisterAdminRoutes(apiAdmin)
	productHandler.RegisterAdminRoutes(apiAdmin)

	// Setup admin GUI
	router.HTMLRender = createAdminGUIRenderer()
	adminHandler, err := admin.NewAdminGUIHandler(productService, svcConfig.Server.SessionAuthKey, authVerifier, svcConfig.Server.JWTSigningKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to register admin handler")
	}
	adminHandler.RegisterRoutes(router)

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

func createAdminGUIRenderer() multitemplate.Renderer {
	pages := map[string][]string{
		"categoriesList":    {"pages/categories_list"},
		"login":             {"pages/login"},
		"manufacturersList": {"pages/manufacturers_list"},
		"productsList":      {"pages/products_list"},
	}

	r := multitemplate.NewRenderer()
	for pageName, templates := range pages {
		templates = append([]string{"layouts/empty", "layouts/base"}, templates...)
		templatePaths := generics.Map(templates, func(i string) string { return fmt.Sprintf("admin/html/%s.html.go.tmpl", i) })
		r.AddFromFiles(pageName, templatePaths...)
	}
	return r
}
