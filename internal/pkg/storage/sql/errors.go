package sql

import "errors"

var (
	ErrEntityExists   = errors.New("Entity already exists")
	ErrEntityNotFound = errors.New("Entity not found")
)
