package presenter

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entity"
	"github.com/google/uuid"
)

func (p *Presenter) ManufacturerFromEntity(e *entity.Manufacturer) openapi.Manufacturer {
	m := openapi.NewManufacturer()
	m.SetId(p.shortIDService.Encode(e.ID))
	m.SetName(e.Name)
	m.SetWebsiteUrl(e.WebsiteURL)
	return *m
}

func (p *Presenter) ManufacturersListFromEntity(input []*entity.Manufacturer) []openapi.Manufacturer {
	output := make([]openapi.Manufacturer, 0, len(input))
	for _, manufacturer := range input {
		output = append(output, p.ManufacturerFromEntity(manufacturer))
	}
	return output
}

func (p *Presenter) ManufacturerToEntity(id uuid.UUID, manufacturer openapi.Manufacturer) *entity.Manufacturer {
	// Build entity
	return &entity.Manufacturer{
		ID:         id,
		Name:       manufacturer.GetName(),
		WebsiteURL: manufacturer.GetWebsiteUrl(),
	}
}
