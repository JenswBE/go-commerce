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

func (id ID) String() string {
	return id.id.String()
}

func (id ID) Bytes() []byte {
	return id.id[:]
}

func (id ID) IsNil() bool {
	return id.id == uuid.Nil
}
