package router

import (
	"context"

	"termcode/internal/domain/model"
)

type RouteResult struct {
	Model    *model.Model
	Provider ProviderInfo
	Stream   bool
}

type ProviderInfo struct {
	ID     string
	Name   string
	URL    string
	APIKey string
}

type ProviderResolver interface {
	ResolveProvider(ctx context.Context, modelID string) (*ProviderInfo, error)
}

type Service struct {
	resolver ProviderResolver
}

func NewService(resolver ProviderResolver) *Service {
	return &Service{resolver: resolver}
}

func (s *Service) Route(ctx context.Context, m *model.Model) (*RouteResult, error) {
	providerInfo, err := s.resolver.ResolveProvider(ctx, m.ID)
	if err != nil {
		return nil, err
	}

	return &RouteResult{
		Model:    m,
		Provider: *providerInfo,
		Stream:   m.Capabilities.Streaming,
	}, nil
}
