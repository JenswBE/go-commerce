package entities

import (
	"strings"
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

func (e *Event) Clean() {
	e.Name = strings.TrimSpace(e.Name)
	e.Description = strings.TrimSpace(e.Description)
}

// Validate cleans and validates the event data
func (e *Event) Validate() error {
	// Clean entity
	e.Clean()

	// Validate simple fields
	if e.End.Before(e.Start) {
		return NewError(400, openapi.GOCOMERRORCODE_EVENT_END_BEFORE_START, e.ID.String(), nil)
	}

	// Entity is valid
	return nil
}
