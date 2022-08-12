package entities

import (
	"github.com/JenswBE/go-commerce/entities"
)

type CategoriesListTemplate struct {
	BaseData
	Categories []*entities.Category
}

func (t CategoriesListTemplate) GetTemplateName() string {
	return "categoriesList"
}

type CategoriesFormTemplate struct {
	BaseData
	IsNew    bool
	Category Category
}

func (t CategoriesFormTemplate) GetTemplateName() string {
	return "categoriesForm"
}

type Category struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}

func CategoryFromEntity(e *entities.Category) Category {
	return Category{
		Name:        e.Name,
		Description: e.Description,
	}
}

func (e Category) ToEntity() entities.Category {
	return entities.Category{
		Name:        e.Name,
		Description: e.Description,
	}
}
