package model

import "errors"

var (
	ErrNotFound      = errors.New("model not found")
	ErrDuplicateID   = errors.New("model ID already exists for provider")
	ErrNotEnabled    = errors.New("model is not enabled")
	ErrNoActiveModel = errors.New("no active model set")
	ErrContextLimit  = errors.New("context window exceeded")
)
