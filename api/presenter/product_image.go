package presenter

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entity"
	"github.com/google/uuid"
)

func (p *Presenter) ImageFromEntity(e *entity.Image) openapi.Image {
	m := openapi.NewImage()
	m.SetId(p.EncodeID(e.ID))
	m.SetExt(e.Extension)
	m.SetUrl(e.URL)
	m.SetOrder(int64(e.Order))
	return *m
}

func (p *Presenter) ImagesListFromEntity(input []*entity.Image) []openapi.Image {
	output := make([]openapi.Image, 0, len(input))
	for _, image := range input {
		output = append(output, p.ImageFromEntity(image))
	}
	return output
}

func (p *Presenter) ImageToEntity(id uuid.UUID, image openapi.Image) (*entity.Image, error) {
	// Build entity
	e := &entity.Image{
		ID:    id,
		Order: int(image.GetOrder()),
	}

	// Successful
	return e, nil
}
