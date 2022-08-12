package internal

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

type Image struct {
	ID        string `gorm:"type:uuid"`
	OwnerID   string `gorm:"type:uuid"`
	OwnerType string

	Extension string // File extension
	Order     int
}

func (image *Image) ToEntity() *entities.Image {
	if image == nil {
		return nil
	}
	return &entities.Image{
		ID:        generics.Must(entities.NewIDFromString(image.ID)),
		Extension: image.Extension,
		Order:     image.Order,
	}
}

func ImagesListPgToEntity(images []Image) []*entities.Image {
	if images == nil {
		return nil
	}
	output := make([]*entities.Image, 0, len(images))
	for _, image := range images {
		output = append(output, image.ToEntity())
	}
	return output
}

func ImageEntityToPg(e *entities.Image) *Image {
	if e == nil {
		return nil
	}
	return &Image{
		ID:        e.ID.String(),
		Extension: e.Extension,
		Order:     e.Order,
	}
}

func ImagesListEntityToPg(images []*entities.Image) []Image {
	if images == nil {
		return nil
	}
	output := make([]Image, 0, len(images))
	for _, image := range images {
		output = append(output, *ImageEntityToPg(image))
	}
	return output
}
