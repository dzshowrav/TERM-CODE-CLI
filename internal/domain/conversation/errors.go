package conversation

import "errors"

var (
	ErrNotFound = errors.New("conversation not found")
	ErrEmpty    = errors.New("conversation is empty")
)
