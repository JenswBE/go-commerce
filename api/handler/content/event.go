package content

import (
	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/openapi"
	presenter "github.com/JenswBE/go-commerce/api/presenter/content"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *ContentHandler) listEvents(c *gin.Context) {
	// Call service
	result, err := h.service.ListEvents()
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

func (h *ContentHandler) createEvent(c *gin.Context) {
	// Parse body
	body := &openapi.Event{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	e := presenter.EventToEntity(h.presenter, uuid.Nil, *body)
	event, err := h.service.CreateEvent(e)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(201, presenter.EventFromEntity(h.presenter, event))
}

func (h *ContentHandler) updateEvent(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Parse body
	body := &openapi.Event{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Call service
	e := presenter.EventToEntity(h.presenter, id, *body)
	event, err := h.service.UpdateEvent(e)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.JSON(200, presenter.EventFromEntity(h.presenter, event))
}

func (h *ContentHandler) deleteEvent(c *gin.Context) {
	// Parse params
	id, ok := handler.ParseIDParam(c, "id", h.presenter)
	if !ok {
		return // Response already set on Gin context
	}

	// Call service
	err := h.service.DeleteEvent(id)
	if err != nil {
		c.JSON(handler.ErrToResponse(err))
		return
	}

	// Handle success
	c.String(204, "")
}
