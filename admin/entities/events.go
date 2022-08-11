package entities

import (
	"github.com/JenswBE/go-commerce/entities"
)

type EventsListData struct {
	BaseData
	Events         []*entities.Event
	ShowPastEvents bool
}

type EventsFormData struct {
	BaseData
	IsNew bool
	Event Event
}

type Event struct {
	Name  string `form:"name"`
	Start string `form:"start"`
	End   string `form:"end"`
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
