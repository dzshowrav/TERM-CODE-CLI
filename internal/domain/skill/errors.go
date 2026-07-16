package skill

import "errors"

var (
	ErrNotFound   = errors.New("skill not found")
	ErrDuplicate  = errors.New("skill name already exists")
	ErrNotEnabled = errors.New("skill is not enabled")
)
