package handler

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-gonic/gin"
)

func AddProductReadRoutes(rg *gin.RouterGroup, p *presenter.Presenter, service product.Usecase) {
	group := rg.Group("/manufacturers")
	group.GET("/", listManufacturers(p, service))
	group.GET("/:id", getManufacturer(p, service))
}

func AddProductWriteRoutes(rg *gin.RouterGroup, p *presenter.Presenter, service product.Usecase) {
	group := rg.Group("/manufacturers")
	group.POST("/", createManufacturer(p, service))
	group.PUT("/:id", updateManufacturer(p, service))
	group.DELETE("/:id", deleteManufacturer(p, service))
}

func listManufacturers(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call service
		search := c.Query("s")
		var err error
		var result []*entity.Manufacturer
		if search != "" {
			result, err = service.SearchManufacturers(search)
		} else {
			result, err = service.ListManufacturers()
		}

		// Handle errors
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
		body := &presenter.ManufacturerData{}
		if err := c.BindJSON(body); err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		manufacturer, err := service.CreateManufacturer(body.Name, body.WebsiteURL)
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
		body := &presenter.ManufacturerData{}
		if err := c.BindJSON(body); err != nil {
			c.String(errToResponse(err))
			return
		}

		// Build entity
		e, err := p.ManufacturerToEntity(id, *body)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
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
