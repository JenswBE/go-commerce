package content

import (
	"github.com/JenswBE/go-commerce/api/handler"
	presenter "github.com/JenswBE/go-commerce/api/presenter/content"
	"github.com/gin-gonic/gin"
)

func (h *ContentHandler) getContent(c *gin.Context) {
	// Call service
	content, err := h.service.GetContent(c.Param("content_name"))
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ContentFromEntity(h.presenter, content))
}
