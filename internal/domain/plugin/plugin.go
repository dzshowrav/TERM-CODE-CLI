package plugin

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusError    Status = "error"
)

type Plugin struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Author      string    `json:"author,omitempty"`
	Description string    `json:"description,omitempty"`
	Path        string    `json:"path,omitempty"`
	FilePath    string    `json:"file_path,omitempty"`
	EntryPoint  string    `json:"entry_point,omitempty"`
	Status      Status    `json:"status"`
	Enabled     bool      `json:"enabled"`
	CreatedAt   time.Time `json:"created_at"`
}

func New(name, version, author, description string) *Plugin {
	return &Plugin{
		ID:          uuid.New().String(),
		Name:        name,
		Version:     version,
		Author:      author,
		Description: description,
		Status:      StatusInactive,
		Enabled:     true,
		CreatedAt:   time.Now(),
	}
}
