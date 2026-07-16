package agent

import (
	"time"

	"github.com/google/uuid"
)

type Agent struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	SystemPrompt string    `json:"system_prompt,omitempty"`
	ModelID      string    `json:"model_id,omitempty"`
	Tools        []string  `json:"tools,omitempty"`
	IsDefault    bool      `json:"is_default"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func New(name, description, systemPrompt string) *Agent {
	now := time.Now()
	return &Agent{
		ID:           uuid.New().String(),
		Name:         name,
		Description:  description,
		SystemPrompt: systemPrompt,
		IsDefault:    false,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
