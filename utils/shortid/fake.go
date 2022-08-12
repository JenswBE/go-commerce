package shortid

import (
	"github.com/JenswBE/go-commerce/entities"
)

// Fake returns the input unchanged
type Fake struct{}

// NewFakeService creates a new fake short ID service
func NewFakeService() *Fake {
	return &Fake{}
}

// Encode converts an UUID to a short ID
func (f *Fake) Encode(input entities.ID) string {
	return input.String()
}

// Decode converts a short ID to an UUID
func (f *Fake) Decode(input string) (entities.ID, error) {
	return entities.NewIDFromString(input)
}
