package entities

import (
	"fmt"

	"github.com/JenswBE/go-commerce/entities"
)

type ServiceCategoriesListTemplate struct {
	BaseData
	ServiceCategories []*entities.ServiceCategory
}

func (t ServiceCategoriesListTemplate) GetTemplateName() string {
	return "serviceCategoriesList"
}

type ServiceCategoriesFormTemplate struct {
	BaseData
	IsNew           bool
	ServiceCategory ServiceCategory
}

func (t ServiceCategoriesFormTemplate) GetTemplateName() string {
	return "serviceCategoriesForm"
}

type ServiceCategory struct {
	Name string `form:"name"`
}

func ServiceCategoryFromEntity(e *entities.ServiceCategory) ServiceCategory {
	return ServiceCategory{Name: e.Name}
}

func (e ServiceCategory) ToEntity() entities.ServiceCategory {
	return entities.ServiceCategory{Name: e.Name}
}

type ServicesListTemplate struct {
	BaseData
	Services        []*entities.Service
	ServiceCategory *entities.ServiceCategory
}

func (t ServicesListTemplate) GetTemplateName() string {
	return "servicesList"
}

type ServicesFormTemplate struct {
	BaseData
	IsNew             bool
	Service           Service
	ServiceCategories []*entities.ServiceCategory
}

func (t ServicesFormTemplate) GetTemplateName() string {
	return "servicesForm"
}

type Service struct {
	Name              string `form:"name"`
	Description       string `form:"description"`
	Price             string `form:"price"`
	ServiceCategoryID string `form:"service_category_id"`
}

func ServiceFromEntity(e *entities.Service) Service {
	return Service{
		Name:              e.Name,
		Description:       e.Description,
		Price:             e.Price.String(),
		ServiceCategoryID: e.ServiceCategoryID.String(),
	}
}

func (e Service) ToEntity() (entities.Service, error) {
	// Parse price
	price, err := entities.NewAmountInCentsFromString(e.Price)
	if err != nil {
		return entities.Service{}, fmt.Errorf("failed to convert price string to AmountInCents: %w", err)
	}

	// Parse service category ID
	serviceCategoryID := entities.NewNilID()
	if e.ServiceCategoryID != "" {
		serviceCategoryID, err = entities.NewIDFromString(e.ServiceCategoryID)
		if err != nil {
			return entities.Service{}, fmt.Errorf("failed to parse service category ID to a UUID: %w", err)
		}
	}

	return entities.Service{
		Name:              e.Name,
		Description:       e.Description,
		Price:             price,
		ServiceCategoryID: serviceCategoryID,
	}, nil
}
