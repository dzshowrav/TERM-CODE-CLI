package theme

import "errors"

var (
	ErrNotFound  = errors.New("theme not found")
	ErrDuplicate = errors.New("theme name already exists")
)
