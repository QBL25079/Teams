package core_errors

import "errors"

var (
	ErrInvalidArgument = errors.New("Invalid Argument")
	ErrNotFound        = errors.New("Not found")
	ErrConflict        = errors.New("Conflict")
)
