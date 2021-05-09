package shortid

import (
	"github.com/google/uuid"
)

// Fake returns the input unchanged
type Fake struct{}

// NewFakeService creates a new fake short ID service
func NewFakeService() *Fake {
	return &Fake{}
}

// Encode converts an UUID to a short ID
func (f *Fake) Encode(input uuid.UUID) (string, error) {
	return input.String(), nil
}

// Decode converts a short ID to an UUID
func (f *Fake) Decode(input string) (uuid.UUID, error) {
	return uuid.Parse(input)
}
