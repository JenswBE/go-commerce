package handler

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func listProducts(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call service
		result, err := service.ListProducts()
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ProductsListFromEntity(result))
	}
}

func getProduct(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		product, err := service.GetProduct(id)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ProductFromEntity(product))
	}
}

func createProduct(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse body
		body := &openapi.Product{}
		if err := c.BindJSON(body); err != nil {
			c.String(errToResponse(err))
			return
		}

		// Convert to entity
		e, err := p.ProductToEntity(uuid.Nil, *body)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		product, err := service.CreateProduct(e)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ProductFromEntity(product))
	}
}

func updateProduct(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Parse body
		body := &openapi.Product{}
		if err := c.BindJSON(body); err != nil {
			c.String(errToResponse(err))
			return
		}

		// Convert to entity
		e, err := p.ProductToEntity(id, *body)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		product, err := service.UpdateProduct(e)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ProductFromEntity(product))
	}
}

func deleteProduct(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		err := service.DeleteProduct(id)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.String(204, "")
	}
}

func addProductImages(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Parse body
		images, err := parseFilesFromMultipart(c.Request)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		product, err := service.AddProductImages(id, images)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ImagesListFromEntity(product.Images))
	}
}
