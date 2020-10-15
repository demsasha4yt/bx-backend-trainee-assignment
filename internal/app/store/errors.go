package store

import "errors"

var (
	// ErrRecordNotFound return in case there is no rows
	ErrRecordNotFound = errors.New("Record not found")
)
