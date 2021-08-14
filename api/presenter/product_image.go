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
		e.URLs,
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

func (p *Presenter) ImageURLsSliceFromEntity(input []*entities.Image) []map[string]string {
	output := make([]map[string]string, 0, len(input))
	for _, image := range input {
		output = append(output, image.URLs)
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
