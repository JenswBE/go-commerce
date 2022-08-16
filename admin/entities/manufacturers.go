package entities

import (
	"github.com/JenswBE/go-commerce/entities"
)

type ManufacturersListTemplate struct {
	BaseData
	Manufacturers []*entities.Manufacturer
}

func (t ManufacturersListTemplate) GetTemplateName() string {
	return "manufacturersList"
}

type ManufacturersFormTemplate struct {
	BaseData
	IsNew        bool
	Manufacturer Manufacturer
}

func (t ManufacturersFormTemplate) GetTemplateName() string {
	return "manufacturersForm"
}

type ManufacturersImageTemplate struct {
	BaseData
	Manufacturer entities.Manufacturer
}

func (t ManufacturersImageTemplate) GetTemplateName() string {
	return "manufacturersImage"
}

type Manufacturer struct {
	Name       string `form:"name"`
	WebsiteURL string `form:"website_url"`
}

func ManufacturerFromEntity(e *entities.Manufacturer) Manufacturer {
	return Manufacturer{
		Name:       e.Name,
		WebsiteURL: e.WebsiteURL,
	}
}

func (e Manufacturer) ToEntity() entities.Manufacturer {
	return entities.Manufacturer{
		Name:       e.Name,
		WebsiteURL: e.WebsiteURL,
	}
}
