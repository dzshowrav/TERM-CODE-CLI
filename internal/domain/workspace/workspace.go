package workspace

import (
	"time"

	"github.com/google/uuid"
)

type Workspace struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	IsDefault bool      `json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name, path string) *Workspace {
	now := time.Now()
	return &Workspace{
		ID:        uuid.New().String(),
		Name:      name,
		Path:      path,
		IsDefault: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
