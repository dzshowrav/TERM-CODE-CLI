package conversation

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusActive Status = "active"
	StatusDone   Status = "done"
	StatusStale  Status = "stale"
)

type Conversation struct {
	ID        string    `json:"id"`
	SessionID string    `json:"session_id"`
	Summary   string    `json:"summary,omitempty"`
	Status    Status    `json:"status"`
	MsgCount  int       `json:"message_count"`
	TokenIn   int       `json:"tokens_in"`
	TokenOut  int       `json:"tokens_out"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(sessionID string) *Conversation {
	now := time.Now()
	return &Conversation{
		ID:        uuid.New().String(),
		SessionID: sessionID,
		Status:    StatusActive,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (c *Conversation) AddTokens(in, out int) {
	c.TokenIn += in
	c.TokenOut += out
	c.MsgCount++
	c.UpdatedAt = time.Now()
}

func (c *Conversation) IsEmpty() bool {
	return c.MsgCount == 0
}
