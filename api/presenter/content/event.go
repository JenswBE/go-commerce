package content

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
)

func EventFromEntity(p *presenter.Presenter, input *entities.Event) openapi.Event {
	// Set basic fields
	output := openapi.NewEvent(
		p.EncodeID(input.ID),
		p.String(input.Name),
		input.EventType,
		input.Start,
		input.End,
	)
	output.SetDescription(p.String(input.Description))
	output.SetWholeDay(input.WholeDay)
	return *output
}

func EventListFromEntity(p *presenter.Presenter, input []*entities.Event) openapi.EventList {
	return *openapi.NewEventList(presenter.SliceFromEntity(p, input, EventFromEntity))
}
