package agent

import "errors"

var (
	ErrNotFound      = errors.New("agent not found")
	ErrDuplicateName = errors.New("agent name already exists")
)
