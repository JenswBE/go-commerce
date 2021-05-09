package shortid

import "github.com/google/uuid"

//Service interface
type Service interface {
	// Encode converts an UUID to a short ID
	Encode(input uuid.UUID) (string, error)

	// Decode converts a short ID to an UUID
	Decode(input string) (uuid.UUID, error)
}
