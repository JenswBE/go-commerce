package fixtures

import (
	"time"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

// #############################
// #           ENTITY          #
// #############################

func Event() *entities.Event {
	return &entities.Event{
		ID:          generics.Must(entities.NewIDFromString(EventID)),
		Name:        "test-name",
		Description: "test-description",
		EventType:   "test-event-type",
		Start:       time.Date(2021, 11, 21, 11, 25, 0, 0, time.UTC),
		End:         time.Date(2021, 11, 22, 18, 45, 0, 0, time.UTC),
		WholeDay:    true,
	}
}

func EventSlice() []*entities.Event {
	return []*entities.Event{
		Event(),
	}
}

// #############################
// #          OPENAPI          #
// #############################

func EventOpenAPI() *openapi.Event {
	return &openapi.Event{
		Id:          openapi.PtrString(EventID),
		Name:        "test-name",
		Description: openapi.PtrString("test-description"),
		EventType:   "test-event-type",
		Start:       time.Date(2021, 11, 21, 11, 25, 0, 0, time.UTC),
		End:         time.Date(2021, 11, 22, 18, 45, 0, 0, time.UTC),
		WholeDay:    openapi.PtrBool(true),
	}
}

func EventOpenAPISlice() []openapi.Event {
	return []openapi.Event{
		*EventOpenAPI(),
	}
}

func EventListOpenAPI() *openapi.EventList {
	return openapi.NewEventList(EventOpenAPISlice())
}
