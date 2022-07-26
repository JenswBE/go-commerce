package product

import (
	"github.com/JenswBE/go-commerce/api/handler"
	presenter "github.com/JenswBE/go-commerce/api/presenter/product"
	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) listManufacturers(c *gin.Context) {
	// Parse params
	imageConfigs, err := parseImageConfigsParam(c)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	result, err := h.service.ListManufacturers(imageConfigs)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ManufacturerListFromEntity(h.presenter, result))
}

func (h *ProductHandler) getManufacturer(c *gin.Context) {
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
	manufacturer, err := h.service.GetManufacturer(id, imageConfigs)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ManufacturerFromEntity(h.presenter, manufacturer))
}
