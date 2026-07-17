package router

import (
	"context"
	"fmt"

	domainllm "termcode/internal/domain/llm"
	"termcode/internal/domain/model"
	"termcode/internal/domain/provider"
)

type ProviderService interface {
	GetByID(ctx context.Context, id string) (*provider.Provider, error)
	GetDefault(ctx context.Context) (*provider.Provider, error)
	DecryptAPIKey(encrypted string) (string, error)
}

type ModelLister interface {
	List(ctx context.Context) ([]*model.Model, error)
}

type RuntimeRouter struct {
	providerSvc ProviderService
	modelLister ModelLister
}

func NewRuntimeRouter(providerSvc ProviderService, modelLister ModelLister) *RuntimeRouter {
	return &RuntimeRouter{
		providerSvc: providerSvc,
		modelLister: modelLister,
	}
}

func (r *RuntimeRouter) Resolve(ctx context.Context, modelID string) (*domainllm.Endpoint, error) {
	var p *provider.Provider

	models, err := r.modelLister.List(ctx)
	if err == nil {
		for _, m := range models {
			if m.ModelID == modelID && m.ProviderID != "" {
				p, err = r.providerSvc.GetByID(ctx, m.ProviderID)
				if err == nil {
					break
				}
			}
		}
	}

	if p == nil {
		p, err = r.providerSvc.GetDefault(ctx)
		if err != nil {
			return nil, fmt.Errorf("resolve provider: %w", err)
		}
	}

	apiKey := p.APIKey
	if apiKey != "" {
		decrypted, err := r.providerSvc.DecryptAPIKey(apiKey)
		if err == nil {
			apiKey = decrypted
		}
	}

	return &domainllm.Endpoint{
		BaseURL:      p.BaseURL,
		APIKey:       apiKey,
		ModelID:      modelID,
		ProviderName: p.Name,
	}, nil
}

func (r *RuntimeRouter) ResolveProvider(ctx context.Context, modelID string) (*ProviderInfo, error) {
	endpoint, err := r.Resolve(ctx, modelID)
	if err != nil {
		return nil, err
	}
	return &ProviderInfo{
		ID:     "",
		Name:   endpoint.ProviderName,
		URL:    endpoint.BaseURL,
		APIKey: endpoint.APIKey,
	}, nil
}
