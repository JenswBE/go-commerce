package product

import (
	"github.com/JenswBE/go-commerce/entity"
)

// GetCategory fetches a single category by ID
func (s *Service) GetCategory(id entity.ID) (*entity.Category, error) {
	return s.repo.GetCategory(id)
}

// ListCategories fetches all categories
func (s *Service) ListCategories() ([]*entity.Category, error) {
	return s.repo.ListCategories()
}

// CreateCategory creates a new category
func (s *Service) CreateCategory(name, description string, parentID entity.ID) (*entity.Category, error) {
	// Create entity
	m, err := entity.NewCategory(name, description, parentID)
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.repo.CreateCategory(m)
}

// UpdateCategory persists the provided category
func (s *Service) UpdateCategory(e *entity.Category) (*entity.Category, error) {
	// Validate entity
	err := e.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.repo.UpdateCategory(e)
}

// DeleteCategory deletes a single category by ID
func (s *Service) DeleteCategory(id entity.ID) error {
	return s.repo.DeleteCategory(id)
}
