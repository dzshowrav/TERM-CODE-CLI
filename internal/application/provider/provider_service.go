package provider

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"termcode/internal/domain/provider"
	"termcode/pkg/helpers"
)

type Repository interface {
	Create(ctx context.Context, p *provider.Provider) error
	GetByID(ctx context.Context, id string) (*provider.Provider, error)
	GetByName(ctx context.Context, name string) (*provider.Provider, error)
	List(ctx context.Context) ([]*provider.Provider, error)
	Update(ctx context.Context, p *provider.Provider) error
	Delete(ctx context.Context, id string) error
	SetDefault(ctx context.Context, id string) error
	GetDefault(ctx context.Context) (*provider.Provider, error)
}

type Service struct {
	repo   Repository
	client *http.Client
	logger *slog.Logger
}

func NewService(repo Repository, logger *slog.Logger) *Service {
	return &Service{
		repo:   repo,
		client: &http.Client{Timeout: 10 * time.Second},
		logger: logger.With("svc", "provider"),
	}
}

func (s *Service) Create(ctx context.Context, name, baseURL, apiKey, desc string) (*provider.Provider, error) {
	name = strings.TrimSpace(name)
	baseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/")

	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if baseURL == "" {
		return nil, fmt.Errorf("base URL is required")
	}
	if !strings.HasPrefix(baseURL, "http") {
		return nil, provider.ErrInvalidURL
	}

	existing, err := s.repo.GetByName(ctx, name)
	if err == nil && existing != nil {
		return nil, provider.ErrDuplicateName
	}

	var encryptedKey string
	if apiKey != "" {
		encryptedKey, err = helpers.EncryptAPIKey(apiKey)
		if err != nil {
			return nil, fmt.Errorf("encrypt key: %w", err)
		}
	}

	p := provider.New(name, baseURL, encryptedKey, desc)
	if err := s.repo.Create(ctx, p); err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	s.logger.Info("provider created", "name", name, "id", p.ID)
	return p, nil
}

func (s *Service) Update(ctx context.Context, id, name, baseURL, apiKey, desc string) (*provider.Provider, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		p.Name = strings.TrimSpace(name)
	}
	if baseURL != "" {
		p.BaseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/")
	}
	if desc != "" {
		p.Description = desc
	}
	if apiKey != "" {
		encrypted, err := helpers.EncryptAPIKey(apiKey)
		if err != nil {
			return nil, fmt.Errorf("encrypt key: %w", err)
		}
		p.APIKey = encrypted
	}

	if err := s.repo.Update(ctx, p); err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}

	s.logger.Info("provider updated", "name", p.Name, "id", p.ID)
	return p, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	s.logger.Info("provider deleted", "id", id)
	return nil
}

func (s *Service) List(ctx context.Context) ([]*provider.Provider, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (*provider.Provider, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetDefault(ctx context.Context) (*provider.Provider, error) {
	return s.repo.GetDefault(ctx)
}

func (s *Service) SetDefault(ctx context.Context, id string) error {
	return s.repo.SetDefault(ctx, id)
}

type ConnectionTestResult struct {
	Success  bool   `json:"success"`
	Latency  int64  `json:"latency_ms"`
	Message  string `json:"message"`
	Provider string `json:"provider"`
	Models   int    `json:"models_count"`
}

func (s *Service) TestConnection(ctx context.Context, id string) (*ConnectionTestResult, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	decryptedKey, err := helpers.DecryptAPIKey(p.APIKey)
	if err != nil {
		return nil, fmt.Errorf("decrypt key: %w", err)
	}

	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.BaseURL+"/v1/models", nil)
	if err != nil {
		return &ConnectionTestResult{Success: false, Message: "invalid request"}, nil
	}

	if decryptedKey != "" {
		req.Header.Set("Authorization", "Bearer "+decryptedKey)
	}

	resp, err := s.client.Do(req)
	latency := time.Since(start).Milliseconds()
	if err != nil {
		return &ConnectionTestResult{
			Success:  false,
			Latency:  latency,
			Message:  fmt.Sprintf("connection failed: %v", err),
			Provider: p.Name,
		}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return &ConnectionTestResult{
			Success:  false,
			Latency:  latency,
			Message:  fmt.Sprintf("HTTP %d: %s", resp.StatusCode, resp.Status),
			Provider: p.Name,
		}, nil
	}

	p.Status = provider.StatusConnected
	p.Latency = latency
	s.repo.Update(ctx, p)

	return &ConnectionTestResult{
		Success:  true,
		Latency:  latency,
		Message:  "connected successfully",
		Provider: p.Name,
	}, nil
}

func (s *Service) DecryptAPIKey(encrypted string) (string, error) {
	return helpers.DecryptAPIKey(encrypted)
}

func (s *Service) MaskAPIKey(key string) string {
	return helpers.MaskAPIKey(key)
}
