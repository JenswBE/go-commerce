package entities

import (
	"errors"
	"fmt"
)

// GoComError allows to bundle a status with the original error.
// This allows to fine-grained response codes at the API level.
type GoComError struct {
	// HTTP status code
	Status int

	// Original error
	Err error
}

func (e *GoComError) Error() string {
	return fmt.Sprintf("%d - %s", e.Status, e.Err.Error())
}

// NewError returns a new GoComError
func NewError(status int, err error) error {
	return &GoComError{
		Status: status,
		Err:    err,
	}
}

// ErrInvalidEntity indicates the provided entity is invalid
var ErrInvalidEntity = errors.New("invalid entity")

// ErrNotFound indicates the requested entity is not found
var ErrNotFound = errors.New("not found")

// ErrInvalidID indicates the provided id is malformed
var ErrInvalidID = errors.New("invalid id")
