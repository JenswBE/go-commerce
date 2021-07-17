package entity

import (
	"errors"
)

type Image struct {
	ID        ID
	Extension string // File extension
	URL       string
	Order     int
}

// Validate validates the product data
func (img *Image) Validate() error {
	// Validate simple fields
	if img.Order < 0 {
		return NewError(400, errors.New("image order cannot be negative"))
	}

	// Entity is valid
	return nil
}
