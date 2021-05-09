package shortid

import (
	"github.com/google/uuid"
	"github.com/itchyny/base58-go"
)

// Base58 shortens ID using a Flickr base58 encoding
type Base58 struct {
	encoding *base58.Encoding
}

// NewBase58Service creates a new base58 short ID service
func NewBase58Service() *Base58 {
	return &Base58{
		encoding: base58.FlickrEncoding,
	}
}

// Encode converts an UUID to a short ID
func (b *Base58) Encode(input uuid.UUID) (string, error) {
	encoded, err := b.encoding.Encode(input[:])
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

// Decode converts a short ID to an UUID
func (b *Base58) Decode(input string) (uuid.UUID, error) {
	// Decode short ID
	decoded, err := b.encoding.Decode([]byte(input))
	if err != nil {
		return uuid.Nil, err
	}

	// Convert to UUID
	return uuid.FromBytes(decoded)
}
