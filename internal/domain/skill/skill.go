package skill

import (
	"time"

	"github.com/google/uuid"
)

type Category string

const (
	CategoryGeneral  Category = "general"
	CategoryCoding   Category = "coding"
	CategoryReview   Category = "review"
	CategoryTesting  Category = "testing"
	CategorySecurity Category = "security"
	CategoryCustom   Category = "custom"
)

type Skill struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Category    Category  `json:"category"`
	Version     string    `json:"version"`
	Path        string    `json:"path,omitempty"`
	Enabled     bool      `json:"enabled"`
	IsBuiltin   bool      `json:"is_builtin"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func New(name, description string, category Category) *Skill {
	now := time.Now()
	return &Skill{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Category:    category,
		Version:     "1.0",
		Enabled:     true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
