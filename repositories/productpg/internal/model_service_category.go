package internal

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

type ServiceCategory struct {
	Base
	Name     string
	Order    int
	Services []Service
}

func (c *ServiceCategory) ToEntity() *entities.ServiceCategory {
	cat := &entities.ServiceCategory{
		ID:    generics.Must(entities.NewIDFromString(c.ID)),
		Name:  c.Name,
		Order: c.Order,
	}
	if len(c.Services) > 0 {
		cat.ServiceIDs = make([]entities.ID, 0, len(c.Services))
		for _, svc := range c.Services {
			cat.ServiceIDs = append(cat.ServiceIDs, generics.Must(entities.NewIDFromString(svc.ID)))
		}
	}
	return cat
}

func ServiceCategoriesListPgToEntity(c []*ServiceCategory) []*entities.ServiceCategory {
	output := make([]*entities.ServiceCategory, 0, len(c))
	for _, cat := range c {
		output = append(output, cat.ToEntity())
	}
	return output
}

func ServiceCategoryEntityToPg(e *entities.ServiceCategory) *ServiceCategory {
	cat := &ServiceCategory{
		Base:     Base{ID: e.ID.String()},
		Name:     e.Name,
		Services: nil,
		Order:    e.Order,
	}
	return cat
}
