package handler

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *ProductHandler) listProducts(c *gin.Context) {
	// Parse params
	imageConfig, err := parseImageConfigParams(c)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Call service
	result, err := h.service.ListProducts(imageConfig)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ProductsListFromEntity(result))
}

func (h *ProductHandler) getProduct(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageConfig, err := parseImageConfigParams(c)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Call service
	product, err := h.service.GetProduct(id, imageConfig)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ProductFromEntity(product))
}

func (h *ProductHandler) listProductImages(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageConfig, err := parseImageConfigParams(c)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Call service
	product, err := h.service.GetProduct(id, imageConfig)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ImagesListFromEntity(product.Images))
}

func (h *ProductHandler) createProduct(c *gin.Context) {
	// Parse body
	body := &openapi.Product{}
	if err := c.BindJSON(body); err != nil {
		c.String(errToResponse(err))
		return
	}

	// Convert to entity
	e, err := h.presenter.ProductToEntity(uuid.Nil, *body)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Call service
	product, err := h.service.CreateProduct(e)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ProductFromEntity(product))
}

func (h *ProductHandler) updateProduct(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
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
	e, err := h.presenter.ProductToEntity(id, *body)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Call service
	product, err := h.service.UpdateProduct(e)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ProductFromEntity(product))
}

func (h *ProductHandler) deleteProduct(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteProduct(id)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}

func (h *ProductHandler) addProductImages(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
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
	product, err := h.service.AddProductImages(id, images, imageConfig)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ImagesListFromEntity(product.Images))
}

func (h *ProductHandler) updateProductImage(c *gin.Context) {
	// Parse params
	productID, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageID, ok := parseIDParam(c, "image_id", h.presenter)
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
	images, err := h.service.UpdateProductImage(productID, imageID, int(body.Order))
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ImagesListFromEntity(images))
}

func (h *ProductHandler) deleteProductImage(c *gin.Context) {
	// Parse params
	productID, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageID, ok := parseIDParam(c, "image_id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteProductImage(productID, imageID)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}
