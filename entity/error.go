package entity

import "errors"

// ErrInvalidEntity indicates the provided entity is invalid
var ErrInvalidEntity = errors.New("invalid entity")

// ErrNotFound indicates the requested entity is not found
var ErrNotFound = errors.New("not found")

// ErrInvalidID indicates the provided id is malformed
var ErrInvalidID = errors.New("invalid id")
