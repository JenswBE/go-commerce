package internal

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

type Manufacturer struct {
	Base
	Name       string
	WebsiteURL string
	Image      *Image `gorm:"polymorphic:Owner;"`
}

func (m *Manufacturer) ToEntity() *entities.Manufacturer {
	return &entities.Manufacturer{
		ID:         generics.Must(entities.NewIDFromString(m.ID)),
		Name:       m.Name,
		WebsiteURL: m.WebsiteURL,
		Image:      m.Image.ToEntity(),
	}
}

func ManufacturerEntityToPg(e *entities.Manufacturer) *Manufacturer {
	return &Manufacturer{
		Base:       Base{ID: e.ID.String()},
		Name:       e.Name,
		WebsiteURL: e.WebsiteURL,
		Image:      ImageEntityToPg(e.Image),
	}
}
