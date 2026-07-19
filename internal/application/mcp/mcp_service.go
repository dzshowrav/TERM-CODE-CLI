package mcp

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	mcpclient "termcode/internal/infrastructure/mcp"

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
	repo    Repository
	logger  *slog.Logger
	manager *mcpclient.MCPManager
	mu      sync.RWMutex
}

func NewService(repo Repository, logger *slog.Logger) *Service {
	return &Service{
		repo:    repo,
		logger:  logger.With("svc", "mcp"),
		manager: mcpclient.NewMCPManager(),
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
	if err := s.manager.Disconnect(id); err != nil {
		s.logger.Warn("disconnect MCP server", "id", id, "err", err)
	}
	return s.repo.Delete(ctx, id)
}

func (s *Service) Connect(ctx context.Context, id string) error {
	server, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("get server: %w", err)
	}

	switch server.Transport {
	case mcp.TransportStdio:
		_, err = s.manager.ConnectStdio(ctx, id, server.Command, server.Args, server.Env)
	case mcp.TransportSSE:
		_, err = s.manager.ConnectSSE(ctx, id, server.URL)
	default:
		return fmt.Errorf("unsupported transport: %s", server.Transport)
	}

	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	server.Status = mcp.StatusConnected
	if err := s.repo.Update(ctx, server); err != nil {
		s.logger.Warn("update server status", "id", id, "err", err)
	}

	s.logger.Info("MCP server connected", "id", id, "name", server.Name)
	return nil
}

func (s *Service) Disconnect(ctx context.Context, id string) error {
	server, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("get server: %w", err)
	}

	if err := s.manager.Disconnect(id); err != nil {
		s.logger.Warn("disconnect", "id", id, "err", err)
	}

	server.Status = mcp.StatusDisconnected
	if err := s.repo.Update(ctx, server); err != nil {
		return fmt.Errorf("update status: %w", err)
	}

	return nil
}

func (s *Service) ListTools(ctx context.Context, serverID string) ([]mcpclient.ToolInfo, error) {
	client, ok := s.manager.Get(serverID)
	if !ok {
		return nil, fmt.Errorf("server %s not connected", serverID)
	}
	return client.ListTools(ctx)
}

func (s *Service) CallTool(ctx context.Context, serverID, toolName string, args map[string]any) (*mcpclient.ToolResult, error) {
	client, ok := s.manager.Get(serverID)
	if !ok {
		return nil, fmt.Errorf("server %s not connected", serverID)
	}
	return client.CallTool(ctx, toolName, args)
}

func (s *Service) GetClient(serverID string) (*mcpclient.MCPClient, bool) {
	return s.manager.Get(serverID)
}

func (s *Service) DisconnectAll() error {
	return s.manager.DisconnectAll()
}
