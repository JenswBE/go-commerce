package product

import (
	"log"

	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

// GetCategory fetches a single category by ID
func (s *Service) GetCategory(id entities.ID, imageConfig *imageproxy.ImageConfig) (*entities.Category, error) {
	return s.db.GetCategory(id)
}

// ListCategories fetches all categories
func (s *Service) ListCategories(imageConfig *imageproxy.ImageConfig) ([]*entities.Category, error) {
	return s.db.ListCategories()
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
func (s *Service) UpsertCategoryImage(categoryID entities.ID, imageName string, imageContent []byte, imageConfig *imageproxy.ImageConfig) (*entities.Category, error) {
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
			log.Printf(`Failed to delete old image for category: %s`, err.Error())
		}
	}

	// Generate URL's
	if imageConfig != nil {
		err := category.Image.SetURLFromConfig(s.imageProxy, *imageConfig)
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
