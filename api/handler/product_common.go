package handler

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-gonic/gin"
)

func AddProductReadRoutes(rg *gin.RouterGroup, p *presenter.Presenter, service product.Usecase) {
	groupCategories := rg.Group("/categories")
	groupCategories.GET("", listCategories(p, service))
	groupCategories.GET("/:id", getCategory(p, service))

	groupManufacturers := rg.Group("/manufacturers")
	groupManufacturers.GET("", listManufacturers(p, service))
	groupManufacturers.GET("/:id", getManufacturer(p, service))

	groupProducts := rg.Group("/products")
	groupProducts.GET("", listProducts(p, service))
	groupProducts.GET("/:id", getProduct(p, service))
}

func AddProductWriteRoutes(rg *gin.RouterGroup, p *presenter.Presenter, service product.Usecase) {
	groupCategories := rg.Group("/categories")
	groupCategories.POST("", createCategory(p, service))
	groupCategories.PUT("/:id", updateCategory(p, service))
	groupCategories.DELETE("/:id", deleteCategory(p, service))

	groupManufacturers := rg.Group("/manufacturers")
	groupManufacturers.POST("", createManufacturer(p, service))
	groupManufacturers.PUT("/:id", updateManufacturer(p, service))
	groupManufacturers.DELETE("/:id", deleteManufacturer(p, service))

	groupProducts := rg.Group("/products")
	groupProducts.POST("", createProduct(p, service))
	groupProducts.PUT("/:id", updateProduct(p, service))
	groupProducts.DELETE("/:id", deleteProduct(p, service))
}
