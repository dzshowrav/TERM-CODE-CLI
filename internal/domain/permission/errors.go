package permission

import "errors"

var (
	ErrNotFound = errors.New("permission entry not found")
	ErrDenied   = errors.New("action not permitted")
	ErrInvalid  = errors.New("invalid permission level")
)
