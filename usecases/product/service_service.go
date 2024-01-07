package product

import (
	"github.com/JenswBE/go-commerce/entities"
)

// GetService fetches a single service by ID
func (s *Service) GetService(id entities.ID) (*entities.Service, error) {
	return s.db.GetService(id)
}

// CreateService creates a new service
func (s *Service) CreateService(service *entities.Service) (*entities.Service, error) {
	// Generate new ID
	service.ID = entities.NewID()

	// Validate entity
	err := service.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.CreateService(service)
}

// UpdateService persists the provided service
func (s *Service) UpdateService(service *entities.Service) (*entities.Service, error) {
	// Validate entity
	err := service.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.UpdateService(service)
}

// DeleteService deletes a single service by ID
func (s *Service) DeleteService(id entities.ID) error {
	return s.db.DeleteService(id)
}
