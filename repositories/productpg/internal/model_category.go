package internal

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
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
		ID:          generics.Must(entities.NewIDFromString(c.ID)),
		Name:        c.Name,
		Description: c.Description,
		ParentID:    entities.NewNilID(),
		Order:       c.Order,
		ProductIDs:  nil,
		Image:       c.Image.ToEntity(),
	}
	if c.ParentID != nil {
		cat.ParentID = generics.Must(entities.NewIDFromString(*c.ParentID))
	}
	if len(c.Products) > 0 {
		cat.ProductIDs = make([]entities.ID, 0, len(c.Products))
		for _, product := range c.Products {
			cat.ProductIDs = append(cat.ProductIDs, generics.Must(entities.NewIDFromString(product.ID)))
		}
	}
	return cat
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
	if !e.ParentID.IsNil() {
		id := e.ParentID.String()
		cat.ParentID = &id
	}
	return cat
}
