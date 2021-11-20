package product

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func ImageFromEntity(p *presenter.Presenter, e *entities.Image) openapi.Image {
	return *openapi.NewImage(
		p.EncodeID(e.ID),
		e.Extension,
		e.URLs,
		int64(e.Order),
	)
}

func ImageSliceFromEntity(p *presenter.Presenter, input []*entities.Image) []openapi.Image {
	output := make([]openapi.Image, 0, len(input))
	for _, image := range input {
		output = append(output, ImageFromEntity(p, image))
	}
	return output
}

func ImageListFromEntity(p *presenter.Presenter, input []*entities.Image) openapi.ImageList {
	return *openapi.NewImageList(ImageSliceFromEntity(p, input))
}

func ImageURLsSliceFromEntity(p *presenter.Presenter, input []*entities.Image) []map[string]string {
	output := make([]map[string]string, 0, len(input))
	for _, image := range input {
		output = append(output, image.URLs)
	}
	return output
}

func ImageToEntity(p *presenter.Presenter, id uuid.UUID, image openapi.Image) (*entities.Image, error) {
	// Build entity
	e := &entities.Image{
		ID:    id,
		Order: int(image.GetOrder()),
	}

	// Successful
	return e, nil
}
