package plugin

import (
	"context"
	"log/slog"

	"termcode/internal/domain/plugin"
)

type Repository interface {
	Create(ctx context.Context, p *plugin.Plugin) error
	GetByID(ctx context.Context, id string) (*plugin.Plugin, error)
	List(ctx context.Context) ([]*plugin.Plugin, error)
	Update(ctx context.Context, p *plugin.Plugin) error
	Delete(ctx context.Context, id string) error
}

type Service struct {
	repo   Repository
	logger *slog.Logger
}

func NewService(repo Repository, logger *slog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger.With("svc", "plugin"),
	}
}

func (s *Service) List(ctx context.Context) ([]*plugin.Plugin, error) {
	return s.repo.List(ctx)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
