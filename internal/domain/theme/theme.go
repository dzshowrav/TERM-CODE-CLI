package theme

import (
	"time"

	"github.com/google/uuid"
)

type Theme struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author,omitempty"`
	Version   string    `json:"version"`
	IsDark    bool      `json:"is_dark"`
	Palette   string    `json:"palette"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name, palette string, isDark bool) *Theme {
	now := time.Now()
	return &Theme{
		ID:        uuid.New().String(),
		Name:      name,
		Version:   "1.0",
		IsDark:    isDark,
		Palette:   palette,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
