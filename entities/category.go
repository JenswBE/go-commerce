package entities

import (
	"github.com/JenswBE/go-commerce/api/openapi"
)

// Category data
type Category struct {
	ID          ID
	Name        string
	Description string

	// ID of the parent.
	// uuid.Nil means it's a root category
	ParentID ID

	// Order (priority) of the category.
	// 1 = highest, inf = lowest
	Order int

	ProductIDs []ID

	Image *Image
}

// Validate validates the category data
func (c *Category) Validate() error {
	// Validate simple fields
	if c.Name == "" {
		return NewError(400, openapi.ERRORCODE_CATEGORY_NAME_EMPTY, c.ID.String(), nil)
	}
	if c.Order < 0 {
		return NewError(400, openapi.ERRORCODE_CATEGORY_ORDER_NEGATIVE, c.ID.String(), nil)
	}

	// Entity is valid
	return nil
}
