package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/JenswBE/go-commerce/api/config"
	configHandler "github.com/JenswBE/go-commerce/api/handler/config"
	contentHandler "github.com/JenswBE/go-commerce/api/handler/content"
	productHandler "github.com/JenswBE/go-commerce/api/handler/product"
	"github.com/JenswBE/go-commerce/api/middlewares"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/repositories/contentpg"
	"github.com/JenswBE/go-commerce/repositories/localstorage"
	"github.com/JenswBE/go-commerce/repositories/productpg"
	"github.com/JenswBE/go-commerce/usecases/content"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/JenswBE/go-commerce/utils/shortid"
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
	apiConfig, err := config.ParseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}

	// Setup Debug logging if enabled
	if apiConfig.Server.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		gin.SetMode(gin.DebugMode)
		log.Debug().Msg("Debug logging enabled")
	}

	// DB
	contentDSN := config.BuildDSN(apiConfig.Database.Content, apiConfig.Database.Default)
	contentDB, err := gorm.Open(postgres.Open(contentDSN), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to content DB")
	}
	productDSN := config.BuildDSN(apiConfig.Database.Product, apiConfig.Database.Default)
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
	imageStorage, err := getStorageRepo(apiConfig.Storage.Images)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create image storage repository")
	}
	allowedImageConfigs, err := config.ParseAllowedImageConfigs(apiConfig.ImageProxy.AllowedConfigs)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse allowed image configs")
	}
	imageProxyService, err := imageproxy.NewImgProxyService(apiConfig.ImageProxy.BaseURL, apiConfig.ImageProxy.Key, apiConfig.ImageProxy.Salt, allowedImageConfigs)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create image proxy service")
	}
	contentService, err := content.NewService(contentDatabase, apiConfig.Features.Content.List.ToEntity())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create content service")
	}
	productService := product.NewService(productDatabase, imageProxyService, imageStorage)
	shortIDService := shortid.NewBase58Service()
	presenter := presenter.New(shortIDService)

	// Setup Gin
	router := gin.Default()
	err = router.SetTrustedProxies(apiConfig.Server.TrustedProxies)
	if err != nil {
		log.Fatal().Err(err).Strs("trusted_proxies", apiConfig.Server.TrustedProxies).Msg("Failed to set trusted proxies")
	}
	router.StaticFile("/", "../docs/index.html")
	router.StaticFile("/index.html", "../docs/index.html")
	router.StaticFile("/oauth2-redirect.html", "../docs/oauth2-redirect.html")
	router.StaticFile("/openapi.yml", "../docs/openapi.yml")

	// Setup handlers
	configHandler := configHandler.NewConfigHandler(presenter, *apiConfig)
	contentHandler := contentHandler.NewContentHandler(presenter, contentService)
	productHandler := productHandler.NewProductHandler(presenter, productService)

	// Public routes
	public := router.Group("/")
	configHandler.RegisterPublicRoutes(public)
	contentHandler.RegisterPublicRoutes(public)
	productHandler.RegisterPublicRoutes(public)

	// Admin routes
	authMW, err := middlewares.NewOIDCMiddleware(apiConfig.Authentication.IssuerURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create OIDC middleware")
	}
	admin := router.Group("/")
	admin.Use(authMW.EnforceRoles([]string{"admin"}))
	contentHandler.RegisterAdminRoutes(admin)
	productHandler.RegisterAdminRoutes(admin)

	// Start Gin
	port := strconv.Itoa(apiConfig.Server.Port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal().Err(err).Int("port", apiConfig.Server.Port).Msg("Failed to start Gin server")
	}
}

func getStorageRepo(apiConfig config.Storage) (product.StorageRepository, error) {
	switch apiConfig.Type {
	case config.StorageTypeFS:
		return localstorage.NewLocalStorage(apiConfig.Path)
	default:
		return nil, fmt.Errorf(`unknown storage type "%s"`, apiConfig.Type)
	}
}
