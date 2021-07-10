package internal

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/google/uuid"
)

type Manufacturer struct {
	Base
	Name       string
	WebsiteURL string
	Image      Image `gorm:"polymorphic:Owner;"`
}

func ManufacturerPgToEntity(m *Manufacturer) *entity.Manufacturer {
	return &entity.Manufacturer{
		ID:         uuid.MustParse(m.ID),
		Name:       m.Name,
		WebsiteURL: m.WebsiteURL,
	}
}

func ManufacturersListPgToEntity(m []*Manufacturer) []*entity.Manufacturer {
	output := make([]*entity.Manufacturer, 0, len(m))
	for _, man := range m {
		output = append(output, ManufacturerPgToEntity(man))
	}
	return output
}

func ManufacturerEntityToPg(e *entity.Manufacturer) *Manufacturer {
	return &Manufacturer{
		Base:       Base{ID: e.ID.String()},
		Name:       e.Name,
		WebsiteURL: e.WebsiteURL,
	}
}
