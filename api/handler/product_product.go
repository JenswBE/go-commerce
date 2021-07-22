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
		// Parse params
		imageConfig, err := parseImageConfigParams(c)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		result, err := service.ListProducts(imageConfig)
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
		id, ok := parseIDParam(c, "id", p)
		if !ok {
			return // Response already set on Gin context
		}
		imageConfig, err := parseImageConfigParams(c)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		product, err := service.GetProduct(id, imageConfig)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ProductFromEntity(product))
	}
}

func listProductImages(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseIDParam(c, "id", p)
		if !ok {
			return // Response already set on Gin context
		}
		imageConfig, err := parseImageConfigParams(c)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		product, err := service.GetProduct(id, imageConfig)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ImagesListFromEntity(product.Images))
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
		id, ok := parseIDParam(c, "id", p)
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
		id, ok := parseIDParam(c, "id", p)
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
		id, ok := parseIDParam(c, "id", p)
		if !ok {
			return // Response already set on Gin context
		}
		imageConfig, err := parseImageConfigParams(c)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Parse body
		images, err := parseFilesFromMultipart(c.Request)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		product, err := service.AddProductImages(id, images, imageConfig)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ImagesListFromEntity(product.Images))
	}
}

func updateProductImage(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		productID, ok := parseIDParam(c, "id", p)
		if !ok {
			return // Response already set on Gin context
		}
		imageID, ok := parseIDParam(c, "image_id", p)
		if !ok {
			return // Response already set on Gin context
		}

		// Parse body
		body := &openapi.Image{}
		if err := c.BindJSON(body); err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		images, err := service.UpdateProductImage(productID, imageID, int(body.Order))
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ImagesListFromEntity(images))
	}
}

func deleteProductImage(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		productID, ok := parseIDParam(c, "id", p)
		if !ok {
			return // Response already set on Gin context
		}
		imageID, ok := parseIDParam(c, "image_id", p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		err := service.DeleteProductImage(productID, imageID)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.String(204, "")
	}
}
