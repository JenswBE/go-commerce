package content

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func EventFromEntity(p *presenter.Presenter, input *entities.Event) openapi.Event {
	// Set basic fields
	output := openapi.NewEvent(p.String(input.Name), input.EventType, input.Start, input.End)
	output.SetId(p.EncodeID(input.ID))
	output.SetDescription(p.String(input.Description))
	output.SetWholeDay(input.WholeDay)
	return *output
}

func EventSliceFromEntity(p *presenter.Presenter, input []*entities.Event) []openapi.Event {
	output := make([]openapi.Event, 0, len(input))
	for _, event := range input {
		output = append(output, EventFromEntity(p, event))
	}
	return output
}

func EventListFromEntity(p *presenter.Presenter, input []*entities.Event) openapi.EventList {
	return *openapi.NewEventList(EventSliceFromEntity(p, input))
}

func EventToEntity(p *presenter.Presenter, id uuid.UUID, event openapi.Event) *entities.Event {
	// Build entity
	return &entities.Event{
		ID:          id,
		Name:        event.GetName(),
		Description: event.GetDescription(),
		EventType:   event.GetEventType(),
		Start:       event.GetStart(),
		End:         event.GetEnd(),
		WholeDay:    event.GetWholeDay(),
	}
}
