package customerrors

import "errors"

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrEmptyID          = errors.New("id param is empty")
)
