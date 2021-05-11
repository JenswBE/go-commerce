package entity

// Category data
type Category struct {
	ID          ID
	Name        string
	Description string
}

// NewCategory creates a new category
func NewCategory(name, description string) (*Category, error) {
	b := &Category{
		Name: name,
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
