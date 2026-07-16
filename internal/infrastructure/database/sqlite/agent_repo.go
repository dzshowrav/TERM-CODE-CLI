package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"termcode/internal/domain/agent"
)

type AgentRepo struct {
	db *sql.DB
}

func NewAgentRepo(db *sql.DB) *AgentRepo {
	return &AgentRepo{db: db}
}

func (r *AgentRepo) Create(ctx context.Context, a *agent.Agent) error {
	tools, _ := json.Marshal(a.Tools)
	_, err := r.db.ExecContext(
		ctx, `
		INSERT INTO agents (id, name, description, system_prompt, model_id, tools, is_default, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		a.ID, a.Name, a.Description, a.SystemPrompt, a.ModelID, string(tools), boolToInt(a.IsDefault),
		a.CreatedAt.UTC().Format(time.RFC3339), a.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create agent: %w", err)
	}
	return nil
}

func (r *AgentRepo) GetByID(ctx context.Context, id string) (*agent.Agent, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, description, system_prompt, model_id, tools, is_default, created_at, updated_at
		FROM agents WHERE id = ?`, id)
	return scanAgent(row)
}

func (r *AgentRepo) List(ctx context.Context) ([]*agent.Agent, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, description, system_prompt, model_id, tools, is_default, created_at, updated_at
		FROM agents ORDER BY is_default DESC, name`)
	if err != nil {
		return nil, fmt.Errorf("list agents: %w", err)
	}
	defer rows.Close()

	var agents []*agent.Agent
	for rows.Next() {
		a, err := scanAgent(rows)
		if err != nil {
			return nil, err
		}
		agents = append(agents, a)
	}
	return agents, rows.Err()
}

func (r *AgentRepo) Update(ctx context.Context, a *agent.Agent) error {
	tools, _ := json.Marshal(a.Tools)
	a.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		UPDATE agents SET name=?, description=?, system_prompt=?, model_id=?, tools=?, is_default=?, updated_at=?
		WHERE id=?`,
		a.Name, a.Description, a.SystemPrompt, a.ModelID, string(tools), boolToInt(a.IsDefault),
		a.UpdatedAt.UTC().Format(time.RFC3339), a.ID)
	return err
}

func (r *AgentRepo) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM agents WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete agent: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return agent.ErrNotFound
	}
	return nil
}

func (r *AgentRepo) SetDefault(ctx context.Context, id string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `UPDATE agents SET is_default = 0`); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE agents SET is_default = 1 WHERE id = ?`, id); err != nil {
		return err
	}
	return tx.Commit()
}

func (r *AgentRepo) GetDefault(ctx context.Context) (*agent.Agent, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, description, system_prompt, model_id, tools, is_default, created_at, updated_at
		FROM agents WHERE is_default = 1 LIMIT 1`)
	return scanAgent(row)
}

func scanAgent(s interface{ Scan(dest ...any) error }) (*agent.Agent, error) {
	var a agent.Agent
	var createdAt, updatedAt, toolsJSON string
	var isDefault int

	err := s.Scan(&a.ID, &a.Name, &a.Description, &a.SystemPrompt, &a.ModelID, &toolsJSON, &isDefault, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, agent.ErrNotFound
		}
		return nil, fmt.Errorf("scan agent: %w", err)
	}

	a.IsDefault = isDefault == 1
	if toolsJSON != "" && toolsJSON != "[]" {
		json.Unmarshal([]byte(toolsJSON), &a.Tools)
	}
	a.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	a.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &a, nil
}

func EnsureDefaultAgent(db *sql.DB) error {
	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM agents`).Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		_, err := db.Exec(`
			INSERT INTO agents (id, name, description, system_prompt, model_id, tools, is_default, created_at, updated_at)
			VALUES ('default', 'General', 'General purpose coding agent', 'You are a helpful coding assistant.', '', '[]', 1, datetime('now'), datetime('now'))`)
		return err
	}
	return nil
}
