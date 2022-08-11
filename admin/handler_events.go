package admin

import (
	"net/http"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const paramEventID = "event_id"

func (h *AdminHandler) handleEventsList(c *gin.Context) {
	htmlWithFlashes(c, http.StatusOK, "eventsList", &entities.EventsListData{
		BaseData: entities.BaseData{
			Title:      "Evenementen",
			ParentPath: "events",
		},
	})
}

func (h *AdminHandler) handleEventsForm(c *gin.Context) {
	// Check if new object
	paramID := c.Param(paramEventID)
	if paramID == "new" {
		c.HTML(http.StatusOK, "eventsForm", entities.EventsFormData{
			BaseData: entities.BaseData{
				Title:      "Evenement toevoegen",
				ParentPath: "events",
			},
			IsNew: true,
		})
		return
	}

	// Get session
	session := sessions.Default(c)

	// Parse ID parameter
	id, err := parseUUID(paramID, objectTypeManufacturer)
	if err != nil {
		log.Debug().Err(err).Str("event_id", paramID).Msg("Invalid event ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "events/")
		return
	}

	// Fetch object
	event, err := h.contentService.GetEvent(id)
	if err != nil {
		log.Debug().Err(err).Str("event_id", paramID).Msg("Failed to fetch event")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "events/")
		return
	}

	// Render page
	c.HTML(http.StatusOK, "eventsForm", entities.EventsFormData{
		BaseData: entities.BaseData{
			Title:      "Evenement toevoegen",
			ParentPath: "events",
		},
		IsNew: false,
		Event: *event,
	})
}

func (h *AdminHandler) handleEventsDelete(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse parameters
	id, err := parseUUID(c.Param(paramEventID), objectTypeEvent)
	if err != nil {
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "events/")
		return
	}

	// Call service
	err = h.contentService.DeleteEvent(id)
	if err != nil {
		msg := i18n.DeleteFailed(i18n.ObjectTypeEvent, id, err)
		redirectWithMessage(c, session, entities.MessageTypeError, msg, "events/")
		return
	}

	// Call successful
	msg := i18n.DeleteSuccessful(i18n.ObjectTypeEvent)
	redirectWithMessage(c, session, entities.MessageTypeSuccess, msg, "events/")
}
