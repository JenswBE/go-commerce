package entity

// Category data
type Category struct {
	ID          ID
	Name        string
	Description string

	// ID of the parent.
	// uuid.Nil means it's a root category
	ParentID ID
}

// NewCategory creates a new category
func NewCategory(name, description string, parentID ID) (*Category, error) {
	b := &Category{
		ID:          NewID(),
		Name:        name,
		Description: description,
		ParentID:    parentID,
	}
	err := b.Validate()
	if err != nil {
		return nil, NewError(400, ErrInvalidEntity)
	}
	return b, nil
}

// Validate validates the category data
func (c *Category) Validate() error {
	// Validate simple fields
	if c.Name == "" {
		return ErrInvalidEntity
	}

	// Entity is valid
	return nil
}
