package provider

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusConnected    Status = "connected"
	StatusConnecting   Status = "connecting"
	StatusDisconnected Status = "disconnected"
	StatusOffline      Status = "offline"
	StatusAuthFailed   Status = "auth_failed"
	StatusTimeout      Status = "timeout"
	StatusUnknown      Status = "unknown"
)

type Provider struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	BaseURL     string    `json:"base_url"`
	APIKey      string    `json:"api_key,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      Status    `json:"status"`
	Latency     int64     `json:"latency_ms"`
	Priority    int       `json:"priority"`
	IsDefault   bool      `json:"is_default"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func New(name, baseURL, apiKey, description string) *Provider {
	now := time.Now()
	return &Provider{
		ID:          uuid.New().String(),
		Name:        name,
		BaseURL:     baseURL,
		APIKey:      apiKey,
		Description: description,
		Status:      StatusDisconnected,
		Priority:    0,
		IsDefault:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
