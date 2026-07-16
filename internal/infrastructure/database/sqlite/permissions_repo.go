package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"termcode/internal/domain/permission"
)

type PermissionsRepo struct {
	db *sql.DB
}

func NewPermissionsRepo(db *sql.DB) *PermissionsRepo {
	return &PermissionsRepo{db: db}
}

func (r *PermissionsRepo) Get(ctx context.Context, toolName string) (*permission.Entry, error) {
	row := r.db.QueryRowContext(ctx, `SELECT tool_name, permission, updated_at FROM permissions WHERE tool_name = ?`, toolName)
	return scanPermissionEntry(row)
}

func (r *PermissionsRepo) Set(ctx context.Context, e *permission.Entry) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO permissions (tool_name, permission, updated_at) VALUES (?, ?, ?)
		ON CONFLICT(tool_name) DO UPDATE SET permission=excluded.permission, updated_at=excluded.updated_at`,
		e.ToolName, string(e.Permission), e.UpdatedAt.UTC().Format(time.RFC3339))
	return err
}

func (r *PermissionsRepo) List(ctx context.Context) ([]*permission.Entry, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT tool_name, permission, updated_at FROM permissions ORDER BY tool_name`)
	if err != nil {
		return nil, fmt.Errorf("list permissions: %w", err)
	}
	defer rows.Close()

	var entries []*permission.Entry
	for rows.Next() {
		e, err := scanPermissionEntry(rows)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, rows.Err()
}

func (r *PermissionsRepo) Delete(ctx context.Context, toolName string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM permissions WHERE tool_name = ?`, toolName)
	return err
}

func scanPermissionEntry(s interface{ Scan(dest ...any) error }) (*permission.Entry, error) {
	var e permission.Entry
	var level, updatedAt string

	err := s.Scan(&e.ToolName, &level, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, permission.ErrNotFound
		}
		return nil, fmt.Errorf("scan permission: %w", err)
	}

	e.Permission = permission.Level(level)
	e.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &e, nil
}
