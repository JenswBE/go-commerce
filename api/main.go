package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/repositories/localstorage"
	"github.com/JenswBE/go-commerce/repositories/productpg"
	"github.com/JenswBE/go-commerce/usecases/product"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/JenswBE/go-commerce/utils/shortid"
	"github.com/gin-contrib/cors"
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

	// Parse config
	config, err := parseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}

	// DB
	dsn := buildDSN(config)
	productDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to db")
	}

	// Services
	productDatabase := productpg.NewProductPostgres(productDB)
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
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "authorization")
	router.Use(cors.New(corsConfig))
	router.StaticFile("/", "../docs/index.html")
	router.StaticFile("/openapi.yml", "../docs/openapi.yml")

	// Public routes
	public := router.Group("/public")
	handler.AddProductReadRoutes(public, presenter, productService)

	// Admin routes
	admin := router.Group("/admin")
	handler.AddProductReadRoutes(admin, presenter, productService)
	handler.AddProductWriteRoutes(admin, presenter, productService)

	// Start Gin
	port := strconv.Itoa(config.Server.Port)
	router.Run(":" + port)
}

func getStorageRepo(config Storage) (product.StorageRepository, error) {
	switch config.Type {
	case StorageTypeFS:
		return localstorage.NewLocalStorage(config.Path)
	default:
		return nil, fmt.Errorf(`unknown storage type "%s"`, config.Type)
	}
}
