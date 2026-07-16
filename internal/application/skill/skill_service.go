package skill

import (
	"context"
	"fmt"
	"log/slog"

	"termcode/internal/domain/skill"
)

type Repository interface {
	Create(ctx context.Context, s *skill.Skill) error
	GetByID(ctx context.Context, id string) (*skill.Skill, error)
	List(ctx context.Context) ([]*skill.Skill, error)
	Update(ctx context.Context, s *skill.Skill) error
	Delete(ctx context.Context, id string) error
}

type Service struct {
	repo   Repository
	logger *slog.Logger
}

func NewService(repo Repository, logger *slog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger.With("svc", "skill"),
	}
}

func (s *Service) List(ctx context.Context) ([]*skill.Skill, error) {
	return s.repo.List(ctx)
}

func (s *Service) Create(ctx context.Context, name, description string, category skill.Category) (*skill.Skill, error) {
	sk := skill.New(name, description, category)
	if err := s.repo.Create(ctx, sk); err != nil {
		return nil, fmt.Errorf("create skill: %w", err)
	}
	return sk, nil
}
