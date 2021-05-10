package presenter

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/google/uuid"
)

type Manufacturer struct {
	ID      entity.ID `json:"id"`
	ShortID string    `json:"short_id"`
	ManufacturerData
}

type ManufacturerData struct {
	Name       string `json:"name"`
	WebsiteURL string `json:"website_url"`
}

func (p *Presenter) ManufacturerFromEntity(e *entity.Manufacturer) Manufacturer {
	return Manufacturer{
		ID:      e.ID,
		ShortID: p.shortIDService.Encode(e.ID),
		ManufacturerData: ManufacturerData{
			Name:       e.Name,
			WebsiteURL: e.WebsiteURL,
		},
	}
}

func (p *Presenter) ManufacturersListFromEntity(input []*entity.Manufacturer) []Manufacturer {
	output := make([]Manufacturer, 0, len(input))
	for _, manufacturer := range input {
		output = append(output, p.ManufacturerFromEntity(manufacturer))
	}
	return output
}

func (p *Presenter) ManufacturerToEntity(id uuid.UUID, data ManufacturerData) (*entity.Manufacturer, error) {
	// Build entity
	return &entity.Manufacturer{
		ID:         id,
		Name:       data.Name,
		WebsiteURL: data.WebsiteURL,
	}, nil
}
