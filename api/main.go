package main

import (
	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Services
	productRepo := product.NewInmem()
	productService := product.NewService(productRepo)

	// Setup Gin
	router := gin.Default()
	router.Use(cors.Default())
	router.StaticFile("/", "../docs/index.html")
	router.StaticFile("/openapi.yml", "../docs/openapi.yml")

	// Public routes
	public := router.Group("/public")
	handler.AddProductReadRoutes(public, productService)

	// Admin routes
	admin := router.Group("/admin")
	handler.AddProductReadRoutes(admin, productService)
	handler.AddProductWriteRoutes(admin, productService)

	// Start Gin
	router.Run(":8090")
}
