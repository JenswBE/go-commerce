package product

import (
	"github.com/JenswBE/go-commerce/api/handler"
	presenter "github.com/JenswBE/go-commerce/api/presenter/product"
	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) listCategories(c *gin.Context) {
	// Parse params
	imageConfigs, err := parseImageConfigsParam(c)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	result, err := h.service.ListCategories(imageConfigs)

	// Handle errors
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
	}

	// Handle success
	c.JSON(200, presenter.CategoryListFromEntity(h.presenter, result))
}

func (h *ProductHandler) getCategory(c *gin.Context) {
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
	category, err := h.service.GetCategory(id, imageConfigs)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.CategoryFromEntity(h.presenter, category))
}
