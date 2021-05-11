package main

import (
	"log"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/infrastructure/repository/productpg"
	"github.com/JenswBE/go-commerce/pkg/shortid"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB
	dsn := "host=localhost user=go_commerce password=go_commerce dbname=go_commerce port=5432"
	productDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err.Error())
	}

	// Services
	productRepo := productpg.NewProductPostgres(productDB)
	productService := product.NewService(productRepo)
	shortIDService := shortid.NewBase58Service()
	presenter := presenter.New(shortIDService)

	// Setup Gin
	router := gin.Default()
	router.Use(cors.Default())
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
	router.Run(":8090")
}
