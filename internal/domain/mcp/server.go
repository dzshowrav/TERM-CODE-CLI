package mcp

import (
	"time"

	"github.com/google/uuid"
)

type Transport string

const (
	TransportStdio     Transport = "stdio"
	TransportSSE       Transport = "sse"
	TransportWebSocket Transport = "websocket"
)

type Status string

const (
	StatusConnected    Status = "connected"
	StatusConnecting   Status = "connecting"
	StatusDisconnected Status = "disconnected"
	StatusError        Status = "error"
)

type Server struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Transport Transport `json:"transport"`
	Command   string    `json:"command,omitempty"`
	Args      []string  `json:"args,omitempty"`
	URL       string    `json:"url,omitempty"`
	Env       []string  `json:"env,omitempty"`
	Status    Status    `json:"status"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewStdio(name, command string, args []string) *Server {
	now := time.Now()
	return &Server{
		ID:        uuid.New().String(),
		Name:      name,
		Transport: TransportStdio,
		Command:   command,
		Args:      args,
		Status:    StatusDisconnected,
		Enabled:   true,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewSSE(name, url string) *Server {
	now := time.Now()
	return &Server{
		ID:        uuid.New().String(),
		Name:      name,
		Transport: TransportSSE,
		URL:       url,
		Status:    StatusDisconnected,
		Enabled:   true,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
