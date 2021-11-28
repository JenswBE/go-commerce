package content

import (
	"github.com/JenswBE/go-commerce/entities"
)

type DatabaseRepository interface {
	GetEvent(id entities.ID) (*entities.Event, error)
	ListEvents(includePastEvents bool) ([]*entities.Event, error)
	CreateEvent(e *entities.Event) (*entities.Event, error)
	UpdateEvent(e *entities.Event) (*entities.Event, error)
	DeleteEvent(id entities.ID) error

	GetContent(name string) (*entities.Content, error)
	ListContent() ([]*entities.Content, error)
	CreateContent(e *entities.Content) (*entities.Content, error)
	UpdateContent(e *entities.Content) (*entities.Content, error)
	DeleteContent(name string) error
}

type Usecase interface {
	GetEvent(id entities.ID) (*entities.Event, error)
	ListEvents(includePastEvents bool) ([]*entities.Event, error)
	CreateEvent(e *entities.Event) (*entities.Event, error)
	UpdateEvent(e *entities.Event) (*entities.Event, error)
	DeleteEvent(id entities.ID) error

	GetContent(name string) (*entities.Content, error)
	ListContent() ([]*entities.Content, error)
	CreateContent(e *entities.Content) (*entities.Content, error)
	UpdateContent(e *entities.Content) (*entities.Content, error)
	DeleteContent(name string) error
}
