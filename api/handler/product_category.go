package handler

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/usecase/product"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func listCategories(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call service
		result, err := service.ListCategories()

		// Handle errors
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.CategoriesListFromEntity(result))
	}
}

func getCategory(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		category, err := service.GetCategory(id)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.CategoryFromEntity(category))
	}
}

func createCategory(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse body
		body := &openapi.Category{}
		if err := c.BindJSON(body); err != nil {
			c.String(errToResponse(err))
			return
		}

		// Build entity
		e, err := p.CategoryToEntity(uuid.Nil, *body)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		category, err := service.CreateCategory(e)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.CategoryFromEntity(category))
	}
}

func updateCategory(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Parse body
		body := &openapi.Category{}
		if err := c.BindJSON(body); err != nil {
			c.String(errToResponse(err))
			return
		}

		// Build entity
		e, err := p.CategoryToEntity(id, *body)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Call service
		category, err := service.UpdateCategory(e)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.JSON(200, p.CategoryFromEntity(category))
	}
}

func deleteCategory(p *presenter.Presenter, service product.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse params
		id, ok := parseParamID(c, p)
		if !ok {
			return // Response already set on Gin context
		}

		// Call service
		err := service.DeleteCategory(id)
		if err != nil {
			c.String(errToResponse(err))
			return
		}

		// Handle success
		c.String(204, "")
	}
}
