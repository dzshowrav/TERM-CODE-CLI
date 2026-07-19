package model

import (
	"time"

	"github.com/google/uuid"
)

type Category string

const (
	CategoryGeneral      Category = "general"
	CategoryCoding       Category = "coding"
	CategoryReasoning    Category = "reasoning"
	CategoryVision       Category = "vision"
	CategoryEmbedding    Category = "embedding"
	CategoryAudio        Category = "audio"
	CategoryExperimental Category = "experimental"
	CategoryCustom       Category = "custom"
)

type Capabilities struct {
	Streaming       bool `json:"streaming"`
	ToolCalling     bool `json:"tool_calling"`
	Reasoning       bool `json:"reasoning"`
	Vision          bool `json:"vision"`
	Embeddings      bool `json:"embeddings"`
	JSONMode        bool `json:"json_mode"`
	FunctionCalling bool `json:"function_calling"`
	SystemPrompt    bool `json:"system_prompt"`
}

type Model struct {
	ID           string       `json:"id"`
	ProviderID   string       `json:"provider_id"`
	ModelID      string       `json:"model_id"`
	DisplayName  string       `json:"display_name"`
	Description  string       `json:"description,omitempty"`
	Category     Category     `json:"category"`
	Capabilities Capabilities `json:"capabilities"`
	MaxContext   int          `json:"max_context"`
	MaxOutput    int          `json:"max_output"`
	PricingInput float64      `json:"pricing_input"`
	PricingOut   float64      `json:"pricing_output"`
	IsFavorite   bool         `json:"is_favorite"`
	Enabled      bool         `json:"enabled"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

func New(providerID, modelID, displayName string, category Category) *Model {
	now := time.Now()
	return &Model{
		ID:          uuid.New().String(),
		ProviderID:  providerID,
		ModelID:     modelID,
		DisplayName: displayName,
		Category:    category,
		Capabilities: Capabilities{
			Streaming:    true,
			SystemPrompt: true,
		},
		MaxContext: 4096,
		MaxOutput:  4096,
		Enabled:    true,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
