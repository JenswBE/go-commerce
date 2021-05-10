package handler

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-gonic/gin"
)

func AddProductReadRoutes(rg *gin.RouterGroup, service product.Usecase) {
	group := rg.Group("/manufacturers")
	group.GET("/", listManufacturers(service))
	group.GET("/:id", getManufacturer(service))
}

func AddProductWriteRoutes(rg *gin.RouterGroup, service product.Usecase) {
	group := rg.Group("/manufacturers")
	group.POST("/", createManufacturer(service))
	group.PUT("/:id", updateManufacturer(service))
	group.DELETE("/:id", deleteManufacturer(service))
}

func listManufacturers(service product.Usecase) gin.HandlerFunc {
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
			c.String(400, err.Error())
			return
		}

		// Handle success
		p := presenter.New()
		c.JSON(200, p.ManufacturersListFromEntity(result))
	}
}

func getManufacturer(service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		p := presenter.New()
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		manufacturer, err := service.GetManufacturer(id)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		// Handle success
		c.JSON(200, p.ManufacturerFromEntity(manufacturer))
	}
}

func createManufacturer(service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse body
		body := &presenter.ManufacturerData{}
		if err := c.BindJSON(body); err != nil {
			c.String(400, err.Error())
			return
		}

		// Call service
		manufacturer, err := service.CreateManufacturer(body.Name, body.WebsiteURL)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		// Handle success
		p := presenter.New()
		c.JSON(200, p.ManufacturerFromEntity(manufacturer))
	}
}

func updateManufacturer(service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		p := presenter.New()
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Parse body
		body := &presenter.ManufacturerData{}
		if err := c.BindJSON(body); err != nil {
			c.String(400, err.Error())
			return
		}

		// Build entity
		e, err := p.ManufacturerToEntity(id, *body)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		// Call service
		manufacturer, err := service.UpdateManufacturer(e)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		// Handle success
		c.JSON(200, p.ManufacturerFromEntity(manufacturer))
	}
}

func deleteManufacturer(service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		p := presenter.New()
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		err := service.DeleteManufacturer(id)
		if err != nil {
			c.String(400, err.Error())
			return
		}

		// Handle success
		c.String(204, "")
	}
}
