package mcp

import (
	"context"
	"fmt"
	"log/slog"

	"termcode/internal/domain/mcp"
)

type Repository interface {
	Create(ctx context.Context, s *mcp.Server) error
	GetByID(ctx context.Context, id string) (*mcp.Server, error)
	List(ctx context.Context) ([]*mcp.Server, error)
	Update(ctx context.Context, s *mcp.Server) error
	Delete(ctx context.Context, id string) error
}

type Service struct {
	repo   Repository
	logger *slog.Logger
}

func NewService(repo Repository, logger *slog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger.With("svc", "mcp"),
	}
}

func (s *Service) Create(ctx context.Context, name, command string, args []string) (*mcp.Server, error) {
	server := mcp.NewStdio(name, command, args)
	if err := s.repo.Create(ctx, server); err != nil {
		return nil, fmt.Errorf("create MCP server: %w", err)
	}
	s.logger.Info("MCP server created", "name", name)
	return server, nil
}

func (s *Service) List(ctx context.Context) ([]*mcp.Server, error) {
	return s.repo.List(ctx)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
