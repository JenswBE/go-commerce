package entities

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

type ProductsListTemplate struct {
	BaseData
	Products         []*entities.Product
	ManufacturersMap map[uuid.UUID]entities.Manufacturer
	PublicURLMap     map[uuid.UUID]string
}

func (t ProductsListTemplate) GetTemplateName() string {
	return "productsList"
}

type ProductsFormTemplate struct {
	BaseData
	IsNew   bool
	Product Product
}

func (t ProductsFormTemplate) GetTemplateName() string {
	return "productsForm"
}

type Product struct {
	Name string `form:"name"`
}

func ProductFromEntity(e *entities.ResolvedProduct) Product {
	return Product{
		Name: e.Name,
	}
}

func (e Product) ToEntity() entities.Product {
	return entities.Product{
		Name: e.Name,
	}
}
