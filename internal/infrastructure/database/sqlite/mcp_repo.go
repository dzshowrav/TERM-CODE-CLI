package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"termcode/internal/domain/mcp"
)

type MCPRepo struct {
	db *sql.DB
}

func NewMCPRepo(db *sql.DB) *MCPRepo {
	return &MCPRepo{db: db}
}

func (r *MCPRepo) Create(ctx context.Context, s *mcp.Server) error {
	args, _ := json.Marshal(s.Args)
	env, _ := json.Marshal(s.Env)
	_, err := r.db.ExecContext(
		ctx, `
		INSERT INTO mcp_servers (id, name, transport, command, args, url, env, status, enabled, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		s.ID, s.Name, string(s.Transport), s.Command, string(args), s.URL, string(env),
		string(s.Status), boolToInt(s.Enabled),
		s.CreatedAt.UTC().Format(time.RFC3339), s.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create MCP server: %w", err)
	}
	return nil
}

func (r *MCPRepo) GetByID(ctx context.Context, id string) (*mcp.Server, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, transport, command, args, url, env, status, enabled, created_at, updated_at
		FROM mcp_servers WHERE id = ?`, id)
	return scanMCPServer(row)
}

func (r *MCPRepo) List(ctx context.Context) ([]*mcp.Server, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, transport, command, args, url, env, status, enabled, created_at, updated_at
		FROM mcp_servers ORDER BY name`)
	if err != nil {
		return nil, fmt.Errorf("list MCP servers: %w", err)
	}
	defer rows.Close()

	var servers []*mcp.Server
	for rows.Next() {
		s, err := scanMCPServer(rows)
		if err != nil {
			return nil, err
		}
		servers = append(servers, s)
	}
	return servers, rows.Err()
}

func (r *MCPRepo) Update(ctx context.Context, s *mcp.Server) error {
	args, _ := json.Marshal(s.Args)
	env, _ := json.Marshal(s.Env)
	s.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		UPDATE mcp_servers SET name=?, transport=?, command=?, args=?, url=?, env=?, status=?, enabled=?, updated_at=?
		WHERE id=?`,
		s.Name, string(s.Transport), s.Command, string(args), s.URL, string(env),
		string(s.Status), boolToInt(s.Enabled), s.UpdatedAt.UTC().Format(time.RFC3339), s.ID)
	return err
}

func (r *MCPRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM mcp_servers WHERE id = ?`, id)
	return err
}

func scanMCPServer(s interface{ Scan(dest ...any) error }) (*mcp.Server, error) {
	var sv mcp.Server
	var transport, status, createdAt, updatedAt, argsJSON, envJSON string
	var enabled int

	err := s.Scan(&sv.ID, &sv.Name, &transport, &sv.Command, &argsJSON, &sv.URL, &envJSON, &status, &enabled, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, mcp.ErrNotFound
		}
		return nil, fmt.Errorf("scan MCP server: %w", err)
	}

	sv.Transport = mcp.Transport(transport)
	sv.Status = mcp.Status(status)
	sv.Enabled = enabled == 1

	if argsJSON != "" {
		json.Unmarshal([]byte(argsJSON), &sv.Args)
	}
	if envJSON != "" {
		json.Unmarshal([]byte(envJSON), &sv.Env)
	}

	sv.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	sv.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &sv, nil
}
