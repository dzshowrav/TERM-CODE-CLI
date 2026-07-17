package llm

import "context"

type Endpoint struct {
	BaseURL      string
	APIKey       string
	ModelID      string
	ProviderName string
}

type Router interface {
	Resolve(ctx context.Context, modelID string) (*Endpoint, error)
}
