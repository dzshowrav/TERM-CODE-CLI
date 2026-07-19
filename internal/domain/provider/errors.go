package provider

import "errors"

var (
	ErrNotFound      = errors.New("provider not found")
	ErrDuplicateName = errors.New("provider name already exists")
	ErrInvalidURL    = errors.New("invalid provider URL")
	ErrNoDefault     = errors.New("no default provider set")
)
