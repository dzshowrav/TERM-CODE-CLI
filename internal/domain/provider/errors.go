package provider

import "errors"

var (
	ErrNotFound      = errors.New("provider not found")
	ErrDuplicateName = errors.New("provider name already exists")
	ErrInvalidURL    = errors.New("invalid provider URL")
	ErrEmptyAPIKey   = errors.New("API key is required for remote providers")
	ErrNoDefault     = errors.New("no default provider set")
	ErrCannotDelete  = errors.New("cannot delete provider with active models")
	ErrConnection    = errors.New("provider connection failed")
)
