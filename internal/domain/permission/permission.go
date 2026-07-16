package permission

import "time"

type Level string

const (
	LevelAlwaysAllow Level = "always_allow"
	LevelAllowOnce   Level = "allow_once"
	LevelAsk         Level = "ask"
	LevelDeny        Level = "deny"
)

type Entry struct {
	ToolName   string    `json:"tool_name"`
	Permission Level     `json:"permission"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func New(toolName string, level Level) *Entry {
	return &Entry{
		ToolName:   toolName,
		Permission: level,
		UpdatedAt:  time.Now(),
	}
}

func (e *Entry) IsAllowed() bool {
	return e.Permission == LevelAlwaysAllow || e.Permission == LevelAllowOnce
}

func (e *Entry) IsDenied() bool {
	return e.Permission == LevelDeny
}
