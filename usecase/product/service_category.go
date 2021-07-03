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
func (s *Service) CreateCategory(category *entity.Category) (*entity.Category, error) {
	// Generate new ID
	category.ID = entity.NewID()

	// Validate entity
	err := category.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.repo.CreateCategory(category)
}

// UpdateCategory persists the provided category
func (s *Service) UpdateCategory(category *entity.Category) (*entity.Category, error) {
	// Validate entity
	err := category.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.repo.UpdateCategory(category)
}

// DeleteCategory deletes a single category by ID
func (s *Service) DeleteCategory(id entity.ID) error {
	return s.repo.DeleteCategory(id)
}
