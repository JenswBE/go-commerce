package entity

import "errors"

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
}

// Validate validates the category data
func (c *Category) Validate() error {
	// Validate simple fields
	if c.Name == "" {
		return NewError(400, errors.New("category name cannot be empty"))
	}
	if c.Order < 0 {
		return NewError(400, errors.New("category order cannot be negative"))
	}

	// Entity is valid
	return nil
}
