package shortid

import (
	"github.com/JenswBE/go-commerce/entities"
)

//Service interface
type Service interface {
	// Encode converts an UUID to a short ID
	Encode(input entities.ID) string

	// Decode converts a short ID to an UUID
	Decode(input string) (entities.ID, error)
}
