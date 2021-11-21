package content

import (
	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/openapi"
	presenter "github.com/JenswBE/go-commerce/api/presenter/content"
	"github.com/gin-gonic/gin"
)

func (h *ContentHandler) listContent(c *gin.Context) {
	// Call service
	result, err := h.service.ListContent()
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ContentListFromEntity(h.presenter, result))
}

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

func (h *ContentHandler) updateContent(c *gin.Context) {
	// Parse body
	body := &openapi.Content{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Parse params
	body.Name = c.Param("content_name")

	// Call service
	e := presenter.ContentToEntity(h.presenter, *body)
	content, err := h.service.UpdateContent(e)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.ContentFromEntity(h.presenter, content))
}
