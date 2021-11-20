package product

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func ManufacturerFromEntity(p *presenter.Presenter, e *entities.Manufacturer) openapi.Manufacturer {
	// Set basic fields
	m := openapi.NewManufacturer(e.Name)
	m.SetId(p.EncodeID(e.ID))
	m.SetWebsiteUrl(e.WebsiteURL)

	// Set image URL
	if e.Image != nil && len(e.Image.URLs) > 0 {
		m.SetImageUrls(e.Image.URLs)
	}
	return *m
}

func ManufacturerSliceFromEntity(p *presenter.Presenter, input []*entities.Manufacturer) []openapi.Manufacturer {
	output := make([]openapi.Manufacturer, 0, len(input))
	for _, manufacturer := range input {
		output = append(output, ManufacturerFromEntity(p, manufacturer))
	}
	return output
}

func ManufacturerListFromEntity(p *presenter.Presenter, input []*entities.Manufacturer) openapi.ManufacturerList {
	return *openapi.NewManufacturerList(ManufacturerSliceFromEntity(p, input))
}

func ManufacturerToEntity(p *presenter.Presenter, id uuid.UUID, manufacturer openapi.Manufacturer) *entities.Manufacturer {
	// Build entity
	return &entities.Manufacturer{
		ID:         id,
		Name:       manufacturer.GetName(),
		WebsiteURL: manufacturer.GetWebsiteUrl(),
	}
}
