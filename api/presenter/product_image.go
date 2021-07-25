package presenter

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func (p *Presenter) ImageFromEntity(e *entities.Image) openapi.Image {
	m := openapi.NewImage(int64(e.Order))
	m.SetId(p.EncodeID(e.ID))
	m.SetExt(e.Extension)
	m.SetUrl(e.URL)
	return *m
}

func (p *Presenter) ImagesListFromEntity(input []*entities.Image) []openapi.Image {
	output := make([]openapi.Image, 0, len(input))
	for _, image := range input {
		output = append(output, p.ImageFromEntity(image))
	}
	return output
}

func (p *Presenter) ImageURLListFromEntity(input []*entities.Image) []string {
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
