package internal

import (
	"github.com/JenswBE/go-commerce/entities"
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
	Image       *Image `gorm:"polymorphic:Owner;"`
}

func (c *Category) ToEntity() *entities.Category {
	cat := &entities.Category{
		ID:          uuid.MustParse(c.ID),
		Name:        c.Name,
		Description: c.Description,
		ParentID:    uuid.Nil,
		Order:       c.Order,
		ProductIDs:  nil,
		Image:       c.Image.ToEntity(),
	}
	if c.ParentID != nil {
		cat.ParentID = uuid.MustParse(*c.ParentID)
	}
	if len(c.Products) > 0 {
		cat.ProductIDs = make([]uuid.UUID, 0, len(c.Products))
		for _, product := range c.Products {
			cat.ProductIDs = append(cat.ProductIDs, uuid.MustParse(product.ID))
		}
	}
	return cat
}

func CategoriesListPgToEntity(c []*Category) []*entities.Category {
	output := make([]*entities.Category, 0, len(c))
	for _, cat := range c {
		output = append(output, cat.ToEntity())
	}
	return output
}

func CategoryEntityToPg(e *entities.Category) *Category {
	cat := &Category{
		Base:        Base{ID: e.ID.String()},
		Name:        e.Name,
		Description: e.Description,
		ParentID:    nil,
		Children:    nil,
		Products:    nil, // Read-only
		Order:       e.Order,
		Image:       ImageEntityToPg(e.Image),
	}
	if e.ParentID != uuid.Nil {
		id := e.ParentID.String()
		cat.ParentID = &id
	}
	return cat
}
