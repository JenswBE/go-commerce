package admin

import (
	"fmt"
	"net/http"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
)

const paramEventID = "event_id"

func (h *AdminHandler) handleEventsList(c *gin.Context) {
	// Get ShowPastEvents state
	showPastEvents, err := handleStatefulBoolFlag(c, "show_past_events")
	if err != nil {
		redirectWithMessage(c, sessions.Default(c), entities.MessageTypeError, err.Error(), "events/")
		return
	}

	// Fetch events
	events, err := h.contentService.ListEvents(showPastEvents)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ophalen van evenementen mislukt: %v", err)
	}

	htmlWithFlashes(c, http.StatusOK, "eventsList", &entities.EventsListData{
		BaseData: entities.BaseData{
			Title:      "Evenementen",
			ParentPath: "events",
		},
		Events:         events,
		ShowPastEvents: showPastEvents,
	})
}

func (h *AdminHandler) handleEventsFormGET(c *gin.Context) {
	// Check if new object
	paramID := c.Param(paramEventID)
	if paramID == "new" {
		c.HTML(http.StatusOK, "eventsForm", entities.EventsFormData{
			BaseData: entities.BaseData{
				Title:      eventsFormTitle(true),
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
			Title:      eventsFormTitle(false),
			ParentPath: "events",
		},
		IsNew: false,
		Event: entities.EventFromEntity(event),
	})
}

func (h *AdminHandler) handleEventsFormPOST(c *gin.Context) {
	// Check if new object
	paramID := c.Param(paramEventID)
	isNew := paramID == "new"

	// Parse body
	event := entities.Event{}
	err := c.MustBindWith(&event, binding.FormPost)
	if err != nil {
		renderEventsFormWithError(c, isNew, event, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Convert form to entity
	eventEntity, err := event.ToEntity()
	if err != nil {
		renderEventsFormWithError(c, isNew, event, fmt.Sprintf("Ongeldige data ontvangen: %v", err))
		return
	}

	// Create new entity
	if isNew {
		_, err := h.contentService.CreateEvent(&eventEntity)
		if err != nil {
			renderEventsFormWithError(c, isNew, event, fmt.Sprintf("Toevoegen van evenement mislukt: %v", err))
			return
		}
	} else {
		// Parse ID parameter
		eventEntity.ID, err = parseUUID(paramID, objectTypeManufacturer)
		if err != nil {
			renderEventsFormWithError(c, isNew, event, fmt.Sprintf("Ongeldige evenement ID %s: %v", paramID, err))
			return
		}

		// Update event
		_, err := h.contentService.UpdateEvent(&eventEntity)
		if err != nil {
			renderEventsFormWithError(c, isNew, event, fmt.Sprintf("Aanpassen van evenement mislukt: %v", err))
			return
		}
	}

	// Upsert successful
	redirectWithMessage(c, sessions.Default(c), entities.MessageTypeSuccess, "Evenement successvol toegevoegd/aangepast.", "events/")
}

func eventsFormTitle(isNew bool) string {
	if isNew {
		return "Evenement toevoegen"
	}
	return "Eventement aanpassen"
}

func renderEventsFormWithError(c *gin.Context, isNew bool, event entities.Event, message string) {
	c.HTML(http.StatusOK, "eventsForm", &entities.EventsFormData{
		BaseData: entities.BaseData{
			Title:      eventsFormTitle(isNew),
			ParentPath: "events",
			Messages: []entities.Message{{
				Type:    entities.MessageTypeError,
				Content: message,
			}},
		},
		IsNew: isNew,
		Event: event,
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
