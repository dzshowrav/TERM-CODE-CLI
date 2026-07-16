package git

import (
	"context"
	"fmt"
	"log/slog"

	gogit "termcode/internal/infrastructure/git"
)

type Service struct {
	git    *gogit.Service
	logger *slog.Logger
}

func NewService(logger *slog.Logger) *Service {
	return &Service{
		git:    gogit.NewService(),
		logger: logger.With("svc", "git"),
	}
}

func (s *Service) Status(ctx context.Context, repoPath string) (*gogit.StatusResult, error) {
	repo, err := s.git.Open(repoPath)
	if err != nil {
		return nil, fmt.Errorf("open repo: %w", err)
	}
	return s.git.Status(repo)
}

func (s *Service) Log(ctx context.Context, repoPath string, count int) ([]gogit.LogEntry, error) {
	repo, err := s.git.Open(repoPath)
	if err != nil {
		return nil, fmt.Errorf("open repo: %w", err)
	}
	return s.git.Log(repo, count)
}

func (s *Service) Diff(ctx context.Context, repoPath string) (*gogit.DiffResult, error) {
	repo, err := s.git.Open(repoPath)
	if err != nil {
		return nil, fmt.Errorf("open repo: %w", err)
	}
	return s.git.Diff(repo)
}

func (s *Service) AddFiles(ctx context.Context, repoPath string, files []string) error {
	repo, err := s.git.Open(repoPath)
	if err != nil {
		return fmt.Errorf("open repo: %w", err)
	}
	if len(files) == 0 {
		return s.git.AddAll(repo)
	}
	return s.git.Add(repo, files)
}

func (s *Service) Commit(ctx context.Context, repoPath, message string) (string, error) {
	repo, err := s.git.Open(repoPath)
	if err != nil {
		return "", fmt.Errorf("open repo: %w", err)
	}
	return s.git.Commit(repo, message)
}
