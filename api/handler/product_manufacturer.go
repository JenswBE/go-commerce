package handler

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *ProductHandler) listManufacturers(c *gin.Context) {
	// Parse params
	imageConfigs, err := parseImageConfigsParam(c)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Call service
	result, err := h.service.ListManufacturers(imageConfigs)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ManufacturerListFromEntity(result))
}

func (h *ProductHandler) getManufacturer(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageConfigs, err := parseImageConfigsParam(c)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Call service
	manufacturer, err := h.service.GetManufacturer(id, imageConfigs)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ManufacturerFromEntity(manufacturer))
}

func (h *ProductHandler) createManufacturer(c *gin.Context) {
	// Parse body
	body := &openapi.Manufacturer{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Call service
	e := h.presenter.ManufacturerToEntity(uuid.Nil, *body)
	manufacturer, err := h.service.CreateManufacturer(e)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(201, h.presenter.ManufacturerFromEntity(manufacturer))
}

func (h *ProductHandler) updateManufacturer(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Parse body
	body := &openapi.Manufacturer{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Call service
	e := h.presenter.ManufacturerToEntity(id, *body)
	manufacturer, err := h.service.UpdateManufacturer(e)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.ManufacturerFromEntity(manufacturer))
}

func (h *ProductHandler) deleteManufacturer(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteManufacturer(id)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}

func (h *ProductHandler) upsertManufacturerImage(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageConfigs, err := parseImageConfigsParam(c)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Parse body
	images, err := parseFilesFromMultipart(c.Request)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}
	if len(images) != 1 {
		err := entities.NewError(400, openapi.ERRORCODE_SINGLE_IMAGE_IN_FORM, "", nil)
		c.JSON(errToResponse(err))
		return
	}

	// Call service
	var manufacturer *entities.Manufacturer
	for imageName, imageContent := range images {
		manufacturer, err = h.service.UpsertManufacturerImage(id, imageName, imageContent, imageConfigs)
		if err != nil {
			c.JSON(errToResponse(err))
			return
		}
	}

	// Handle success
	c.JSON(200, h.presenter.ImageFromEntity(manufacturer.Image))
}

func (h *ProductHandler) deleteManufacturerImage(c *gin.Context) {
	// Parse params
	manufacturerID, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteManufacturerImage(manufacturerID)
	if err != nil {
		c.JSON(errToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}
