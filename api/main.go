package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/middlewares"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/repositories/localstorage"
	"github.com/JenswBE/go-commerce/repositories/productpg"
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
	config, err := parseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}

	// Setup Debug logging if enabled
	if config.Server.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		gin.SetMode(gin.DebugMode)
		log.Debug().Msg("Debug logging enabled")
	}

	// DB
	dsn := buildDSN(config)
	productDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to DB")
	}

	// Services
	productDatabase, err := productpg.NewProductPostgres(productDB)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate DB and create products repository")
	}
	imageStorage, err := getStorageRepo(config.Storage.Images)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create image storage repository")
	}
	allowedImageConfigs, err := parseAllowedImageConfigs(config.ImageProxy.AllowedConfigs)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse allowed image configs")
	}
	imageProxyService, err := imageproxy.NewImgProxyService(config.ImageProxy.BaseURL, config.ImageProxy.Key, config.ImageProxy.Salt, allowedImageConfigs)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create image proxy service")
	}
	productService := product.NewService(productDatabase, imageProxyService, imageStorage)
	shortIDService := shortid.NewBase58Service()
	presenter := presenter.New(shortIDService)

	// Setup Gin
	router := gin.Default()
	router.StaticFile("/", "../docs/index.html")
	router.StaticFile("/openapi.yml", "../docs/openapi.yml")

	// Setup handlers
	productHandler := handler.NewProductHandler(presenter, productService)

	// Public routes
	public := router.Group("/public")
	productHandler.RegisterReadRoutes(public)

	// Admin routes
	authMW, err := middlewares.NewOIDCMiddleware(config.Authentication.IssuerURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create OIDC middleware")
	}
	admin := router.Group("/admin")
	admin.Use(authMW.Handle)
	productHandler.RegisterReadRoutes(admin)
	productHandler.RegisterWriteRoutes(admin)

	// Start Gin
	port := strconv.Itoa(config.Server.Port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal().Err(err).Int("port", config.Server.Port).Msg("Failed to start Gin server")
	}
}

func getStorageRepo(config Storage) (product.StorageRepository, error) {
	switch config.Type {
	case StorageTypeFS:
		return localstorage.NewLocalStorage(config.Path)
	default:
		return nil, fmt.Errorf(`unknown storage type "%s"`, config.Type)
	}
}
