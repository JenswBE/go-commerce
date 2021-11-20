package product

import (
	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/openapi"
	presenter "github.com/JenswBE/go-commerce/api/presenter/product"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *ProductHandler) createManufacturer(c *gin.Context) {
	// Parse body
	body := &openapi.Manufacturer{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	e := presenter.ManufacturerToEntity(h.presenter, uuid.Nil, *body)
	manufacturer, err := h.service.CreateManufacturer(e)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(201, presenter.ManufacturerFromEntity(h.presenter, manufacturer))
}

func (h *ProductHandler) updateManufacturer(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Parse body
	body := &openapi.Manufacturer{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	e := presenter.ManufacturerToEntity(h.presenter, id, *body)
	manufacturer, err := h.service.UpdateManufacturer(e)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ManufacturerFromEntity(h.presenter, manufacturer))
}

func (h *ProductHandler) deleteManufacturer(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteManufacturer(id)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}

func (h *ProductHandler) upsertManufacturerImage(c *gin.Context) {
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
	if len(images) != 1 {
		err := entities.NewError(400, openapi.GOCOMERRORCODE_SINGLE_IMAGE_IN_FORM, "", nil)
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	var manufacturer *entities.Manufacturer
	for imageName, imageContent := range images {
		manufacturer, err = h.service.UpsertManufacturerImage(id, imageName, imageContent, imageConfigs)
		if err != nil {
			c.JSON(handler.ErrToResponse(err))
			return
		}
	}

	// Handle success
	c.JSON(200, presenter.ImageFromEntity(h.presenter, manufacturer.Image))
}

func (h *ProductHandler) deleteManufacturerImage(c *gin.Context) {
	// Parse params
	manufacturerID, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteManufacturerImage(manufacturerID)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}
