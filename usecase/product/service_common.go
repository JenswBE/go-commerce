package product

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
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

func (s *Service) setImageURLsFromConfig(images []*entity.Image, config imageproxy.ImageConfig) error {
	for _, image := range images {
		err := image.SetURLFromConfig(s.imageProxy, config)
		if err != nil {
			return err
		}
	}
	return nil
}
