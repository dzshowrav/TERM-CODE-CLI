package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"termcode/internal/domain/model"
	"termcode/internal/domain/provider"
)

type DiscoveryService struct {
	client    *http.Client
	modelRepo ModelRepository
}

type ModelRepository interface {
	Create(ctx context.Context, m *model.Model) error
	DeleteByProvider(ctx context.Context, providerID string) error
}

func NewDiscoveryService(modelRepo ModelRepository) *DiscoveryService {
	return &DiscoveryService{
		client:    &http.Client{Timeout: 30 * time.Second},
		modelRepo: modelRepo,
	}
}

type openAIModel struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	OwnedBy string `json:"owned_by"`
}

type openAIModelsResponse struct {
	Object string        `json:"object"`
	Data   []openAIModel `json:"data"`
}

func normalizeBaseURL(raw string) string {
	u := strings.TrimRight(raw, "/")
	if strings.HasSuffix(u, "/v1") {
		u = u[:len(u)-3]
	}
	return u
}

func (s *DiscoveryService) DiscoverFromProvider(ctx context.Context, p *provider.Provider, apiKey string) ([]*model.Model, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, normalizeBaseURL(p.BaseURL)+"/v1/models", nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	if apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		msg := strings.TrimSpace(string(body))
		if len(msg) > 300 {
			msg = msg[:300] + "..."
		}
		if strings.Contains(msg, "<html") || strings.Contains(msg, "<!DOCTYPE") {
			msg = "server returned non-JSON response (HTML)"
		}
		return nil, fmt.Errorf("API error (HTTP %d): %s", resp.StatusCode, msg)
	}

	var modelsResp openAIModelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&modelsResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	var models []*model.Model
	for _, apiModel := range modelsResp.Data {
		cat := classifyModel(apiModel.ID)
		m := model.New(p.ID, apiModel.ID, apiModel.ID, cat)
		m.MaxContext = defaultContext(cat)
		models = append(models, m)
	}

	return models, nil
}

func (s *DiscoveryService) SyncFromProvider(ctx context.Context, p *provider.Provider, apiKey string) (int, error) {
	models, err := s.DiscoverFromProvider(ctx, p, apiKey)
	if err != nil {
		return 0, err
	}

	if err := s.modelRepo.DeleteByProvider(ctx, p.ID); err != nil {
		return 0, fmt.Errorf("clear models: %w", err)
	}

	for _, m := range models {
		if err := s.modelRepo.Create(ctx, m); err != nil {
			return 0, fmt.Errorf("save model %s: %w", m.ModelID, err)
		}
	}

	return len(models), nil
}

func (s *DiscoveryService) discoverAndSync(ctx context.Context, providerID, baseURL, apiKey string) (int, error) {
	p := &provider.Provider{
		ID:      providerID,
		BaseURL: baseURL,
	}
	return s.SyncFromProvider(ctx, p, apiKey)
}

func classifyModel(id string) model.Category {
	id = strings.ToLower(id)

	switch {
	case strings.Contains(id, "gpt-4") || strings.Contains(id, "claude-3") || strings.Contains(id, "gemini"):
		if strings.Contains(id, "vision") || strings.Contains(id, "v1") && strings.Contains(id, "vision") {
			return model.CategoryVision
		}
		return model.CategoryGeneral

	case strings.Contains(id, "gpt-3.5"):
		return model.CategoryGeneral

	case strings.Contains(id, "embed") || strings.Contains(id, "ada"):
		return model.CategoryEmbedding

	case strings.Contains(id, "dall-e") || strings.Contains(id, "image"):
		return model.CategoryVision

	case strings.Contains(id, "whisper") || strings.Contains(id, "tts"):
		return model.CategoryAudio

	case strings.Contains(id, "code") || strings.Contains(id, "coder") || strings.Contains(id, "deepseek"):
		return model.CategoryCoding

	case strings.Contains(id, "reason") || strings.Contains(id, "o1") || strings.Contains(id, "o3"):
		return model.CategoryReasoning

	default:
		return model.CategoryGeneral
	}
}

func defaultContext(cat model.Category) int {
	switch cat {
	case model.CategoryReasoning:
		return 128000
	case model.CategoryCoding:
		return 64000
	case model.CategoryVision:
		return 64000
	default:
		return 4096
	}
}
