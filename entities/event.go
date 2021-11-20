package entities

import (
	"time"

	"github.com/JenswBE/go-commerce/api/openapi"
)

// Event data
type Event struct {
	ID          ID
	Name        string
	Description string
	EventType   string
	Start       time.Time
	End         time.Time
	WholeDay    bool
}

// Validate validates the event data
func (e *Event) Validate() error {
	// Validate simple fields
	if !e.End.IsZero() && e.End.Before(e.Start) {
		return NewError(400, openapi.GOCOMERRORCODE_EVENT_END_BEFORE_START, e.ID.String(), nil)
	}

	// Entity is valid
	return nil
}
