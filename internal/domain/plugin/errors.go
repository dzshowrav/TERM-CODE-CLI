package plugin

import "errors"

var (
	ErrNotFound   = errors.New("plugin not found")
	ErrDuplicate  = errors.New("plugin name already exists")
	ErrNotEnabled = errors.New("plugin is not enabled")
	ErrLoadFailed = errors.New("plugin load failed")
)
