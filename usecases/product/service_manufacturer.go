package product

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/rs/zerolog/log"
)

// GetManufacturer fetches a single manufacturer by ID
func (s *Service) GetManufacturer(id entities.ID, imageConfig *imageproxy.ImageConfig) (*entities.Manufacturer, error) {
	return s.db.GetManufacturer(id)
}

// ListManufacturers fetches all manufacturers
func (s *Service) ListManufacturers(imageConfig *imageproxy.ImageConfig) ([]*entities.Manufacturer, error) {
	return s.db.ListManufacturers()
}

// CreateManufacturer creates a new manufacturer
func (s *Service) CreateManufacturer(manufacturer *entities.Manufacturer) (*entities.Manufacturer, error) {
	// Generate new ID
	manufacturer.ID = entities.NewID()

	// Validate entity
	err := manufacturer.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.CreateManufacturer(manufacturer)
}

// UpdateManufacturer persists the provided manufacturer
func (s *Service) UpdateManufacturer(manufacturer *entities.Manufacturer) (*entities.Manufacturer, error) {
	// Validate entity
	err := manufacturer.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.UpdateManufacturer(manufacturer)
}

// DeleteManufacturer deletes a single manufacturer by ID
func (s *Service) DeleteManufacturer(id entities.ID) error {
	// Fetch manufacturer
	manufacturer, err := s.db.GetManufacturer(id)
	if err != nil {
		return err
	}

	// Delete image if set
	if manufacturer.Image != nil {
		err := s.deleteImage(manufacturer.Image)
		if err != nil {
			return err
		}
	}

	// Delete manufacturer
	return s.db.DeleteManufacturer(id)
}

// UpsertManufacturerImage adds/updates the image of the manufacturer with the provided one.
func (s *Service) UpsertManufacturerImage(manufacturerID entities.ID, imageName string, imageContent []byte, imageConfig *imageproxy.ImageConfig) (*entities.Manufacturer, error) {
	// Fetch manufacturer
	manufacturer, err := s.db.GetManufacturer(manufacturerID)
	if err != nil {
		return nil, err
	}

	// Set image
	oldImage := manufacturer.Image
	manufacturer.Image, err = s.saveImage(imageName, imageContent)
	if err != nil {
		return nil, err
	}

	// Save manufacturer
	manufacturer, err = s.db.UpdateManufacturer(manufacturer)
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
	if imageConfig != nil {
		err := manufacturer.Image.SetURLFromConfig(s.imageProxy, *imageConfig)
		if err != nil {
			return nil, err
		}
	}

	// Add images successful
	return manufacturer, nil
}

// DeleteManufacturerImage cleares the image of the manufacturer
func (s *Service) DeleteManufacturerImage(manufacturerID entities.ID) error {
	// Fetch manufacturer
	manufacturer, err := s.db.GetManufacturer(manufacturerID)
	if err != nil {
		return err
	}

	// Manufacturer has no image => No action needed
	if manufacturer.Image == nil {
		return nil
	}

	// Delete image
	return s.deleteImage(manufacturer.Image)
}