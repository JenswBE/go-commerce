package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
)

const paramEventID = "event_id"

func (h *Handler) handleEventsList(c *gin.Context) {
	// Define base data
	baseData := entities.BaseData{
		Title:      "Evenementen",
		ParentPath: "events",
	}

	// Get ShowPastEvents state
	showPastEvents, err := handleStatefulBoolFlag(c, "show_past_events")
	if err != nil {
		baseData.AddMessage(entities.MessageTypeError, "Verwerken van show_past_events mislukt: %v", err)
		html(c, http.StatusBadRequest, &entities.EventsListTemplate{BaseData: baseData})
		return
	}

	// Fetch events
	events, err := h.contentService.ListEvents(showPastEvents)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ophalen van evenementen mislukt: %v", err)
		return
	}

	htmlWithFlashes(c, http.StatusOK, &entities.EventsListTemplate{
		BaseData: entities.BaseData{
			Title:      "Evenementen",
			ParentPath: "events",
		},
		Events:         events,
		ShowPastEvents: showPastEvents,
	})
}

func (h *Handler) handleEventsFormGET(c *gin.Context) {
	// Check if new event
	paramID := c.Param(paramEventID)
	if paramID == "new" {
		html(c, http.StatusOK, &entities.EventsFormTemplate{
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
	id, err := parseID(paramID, i18n.ObjectTypeEvent)
	if err != nil {
		log.Debug().Err(err).Str("event_id", paramID).Msg("Invalid event ID provided")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "events/")
		return
	}

	// Fetch event
	event, err := h.contentService.GetEvent(id)
	if err != nil {
		log.Debug().Err(err).Str("event_id", paramID).Msg("Failed to fetch event")
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "events/")
		return
	}

	// Render page
	html(c, http.StatusOK, &entities.EventsFormTemplate{
		BaseData: entities.BaseData{
			Title:      eventsFormTitle(false),
			ParentPath: "events",
		},
		IsNew: false,
		Event: entities.EventFromEntity(event),
	})
}

func (h *Handler) handleEventsFormPOST(c *gin.Context) {
	// Check if new event
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
		eventEntity.ID, err = parseID(paramID, i18n.ObjectTypeEvent)
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
	html(c, http.StatusOK, &entities.EventsFormTemplate{
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

func (h *Handler) handleEventsDelete(c *gin.Context) {
	// Get session
	session := sessions.Default(c)

	// Parse parameters
	id, err := parseID(c.Param(paramEventID), i18n.ObjectTypeEvent)
	if err != nil {
		redirectWithMessage(c, session, entities.MessageTypeError, err.Error(), "events/")
		return
	}

	// Call service
	err = h.contentService.DeleteEvent(id)
	if err != nil {
		msg := i18n.DeleteFailed(i18n.ObjectTypeEvent, "", err)
		redirectWithMessage(c, session, entities.MessageTypeError, msg, "events/")
		return
	}

	// Call successful
	msg := i18n.DeleteSuccessful(i18n.ObjectTypeEvent)
	redirectWithMessage(c, session, entities.MessageTypeSuccess, msg, "events/")
}
