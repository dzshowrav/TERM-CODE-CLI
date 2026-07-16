package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"termcode/internal/domain/workspace"
)

type WorkspaceRepo struct {
	db *sql.DB
}

func NewWorkspaceRepo(db *sql.DB) *WorkspaceRepo {
	return &WorkspaceRepo{db: db}
}

func (r *WorkspaceRepo) Create(ctx context.Context, w *workspace.Workspace) error {
	_, err := r.db.ExecContext(
		ctx, `
		INSERT INTO workspaces (id, name, path, is_default, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)`,
		w.ID, w.Name, w.Path, boolToInt(w.IsDefault),
		w.CreatedAt.UTC().Format(time.RFC3339), w.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create workspace: %w", err)
	}
	return nil
}

func (r *WorkspaceRepo) GetByID(ctx context.Context, id string) (*workspace.Workspace, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, path, is_default, created_at, updated_at
		FROM workspaces WHERE id = ?`, id)
	return scanWorkspace(row)
}

func (r *WorkspaceRepo) GetByPath(ctx context.Context, path string) (*workspace.Workspace, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, path, is_default, created_at, updated_at
		FROM workspaces WHERE path = ?`, path)
	return scanWorkspace(row)
}

func (r *WorkspaceRepo) List(ctx context.Context) ([]*workspace.Workspace, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, path, is_default, created_at, updated_at
		FROM workspaces ORDER BY name`)
	if err != nil {
		return nil, fmt.Errorf("list workspaces: %w", err)
	}
	defer rows.Close()

	var ws []*workspace.Workspace
	for rows.Next() {
		w, err := scanWorkspace(rows)
		if err != nil {
			return nil, err
		}
		ws = append(ws, w)
	}
	return ws, rows.Err()
}

func (r *WorkspaceRepo) Update(ctx context.Context, w *workspace.Workspace) error {
	w.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		UPDATE workspaces SET name=?, path=?, is_default=?, updated_at=? WHERE id=?`,
		w.Name, w.Path, boolToInt(w.IsDefault), w.UpdatedAt.UTC().Format(time.RFC3339), w.ID)
	return err
}

func (r *WorkspaceRepo) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM workspaces WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete workspace: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return workspace.ErrNotFound
	}
	return nil
}

func (r *WorkspaceRepo) SetDefault(ctx context.Context, id string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `UPDATE workspaces SET is_default = 0`); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE workspaces SET is_default = 1 WHERE id = ?`, id); err != nil {
		return err
	}
	return tx.Commit()
}

func (r *WorkspaceRepo) GetDefault(ctx context.Context) (*workspace.Workspace, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, path, is_default, created_at, updated_at
		FROM workspaces WHERE is_default = 1 LIMIT 1`)
	w, err := scanWorkspace(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, workspace.ErrNotFound
		}
		return nil, err
	}
	return w, nil
}

func scanWorkspace(s interface{ Scan(dest ...any) error }) (*workspace.Workspace, error) {
	var w workspace.Workspace
	var createdAt, updatedAt string
	var isDefault int

	err := s.Scan(&w.ID, &w.Name, &w.Path, &isDefault, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, workspace.ErrNotFound
		}
		return nil, fmt.Errorf("scan workspace: %w", err)
	}

	w.IsDefault = isDefault == 1
	w.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	w.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &w, nil
}
