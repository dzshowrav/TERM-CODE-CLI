package session

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
	StatusDeleted  Status = "deleted"
)

type Session struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	ProviderID string    `json:"provider_id"`
	ModelID    string    `json:"model_id"`
	AgentID    string    `json:"agent_id,omitempty"`
	Workspace  string    `json:"workspace,omitempty"`
	Status     Status    `json:"status"`
	MessageCnt int       `json:"message_count"`
	TokenIn    int       `json:"tokens_in"`
	TokenOut   int       `json:"tokens_out"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func New(name, providerID, modelID string) *Session {
	now := time.Now()
	return &Session{
		ID:         uuid.New().String(),
		Name:       name,
		ProviderID: providerID,
		ModelID:    modelID,
		Status:     StatusActive,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
