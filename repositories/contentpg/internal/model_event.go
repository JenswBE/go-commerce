package internal

import (
	"time"

	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

type Event struct {
	Base
	Name        string
	Description string
	EventType   string
	Start       time.Time
	End         time.Time
	WholeDay    bool
}

func (e *Event) ToEntity() *entities.Event {
	return &entities.Event{
		ID:          generics.Must(entities.NewIDFromString(e.ID)),
		Name:        e.Name,
		Description: e.Description,
		EventType:   e.EventType,
		Start:       e.Start,
		End:         e.End,
		WholeDay:    e.WholeDay,
	}
}

func EventEntityToPg(e *entities.Event) *Event {
	return &Event{
		Base:        Base{ID: e.ID.String()},
		Name:        e.Name,
		Description: e.Description,
		EventType:   e.EventType,
		Start:       e.Start,
		End:         e.End,
		WholeDay:    e.WholeDay,
	}
}
