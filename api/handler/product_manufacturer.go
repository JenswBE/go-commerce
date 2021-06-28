package handler

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-gonic/gin"
)

func listManufacturers(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call service
		result, err := service.ListManufacturers()
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
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		manufacturer, err := service.GetManufacturer(id)
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
		manufacturer, err := service.CreateManufacturer(body.GetName(), body.GetWebsiteUrl())
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
		id, ok := parseParamID(c, p)
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
		id, ok := parseParamID(c, p)
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
