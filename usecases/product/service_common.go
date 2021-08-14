package product

import (
	"errors"
	"path/filepath"

	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/google/uuid"
)

type Service struct {
	db           DatabaseRepository
	imageProxy   imageproxy.Service
	imageStorage StorageRepository
}

func NewService(db DatabaseRepository, imageProxy imageproxy.Service, imageStorage StorageRepository) *Service {
	return &Service{
		db:           db,
		imageProxy:   imageProxy,
		imageStorage: imageStorage,
	}
}

func (s *Service) setImageURLsFromConfig(images []*entities.Image, configs map[string]imageproxy.ImageConfig) error {
	for _, image := range images {
		err := image.SetURLsFromConfigs(s.imageProxy, configs)
		if err != nil {
			return err
		}
	}
	return nil
}

// saveImage saves a single image
func (s *Service) saveImage(filename string, content []byte) (*entities.Image, error) {
	// Extract extension from filename
	imageExt := filepath.Ext(filename)
	if imageExt == "" {
		return nil, errors.New("cannot save image without knowing extension")
	}

	// Save images as files
	imageID := uuid.New()
	err := s.imageStorage.SaveFile(imageID.String()+imageExt, content)
	if err != nil {
		return nil, err
	}

	// Build and return image entity
	image := &entities.Image{
		ID:        imageID,
		Extension: imageExt,
	}
	return image, nil
}

// deleteImage deletes a single image
func (s *Service) deleteImage(image *entities.Image) error {
	// Delete from DB first as it has the highest risk to fail
	err := s.db.DeleteImage(image.ID)
	if err != nil {
		return err
	}

	// Delete from storage
	return s.imageStorage.DeleteFile(image.ID.String() + image.Extension)
}
