package product

import (
	"github.com/JenswBE/go-commerce/api/handler"
	presenter "github.com/JenswBE/go-commerce/api/presenter/product"
	"github.com/gin-gonic/gin"
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
