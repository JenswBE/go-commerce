package main

import (
	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/pkg/shortid"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Services
	productRepo := product.NewInmem()
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
