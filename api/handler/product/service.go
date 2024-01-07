package product

import (
	"github.com/gin-gonic/gin"

	"github.com/JenswBE/go-commerce/api/handler"
	presenter "github.com/JenswBE/go-commerce/api/presenter/product"
)

func (h *ProductHandler) listServiceCategories(c *gin.Context) {
	// Call service
	result, err := h.service.ListServiceCategories(true)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Convert to OpenAPI model
	output, err := presenter.ResolvedServiceCategoryListFromEntity(h.presenter, result)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, output)
}
