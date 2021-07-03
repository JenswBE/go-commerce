package internal

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/google/uuid"
)

type Category struct {
	Base
	Name        string
	Description string
	ParentID    *string    `gorm:"type:uuid"`
	Children    []Category `gorm:"foreignkey:ParentID"`
	Products    []*Product `gorm:"many2many:product_categories;"`
	Order       int
}

func CategoryPgToEntity(c *Category) *entity.Category {
	cat := &entity.Category{
		ID:          uuid.MustParse(c.ID),
		Name:        c.Name,
		Description: c.Description,
		ParentID:    uuid.Nil,
		Order:       c.Order,
	}
	if c.ParentID != nil {
		cat.ParentID = uuid.MustParse(*c.ParentID)
	}
	return cat
}

func CategoriesListPgToEntity(c []*Category) []*entity.Category {
	output := make([]*entity.Category, 0, len(c))
	for _, cat := range c {
		output = append(output, CategoryPgToEntity(cat))
	}
	return output
}

func CategoryEntityToPg(e *entity.Category) *Category {
	cat := &Category{
		Base:        Base{ID: e.ID.String()},
		Name:        e.Name,
		Description: e.Description,
		ParentID:    nil,
		Children:    nil,
		Products:    nil,
		Order:       e.Order,
	}
	if e.ParentID != uuid.Nil {
		id := e.ParentID.String()
		cat.ParentID = &id
	}
	return cat
}
