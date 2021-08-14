package presenter

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func (p *Presenter) ManufacturerFromEntity(e *entities.Manufacturer) openapi.Manufacturer {
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

func (p *Presenter) ManufacturerSliceFromEntity(input []*entities.Manufacturer) []openapi.Manufacturer {
	output := make([]openapi.Manufacturer, 0, len(input))
	for _, manufacturer := range input {
		output = append(output, p.ManufacturerFromEntity(manufacturer))
	}
	return output
}

func (p *Presenter) ManufacturerListFromEntity(input []*entities.Manufacturer) openapi.ManufacturerList {
	return *openapi.NewManufacturerList(p.ManufacturerSliceFromEntity(input))
}

func (p *Presenter) ManufacturerToEntity(id uuid.UUID, manufacturer openapi.Manufacturer) *entities.Manufacturer {
	// Build entity
	return &entities.Manufacturer{
		ID:         id,
		Name:       manufacturer.GetName(),
		WebsiteURL: manufacturer.GetWebsiteUrl(),
	}
}
