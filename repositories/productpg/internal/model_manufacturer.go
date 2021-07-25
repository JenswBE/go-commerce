package internal

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

type Manufacturer struct {
	Base
	Name       string
	WebsiteURL string
	Image      *Image `gorm:"polymorphic:Owner;"`
}

func ManufacturerPgToEntity(m *Manufacturer) *entities.Manufacturer {
	return &entities.Manufacturer{
		ID:         uuid.MustParse(m.ID),
		Name:       m.Name,
		WebsiteURL: m.WebsiteURL,
		Image:      ImagePgToEntity(m.Image),
	}
}

func ManufacturersListPgToEntity(m []*Manufacturer) []*entities.Manufacturer {
	output := make([]*entities.Manufacturer, 0, len(m))
	for _, man := range m {
		output = append(output, ManufacturerPgToEntity(man))
	}
	return output
}

func ManufacturerEntityToPg(e *entities.Manufacturer) *Manufacturer {
	return &Manufacturer{
		Base:       Base{ID: e.ID.String()},
		Name:       e.Name,
		WebsiteURL: e.WebsiteURL,
		Image:      ImageEntityToPg(e.Image),
	}
}
