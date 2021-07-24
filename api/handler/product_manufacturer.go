package handler

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func listManufacturers(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		imageConfig, err := parseImageConfigParams(c)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		result, err := service.ListManufacturers(imageConfig)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ManufacturersListFromEntity(result))
	}
}

func getManufacturer(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
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
		manufacturer, err := service.GetManufacturer(id, imageConfig)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ManufacturerFromEntity(manufacturer))
	}
}

func createManufacturer(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse body
		body := &openapi.Manufacturer{}
		if err := c.BindJSON(body); err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		e := p.ManufacturerToEntity(uuid.Nil, *body)
		manufacturer, err := service.CreateManufacturer(e)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ManufacturerFromEntity(manufacturer))
	}
}

func updateManufacturer(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseIDParam(c, "id", p)
		if !ok {
			return // Response already set on Gin context
		}

		// Parse body
		body := &openapi.Manufacturer{}
		if err := c.BindJSON(body); err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		e := p.ManufacturerToEntity(id, *body)
		manufacturer, err := service.UpdateManufacturer(e)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.ManufacturerFromEntity(manufacturer))
	}
}

func deleteManufacturer(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseIDParam(c, "id", p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		err := service.DeleteManufacturer(id)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.String(204, "")
	}
}

func upsertManufacturerImage(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
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
		if len(images) != 1 {
			c.String(400, "expects exactly 1 image in multipart form")
			return
		}

		// Call service
		var manufacturer *entity.Manufacturer
		for imageName, imageContent := range images {
			manufacturer, err = service.UpsertManufacturerImage(id, imageName, imageContent, imageConfig)
			if err != nil {
				c.String(errToResponse(err))
				return
			}
		}

		// Handle success
		c.JSON(200, p.ImageFromEntity(manufacturer.Image))
	}
}

func deleteManufacturerImage(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		manufacturerID, ok := parseIDParam(c, "id", p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		err := service.DeleteManufacturerImage(manufacturerID)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.String(204, "")
	}
}
