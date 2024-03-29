package entities

import (
	"github.com/JenswBE/go-commerce/entities"
)

type EventsListTemplate struct {
	BaseData
	Events         []*entities.Event
	ShowPastEvents bool
}

func (t EventsListTemplate) GetTemplateName() string {
	return "eventsList"
}

type EventsFormTemplate struct {
	BaseData
	IsNew bool
	Event Event
}

func (t EventsFormTemplate) GetTemplateName() string {
	return "eventsForm"
}

type Event struct {
	Name  string `form:"name"`
	Start string `form:"start"`
	End   string `form:"end"`
}

func EventFromEntity(e *entities.Event) Event {
	return Event{
		Name:  e.Name,
		Start: e.Start.Format(TimeFormatDate),
		End:   e.End.Format(TimeFormatDate),
	}
}

func (e Event) ToEntity() (entities.Event, error) {
	start, err := parseDateString(e.Start)
	if err != nil {
		return entities.Event{}, err
	}

	end, err := parseDateString(e.End)
	if err != nil {
		return entities.Event{}, err
	}

	return entities.Event{
		Name:  e.Name,
		Start: start,
		End:   end,
	}, nil
}
