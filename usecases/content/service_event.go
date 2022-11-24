package content

import (
	"github.com/JenswBE/go-commerce/entities"
)

// GetEvent fetches a single event by ID
func (s *Service) GetEvent(id entities.ID) (*entities.Event, error) {
	// Fetch event
	event, err := s.db.GetEvent(id)
	if err != nil {
		return nil, err
	}

	// Get successful
	return event, nil
}

// ListEvents fetches all events
func (s *Service) ListEvents(includePastEvents bool) ([]*entities.Event, error) {
	// Fetch events
	events, err := s.db.ListEvents(includePastEvents)
	if err != nil {
		return nil, err
	}

	// Get successful
	return events, nil
}

// CreateEvent creates a new event
func (s *Service) CreateEvent(event *entities.Event) (*entities.Event, error) {
	// Generate new ID
	event.ID = entities.NewID()
	if s.eventsWholeDaysOnly {
		event.WholeDay = true
	}

	// Validate entity
	err := event.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.CreateEvent(event)
}

// UpdateEvent persists the provided event
func (s *Service) UpdateEvent(event *entities.Event) (*entities.Event, error) {
	// Validate entity
	err := event.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.UpdateEvent(event)
}

// DeleteEvent deletes a single event by ID
func (s *Service) DeleteEvent(id entities.ID) error {
	return s.db.DeleteEvent(id)
}
