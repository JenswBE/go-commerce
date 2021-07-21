package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/repository/localstorage"
	"github.com/JenswBE/go-commerce/repository/productpg"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/JenswBE/go-commerce/utils/shortid"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Parse config
	config, err := parseConfig()
	if err != nil {
		log.Fatalf("Failed to parse config: %s", err)
	}

	// DB
	dsn := buildDSN(config)
	productDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err.Error())
	}

	// Services
	productDatabase := productpg.NewProductPostgres(productDB)
	imageStorage, err := getStorageRepo(config.Storage.Images)
	if err != nil {
		log.Fatalf("Failed to create image storage repository: %s", err.Error())
	}
	productService := product.NewService(productDatabase, imageStorage)
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
