package internal

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

type Service struct {
	Base
	Name              string
	Description       string
	Price             int
	ServiceCategoryID string `gorm:"type:uuid;uniqueIndex:idx_service_category_order,priority:1"`
	Order             int    `gorm:"uniqueIndex:idx_service_category_order,priority:2"`
}

func (s *Service) ToEntity() *entities.Service {
	return &entities.Service{
		ID:                generics.Must(entities.NewIDFromString(s.ID)),
		Name:              s.Name,
		Description:       s.Description,
		Price:             entities.NewAmountInCents(s.Price),
		Order:             s.Order,
		ServiceCategoryID: generics.Must(entities.NewIDFromString(s.ServiceCategoryID)),
	}
}

func ServiceEntityToPg(e *entities.Service) *Service {
	return &Service{
		Base:              Base{ID: e.ID.String()},
		Name:              e.Name,
		Description:       e.Description,
		Price:             e.Price.Int(),
		Order:             e.Order,
		ServiceCategoryID: e.ServiceCategoryID.String(),
	}
}
