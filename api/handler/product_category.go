package handler

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *ProductHandler) listCategories(c *gin.Context) {
	// Parse params
	imageConfig, err := parseImageConfigParams(c)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Call service
	result, err := h.service.ListCategories(imageConfig)

	// Handle errors
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.CategoryListFromEntity(result))
}

func (h *ProductHandler) getCategory(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}
	imageConfig, err := parseImageConfigParams(c)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Call service
	category, err := h.service.GetCategory(id, imageConfig)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.CategoryFromEntity(category))
}

func (h *ProductHandler) createCategory(c *gin.Context) {
	// Parse body
	body := &openapi.Category{}
	if err := c.BindJSON(body); err != nil {
		c.String(errToResponse(err))
		return
	}

	// Build entity
	e, err := h.presenter.CategoryToEntity(uuid.Nil, *body)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Call service
	category, err := h.service.CreateCategory(e)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(201, h.presenter.CategoryFromEntity(category))
}

func (h *ProductHandler) updateCategory(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
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
	e, err := h.presenter.CategoryToEntity(id, *body)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Call service
	category, err := h.service.UpdateCategory(e)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, h.presenter.CategoryFromEntity(category))
}

func (h *ProductHandler) deleteCategory(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteCategory(id)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}

func (h *ProductHandler) upsertCategoryImage(c *gin.Context) {
	// Parse params
	id, ok := parseIDParam(c, "id", h.presenter)
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
	var category *entities.Category
	for imageName, imageContent := range images {
		category, err = h.service.UpsertCategoryImage(id, imageName, imageContent, imageConfig)
		if err != nil {
			c.String(errToResponse(err))
			return
		}
	}

	// Handle success
	c.JSON(200, h.presenter.ImageFromEntity(category.Image))
}

func (h *ProductHandler) deleteCategoryImage(c *gin.Context) {
	// Parse params
	categoryID, ok := parseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteCategoryImage(categoryID)
	if err != nil {
		c.String(errToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}
