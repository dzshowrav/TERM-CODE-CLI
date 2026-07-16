package model

import (
	"context"
	"fmt"

	"termcode/internal/domain/model"
)

type Service struct {
	repo     FullRepository
	discover *DiscoveryService
}

type FullRepository interface {
	Create(ctx context.Context, m *model.Model) error
	GetByID(ctx context.Context, id string) (*model.Model, error)
	List(ctx context.Context) ([]*model.Model, error)
	ListByProvider(ctx context.Context, providerID string) ([]*model.Model, error)
	Update(ctx context.Context, m *model.Model) error
	Delete(ctx context.Context, id string) error
	DeleteByProvider(ctx context.Context, providerID string) error
	SetFavorite(ctx context.Context, id string, favorite bool) error
}

func NewService(repo FullRepository) *Service {
	return &Service{
		repo:     repo,
		discover: NewDiscoveryService(repo),
	}
}

func (s *Service) GetByID(ctx context.Context, id string) (*model.Model, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]*model.Model, error) {
	return s.repo.List(ctx)
}

func (s *Service) ListByProvider(ctx context.Context, providerID string) ([]*model.Model, error) {
	return s.repo.ListByProvider(ctx, providerID)
}

func (s *Service) SyncFromProvider(ctx context.Context, providerID, baseURL, apiKey string) (int, error) {
	n, err := s.discover.discoverAndSync(ctx, providerID, baseURL, apiKey)
	if err != nil {
		return 0, fmt.Errorf("sync: %w", err)
	}
	return n, nil
}

func (s *Service) SetFavorite(ctx context.Context, id string, favorite bool) error {
	return s.repo.SetFavorite(ctx, id, favorite)
}
