package config

import "errors"

var (
	ErrNotFound = errors.New("config key not found")
	ErrInvalid  = errors.New("invalid config value")
)
