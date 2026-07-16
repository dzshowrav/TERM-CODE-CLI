package agent

import (
	"context"
	"fmt"
	"log/slog"

	"termcode/internal/domain/agent"
)

type Repository interface {
	Create(ctx context.Context, a *agent.Agent) error
	GetByID(ctx context.Context, id string) (*agent.Agent, error)
	List(ctx context.Context) ([]*agent.Agent, error)
	Update(ctx context.Context, a *agent.Agent) error
	Delete(ctx context.Context, id string) error
	SetDefault(ctx context.Context, id string) error
	GetDefault(ctx context.Context) (*agent.Agent, error)
}

type Service struct {
	repo   Repository
	logger *slog.Logger
}

func NewService(repo Repository, logger *slog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger.With("svc", "agent"),
	}
}

func (s *Service) List(ctx context.Context) ([]*agent.Agent, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetDefault(ctx context.Context) (*agent.Agent, error) {
	return s.repo.GetDefault(ctx)
}

func (s *Service) Create(ctx context.Context, name, description, systemPrompt string) (*agent.Agent, error) {
	a := agent.New(name, description, systemPrompt)
	if err := s.repo.Create(ctx, a); err != nil {
		return nil, fmt.Errorf("create agent: %w", err)
	}
	return a, nil
}
