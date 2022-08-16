package entities

import (
	"github.com/google/uuid"
)

// ID contains an entity ID, which is a valid UUID.
// ID type is created to make it library agnostic and to add useful helpers like IsNil.
type ID struct {
	id uuid.UUID
}

// NewID create a new entity ID
func NewID() ID {
	return ID{id: uuid.New()}
}

// NewNilID create a new entity ID with value Nil
func NewNilID() ID {
	return ID{id: uuid.Nil}
}

// NewIDFromString convert a string to an entity ID
func NewIDFromString(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID{id: id}, err
}

// NewIDFromString convert a bytes slice to an entity ID
func NewIDFromBytes(b []byte) (ID, error) {
	id, err := uuid.FromBytes(b)
	return ID{id: id}, err
}

// String converts the ID to a string.
// Returns an empty string in case the ID is Nil.
func (id ID) String() string {
	if id.IsNil() {
		return ""
	}
	return id.id.String()
}

// Bytes converts the ID to a byte slice.
// Returns an empty slice in case the ID is Nil.
func (id ID) Bytes() []byte {
	if id.IsNil() {
		return []byte{}
	}
	return id.id[:]
}

func (id ID) IsNil() bool {
	return id.id == uuid.Nil
}
