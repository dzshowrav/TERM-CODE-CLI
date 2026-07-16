package tool

import "time"

type Status string

const (
	StatusPending  Status = "pending"
	StatusRunning  Status = "running"
	StatusApproved Status = "approved"
	StatusDenied   Status = "denied"
	StatusSuccess  Status = "success"
	StatusFailed   Status = "failed"
)

type Tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema any    `json:"input_schema"`
}

type Result struct {
	Tool     string    `json:"tool"`
	Input    string    `json:"input"`
	Output   string    `json:"output"`
	Error    string    `json:"error,omitempty"`
	Status   Status    `json:"status"`
	Duration int64     `json:"duration_ms"`
	Time     time.Time `json:"time"`
}
