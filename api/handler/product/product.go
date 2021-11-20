package product

import (
	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/openapi"
	presenter "github.com/JenswBE/go-commerce/api/presenter/product"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *ProductHandler) listProducts(c *gin.Context) {
	// Parse params
	imageConfigs, err := parseImageConfigsParam(c)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	result, err := h.service.ListProducts(imageConfigs)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ProductListFromEntity(h.presenter, result))
}

func (h *ProductHandler) getProduct(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageConfigs, err := parseImageConfigsParam(c)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}
	resolve := c.Query("resolve") == "true"

	// Call service
	product, err := h.service.GetProduct(id, resolve, imageConfigs)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Convert to OpenAPI model
	output, err := presenter.ResolvedProductFromEntity(h.presenter, product)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, output)
}

func (h *ProductHandler) createProduct(c *gin.Context) {
	// Parse body
	body := &openapi.Product{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Convert to entity
	e, err := presenter.ProductToEntity(h.presenter, uuid.Nil, *body)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	product, err := h.service.CreateProduct(e)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(201, presenter.ProductFromEntity(h.presenter, product))
}

func (h *ProductHandler) updateProduct(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Parse body
	body := &openapi.Product{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Convert to entity
	e, err := presenter.ProductToEntity(h.presenter, id, *body)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	product, err := h.service.UpdateProduct(e)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ProductFromEntity(h.presenter, product))
}

func (h *ProductHandler) deleteProduct(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteProduct(id)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}

func (h *ProductHandler) listProductImages(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageConfigs, err := parseImageConfigsParam(c)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	product, err := h.service.GetProduct(id, false, imageConfigs)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ImageListFromEntity(h.presenter, product.Images))
}

func (h *ProductHandler) addProductImages(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageConfigs, err := parseImageConfigsParam(c)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Parse body
	images, err := parseFilesFromMultipart(c.Request)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	product, err := h.service.AddProductImages(id, images, imageConfigs)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ImageListFromEntity(h.presenter, product.Images))
}

func (h *ProductHandler) updateProductImage(c *gin.Context) {
	// Parse params
	productID, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageID, ok := handler.ParseIDParam(c, "image_id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Parse body
	body := &openapi.Image{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	images, err := h.service.UpdateProductImage(productID, imageID, int(body.Order))
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ImageSliceFromEntity(h.presenter, images))
}

func (h *ProductHandler) deleteProductImage(c *gin.Context) {
	// Parse params
	productID, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageID, ok := handler.ParseIDParam(c, "image_id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteProductImage(productID, imageID)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}
