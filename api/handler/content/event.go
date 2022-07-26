package content

import (
	"strings"

	"github.com/JenswBE/go-commerce/api/handler"
	presenter "github.com/JenswBE/go-commerce/api/presenter/content"
	"github.com/gin-gonic/gin"
)

func (h *ContentHandler) listEvents(c *gin.Context) {
	// Parse params
	includePastEvents := strings.ToLower(c.Query("include_past_events")) == "true"

	// Call service
	result, err := h.service.ListEvents(includePastEvents)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.EventListFromEntity(h.presenter, result))
}

func (h *ContentHandler) getEvent(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	event, err := h.service.GetEvent(id)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.EventFromEntity(h.presenter, event))
}
