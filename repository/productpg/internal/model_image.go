package internal

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/google/uuid"
)

type Image struct {
	ID        string `gorm:"type:uuid"`
	OwnerID   string `gorm:"type:uuid"`
	OwnerType string

	Extension string // File extension
	Order     int
}

func ImagePgToEntity(image *Image) *entity.Image {
	if image == nil {
		return nil
	}
	return &entity.Image{
		ID:        uuid.MustParse(image.ID),
		Extension: image.Extension,
		Order:     image.Order,
	}
}

func ImagesListPgToEntity(images []Image) []*entity.Image {
	if images == nil {
		return nil
	}
	output := make([]*entity.Image, 0, len(images))
	for _, image := range images {
		output = append(output, ImagePgToEntity(&image))
	}
	return output
}

func ImageEntityToPg(e *entity.Image) *Image {
	if e == nil {
		return nil
	}
	return &Image{
		ID:        e.ID.String(),
		Extension: e.Extension,
		Order:     e.Order,
	}
}

func ImagesListEntityToPg(images []*entity.Image) []Image {
	if images == nil {
		return nil
	}
	output := make([]Image, 0, len(images))
	for _, image := range images {
		output = append(output, *ImageEntityToPg(image))
	}
	return output
}
