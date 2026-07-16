package config

import "time"

type Config struct {
	DefaultProvider  string    `json:"default_provider"`
	DefaultModel     string    `json:"default_model"`
	DefaultAgent     string    `json:"default_agent"`
	DefaultWorkspace string    `json:"default_workspace"`
	Temperature      float64   `json:"temperature"`
	MaxTokens        int       `json:"max_tokens"`
	Streaming        bool      `json:"streaming"`
	Theme            string    `json:"theme"`
	CompactMode      bool      `json:"compact_mode"`
	ShowTokenGauge   bool      `json:"show_token_gauge"`
	AutoSave         bool      `json:"auto_save"`
	AutoSaveInterval int       `json:"auto_save_interval"`
	EncryptKeys      bool      `json:"encrypt_keys"`
	DefaultPerm      string    `json:"default_permission"`
	Timeout          int       `json:"timeout"`
	RetryCount       int       `json:"retry_count"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func Defaults() *Config {
	return &Config{
		Temperature:      0.7,
		MaxTokens:        4096,
		Streaming:        true,
		CompactMode:      false,
		ShowTokenGauge:   true,
		AutoSave:         true,
		AutoSaveInterval: 30,
		EncryptKeys:      true,
		DefaultPerm:      "ask",
		Timeout:          60,
		RetryCount:       3,
		UpdatedAt:        time.Now(),
	}
}
