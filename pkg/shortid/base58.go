package shortid

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

// Base58 shortens ID using a Flickr base58 encoding
type Base58 struct{}

// NewBase58Service creates a new base58 short ID service
func NewBase58Service() *Base58 {
	return &Base58{}
}

// Encode converts an UUID to a short ID
func (b *Base58) Encode(input uuid.UUID) string {
	encoded := base58.Encode(input[:])
	return string(encoded)
}

// Decode converts a short ID to an UUID
func (b *Base58) Decode(input string) (uuid.UUID, error) {
	decoded := base58.Decode(input)
	return uuid.FromBytes(decoded)
}
