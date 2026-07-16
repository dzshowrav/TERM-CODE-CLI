package llm

import (
	"context"
	"fmt"
	"log/slog"

	"termcode/internal/domain/provider"
)

type ProviderRepo interface {
	GetByID(ctx context.Context, id string) (*provider.Provider, error)
	GetDefault(ctx context.Context) (*provider.Provider, error)
}

type Manager struct {
	repo   ProviderRepo
	logger *slog.Logger
}

func NewManager(repo ProviderRepo, logger *slog.Logger) *Manager {
	return &Manager{
		repo:   repo,
		logger: logger.With("svc", "llm"),
	}
}

func (m *Manager) ResolveProvider(ctx context.Context, providerID string) (*provider.Provider, error) {
	if providerID != "" {
		return m.repo.GetByID(ctx, providerID)
	}
	return m.repo.GetDefault(ctx)
}

type ResolvedProvider struct {
	Provider *provider.Provider
	APIKey   string
}

func (m *Manager) ResolveWithKey(ctx context.Context, providerID string, decryptFn func(string) (string, error)) (*ResolvedProvider, error) {
	p, err := m.ResolveProvider(ctx, providerID)
	if err != nil {
		return nil, fmt.Errorf("resolve provider: %w", err)
	}

	apiKey := p.APIKey
	if apiKey != "" {
		decrypted, err := decryptFn(apiKey)
		if err == nil {
			apiKey = decrypted
		}
	}

	return &ResolvedProvider{
		Provider: p,
		APIKey:   apiKey,
	}, nil
}
