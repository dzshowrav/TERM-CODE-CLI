package workspace

import (
	"context"
	"fmt"
	"log/slog"

	"termcode/internal/domain/workspace"
)

type Repository interface {
	Create(ctx context.Context, w *workspace.Workspace) error
	GetByID(ctx context.Context, id string) (*workspace.Workspace, error)
	GetByPath(ctx context.Context, path string) (*workspace.Workspace, error)
	List(ctx context.Context) ([]*workspace.Workspace, error)
	Update(ctx context.Context, w *workspace.Workspace) error
	Delete(ctx context.Context, id string) error
	SetDefault(ctx context.Context, id string) error
	GetDefault(ctx context.Context) (*workspace.Workspace, error)
}

type Service struct {
	repo   Repository
	logger *slog.Logger
}

func NewService(repo Repository, logger *slog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger.With("svc", "workspace"),
	}
}

func (s *Service) Create(ctx context.Context, name, path string) (*workspace.Workspace, error) {
	w := workspace.New(name, path)
	if err := s.repo.Create(ctx, w); err != nil {
		return nil, fmt.Errorf("create workspace: %w", err)
	}
	s.logger.Info("workspace created", "name", name, "path", path)
	return w, nil
}

func (s *Service) List(ctx context.Context) ([]*workspace.Workspace, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetDefault(ctx context.Context) (*workspace.Workspace, error) {
	return s.repo.GetDefault(ctx)
}

func (s *Service) SetDefault(ctx context.Context, id string) error {
	return s.repo.SetDefault(ctx, id)
}
