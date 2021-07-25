package fixtures

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func Image() *entities.Image {
	return &entities.Image{
		ID:        uuid.MustParse(ImageID),
		Extension: ".jpg",
		Order:     1,
		URL:       "http://image.test",
	}
}
