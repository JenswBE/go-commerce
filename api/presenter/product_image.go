package presenter

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func (p *Presenter) ImageFromEntity(e *entities.Image) openapi.Image {
	return *openapi.NewImage(
		p.EncodeID(e.ID),
		e.Extension,
		e.URL,
		int64(e.Order),
	)
}

func (p *Presenter) ImageSliceFromEntity(input []*entities.Image) []openapi.Image {
	output := make([]openapi.Image, 0, len(input))
	for _, image := range input {
		output = append(output, p.ImageFromEntity(image))
	}
	return output
}

func (p *Presenter) ImageListFromEntity(input []*entities.Image) openapi.ImageList {
	return *openapi.NewImageList(p.ImageSliceFromEntity(input))
}

func (p *Presenter) ImageURLSliceFromEntity(input []*entities.Image) []string {
	output := make([]string, 0, len(input))
	for _, image := range input {
		if image.URL != "" {
			output = append(output, image.URL)
		}
	}
	return output
}

func (p *Presenter) ImageToEntity(id uuid.UUID, image openapi.Image) (*entities.Image, error) {
	// Build entity
	e := &entities.Image{
		ID:    id,
		Order: int(image.GetOrder()),
	}

	// Successful
	return e, nil
}
