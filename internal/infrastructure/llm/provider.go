package llm

import (
	"context"
	"fmt"
	"strings"

	"termcode/pkg/apitypes"
)

type ProviderType string

const (
	ProviderOpenAI    ProviderType = "openai"
	ProviderOllama    ProviderType = "ollama"
	ProviderAnthropic ProviderType = "anthropic"
	ProviderGoogle    ProviderType = "google"
	ProviderCustom    ProviderType = "custom"
)

type LLMProvider interface {
	Chat(ctx context.Context, req *apitypes.ChatRequest) (*apitypes.ChatResponse, error)
	ChatStream(ctx context.Context, req *apitypes.ChatRequest) (<-chan apitypes.StreamChunk, <-chan error)
}

func NewProvider(pt ProviderType, baseURL, apiKey string) (LLMProvider, error) {
	switch pt {
	case ProviderOpenAI:
		return NewOpenAIAdapter(baseURL, apiKey), nil
	case ProviderOllama:
		return NewOllamaAdapter(baseURL), nil
	case ProviderAnthropic:
		return NewAnthropicAdapter(baseURL, apiKey), nil
	case ProviderGoogle:
		return NewOpenAIAdapter(baseURL, apiKey), nil
	case ProviderCustom:
		return NewOpenAIAdapter(baseURL, apiKey), nil
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", pt)
	}
}

func DetectProviderType(baseURL string) ProviderType {
	u := strings.ToLower(baseURL)
	switch {
	case strings.Contains(u, "ollama"):
		return ProviderOllama
	case strings.Contains(u, "anthropic"):
		return ProviderAnthropic
	case strings.Contains(u, "googleapis") || strings.Contains(u, "generativelanguage"):
		return ProviderGoogle
	case strings.Contains(u, "openai"):
		return ProviderOpenAI
	default:
		return ProviderCustom
	}
}
