package shortid

import (
	"github.com/btcsuite/btcutil/base58"

	"github.com/JenswBE/go-commerce/entities"
)

// Base58 shortens ID using a Flickr base58 encoding
type Base58 struct{}

// NewBase58Service creates a new base58 short ID service
func NewBase58Service() *Base58 {
	return &Base58{}
}

// Encode converts an UUID to a short ID
func (b *Base58) Encode(input entities.ID) string {
	return base58.Encode(input.Bytes())
}

// Decode converts a short ID to an UUID
func (b *Base58) Decode(input string) (entities.ID, error) {
	decoded := base58.Decode(input)
	return entities.NewIDFromBytes(decoded)
}
