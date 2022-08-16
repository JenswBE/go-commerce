package product

import (
	"github.com/rs/zerolog/log"

	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

// GetCategory fetches a single category by ID
func (s *Service) GetCategory(id entities.ID, imageConfigs map[string]imageproxy.ImageConfig) (*entities.Category, error) {
	// Fetch category
	category, err := s.db.GetCategory(id)
	if err != nil {
		return nil, err
	}

	// Generate URL's
	if len(imageConfigs) > 0 {
		err = category.Image.SetURLsFromConfigs(s.imageProxy, imageConfigs)
		if err != nil {
			return nil, err
		}
	}

	// Get successful
	return category, nil
}

// ListCategories fetches all categories
func (s *Service) ListCategories(imageConfigs map[string]imageproxy.ImageConfig) ([]*entities.Category, error) {
	// Fetch categories
	categories, err := s.db.ListCategories()
	if err != nil {
		return nil, err
	}

	// Generate URL's
	if len(imageConfigs) > 0 {
		for _, category := range categories {
			err = category.Image.SetURLsFromConfigs(s.imageProxy, imageConfigs)
			if err != nil {
				return nil, err
			}
		}
	}

	// Get successful
	return categories, nil
}

// CreateCategory creates a new category
func (s *Service) CreateCategory(category *entities.Category) (*entities.Category, error) {
	// Generate new ID
	category.ID = entities.NewID()

	// Validate entity
	err := category.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.CreateCategory(category)
}

// UpdateCategory persists the provided category
func (s *Service) UpdateCategory(category *entities.Category) (*entities.Category, error) {
	// Validate entity
	err := category.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.UpdateCategory(category)
}

// DeleteCategory deletes a single category by ID
func (s *Service) DeleteCategory(id entities.ID) error {
	// Fetch category
	category, err := s.db.GetCategory(id)
	if err != nil {
		return err
	}

	// Delete image if set
	if category.Image != nil {
		err := s.deleteImage(category.Image)
		if err != nil {
			return err
		}
	}

	// Delete category
	return s.db.DeleteCategory(id)
}

// UpsertCategoryImage adds/updates the image of the category with the provided one.
func (s *Service) UpsertCategoryImage(categoryID entities.ID, imageName string, imageContent []byte, imageConfigs map[string]imageproxy.ImageConfig) (*entities.Category, error) {
	// Fetch category
	category, err := s.db.GetCategory(categoryID)
	if err != nil {
		return nil, err
	}

	// Set image
	oldImage := category.Image
	category.Image, err = s.saveImage(imageName, imageContent)
	if err != nil {
		return nil, err
	}

	// Save category
	category, err = s.db.UpdateCategory(category)
	if err != nil {
		return nil, err
	}

	// Delete old image
	if oldImage != nil {
		err = s.deleteImage(oldImage)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to delete old image for category")
		}
	}

	// Generate URL's
	if len(imageConfigs) > 0 {
		err := category.Image.SetURLsFromConfigs(s.imageProxy, imageConfigs)
		if err != nil {
			return nil, err
		}
	}

	// Add images successful
	return category, nil
}

// DeleteCategoryImage cleares the image of the category
func (s *Service) DeleteCategoryImage(categoryID entities.ID) error {
	// Fetch category
	category, err := s.db.GetCategory(categoryID)
	if err != nil {
		return err
	}

	// Category has no image => No action needed
	if category.Image == nil {
		return nil
	}

	// Delete image
	return s.deleteImage(category.Image)
}
