package workspace

import "errors"

var (
	ErrNotFound      = errors.New("workspace not found")
	ErrDuplicatePath = errors.New("workspace path already exists")
)
