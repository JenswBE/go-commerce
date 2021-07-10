package product

import (
	"github.com/JenswBE/go-commerce/entity"
)

// GetManufacturer fetches a single manufacturer by ID
func (s *Service) GetManufacturer(id entity.ID) (*entity.Manufacturer, error) {
	return s.db.GetManufacturer(id)
}

// ListManufacturers fetches all manufacturers
func (s *Service) ListManufacturers() ([]*entity.Manufacturer, error) {
	return s.db.ListManufacturers()
}

// CreateManufacturer creates a new manufacturer
func (s *Service) CreateManufacturer(manufacturer *entity.Manufacturer) (*entity.Manufacturer, error) {
	// Generate new ID
	manufacturer.ID = entity.NewID()

	// Validate entity
	err := manufacturer.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.CreateManufacturer(manufacturer)
}

// UpdateManufacturer persists the provided manufacturer
func (s *Service) UpdateManufacturer(manufacturer *entity.Manufacturer) (*entity.Manufacturer, error) {
	// Validate entity
	err := manufacturer.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.UpdateManufacturer(manufacturer)
}

// DeleteManufacturer deletes a single manufacturer by ID
func (s *Service) DeleteManufacturer(id entity.ID) error {
	return s.db.DeleteManufacturer(id)
}
