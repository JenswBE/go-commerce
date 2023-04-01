package content

import (
	"github.com/gin-gonic/gin"

	"github.com/JenswBE/go-commerce/api/handler"
	presenter "github.com/JenswBE/go-commerce/api/presenter/content"
)

func (h *ContentHandler) getContent(c *gin.Context) {
	// Call service
	content, err := h.service.GetContent(c.Param("content_name"))
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	response := presenter.ContentFromEntity(h.presenter, content)
	c.JSON(200, response)
}
