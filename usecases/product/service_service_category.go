package product

import (
	"github.com/JenswBE/go-commerce/entities"
)

// GetServiceCategory fetches a single service category by ID
func (s *Service) GetServiceCategory(id entities.ID, resolved bool) (*entities.ResolvedServiceCategory, error) {
	// Fetch serviceCategory
	serviceCategory, err := s.db.GetServiceCategory(id)
	if err != nil {
		return nil, err
	}

	// Resolve service category
	resolvedServiceCategory := &entities.ResolvedServiceCategory{ServiceCategory: *serviceCategory}
	if resolved {
		resolvedServiceCategory.Services, err = s.db.ListServices(id)
		if err != nil {
			return nil, err
		}
	}

	// Get successful
	return resolvedServiceCategory, nil
}

// ListServiceCategories fetches all service categories
func (s *Service) ListServiceCategories(resolved bool) ([]*entities.ResolvedServiceCategory, error) {
	// Fetch service categories
	serviceCategories, err := s.db.ListServiceCategories()
	if err != nil {
		return nil, err
	}

	// Resolve service category
	resolvedServiceCategories := make([]*entities.ResolvedServiceCategory, 0, len(serviceCategories))
	for _, svcCat := range serviceCategories {
		resolvedServiceCategory := &entities.ResolvedServiceCategory{ServiceCategory: *svcCat}
		if resolved {
			resolvedServiceCategory.Services, err = s.db.ListServices(svcCat.ID)
			if err != nil {
				return nil, err
			}
		}
		resolvedServiceCategories = append(resolvedServiceCategories, resolvedServiceCategory)
	}

	// Get successful
	return resolvedServiceCategories, nil
}

// CreateServiceCategory creates a new service category
func (s *Service) CreateServiceCategory(serviceCategory *entities.ServiceCategory) (*entities.ServiceCategory, error) {
	// Generate new ID
	serviceCategory.ID = entities.NewID()

	// Validate entity
	err := serviceCategory.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.CreateServiceCategory(serviceCategory)
}

// UpdateServiceCategory persists the provided service category
func (s *Service) UpdateServiceCategory(serviceCategory *entities.ServiceCategory) (*entities.ServiceCategory, error) {
	// Validate entity
	err := serviceCategory.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.UpdateServiceCategory(serviceCategory)
}

// DeleteServiceCategory deletes a single serviceCategory by ID
func (s *Service) DeleteServiceCategory(id entities.ID) error {
	return s.db.DeleteServiceCategory(id)
}
