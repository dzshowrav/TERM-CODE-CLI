package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"termcode/internal/domain/plugin"
)

type PluginRepo struct {
	db *sql.DB
}

func NewPluginRepo(db *sql.DB) *PluginRepo {
	return &PluginRepo{db: db}
}

func (r *PluginRepo) Create(ctx context.Context, p *plugin.Plugin) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO plugins (id, name, version, author, description, status, enabled, installed_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		p.ID, p.Name, p.Version, p.Author, p.Description, string(p.Status), boolToInt(p.Enabled),
		p.CreatedAt.UTC().Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("create plugin: %w", err)
	}
	return nil
}

func (r *PluginRepo) GetByID(ctx context.Context, id string) (*plugin.Plugin, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, version, author, description, status, enabled, installed_at
		FROM plugins WHERE id = ?`, id)
	return scanPlugin(row)
}

func (r *PluginRepo) List(ctx context.Context) ([]*plugin.Plugin, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, version, author, description, status, enabled, installed_at
		FROM plugins ORDER BY name`)
	if err != nil {
		return nil, fmt.Errorf("list plugins: %w", err)
	}
	defer rows.Close()

	var plugins []*plugin.Plugin
	for rows.Next() {
		p, err := scanPlugin(rows)
		if err != nil {
			return nil, err
		}
		plugins = append(plugins, p)
	}
	return plugins, rows.Err()
}

func (r *PluginRepo) Update(ctx context.Context, p *plugin.Plugin) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE plugins SET name=?, version=?, author=?, description=?, status=?, enabled=? WHERE id=?`,
		p.Name, p.Version, p.Author, p.Description, string(p.Status), boolToInt(p.Enabled), p.ID)
	return err
}

func (r *PluginRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM plugins WHERE id = ?`, id)
	return err
}

func scanPlugin(s interface{ Scan(dest ...any) error }) (*plugin.Plugin, error) {
	var p plugin.Plugin
	var status, installedAt string
	var enabled int

	err := s.Scan(&p.ID, &p.Name, &p.Version, &p.Author, &p.Description, &status, &enabled, &installedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, plugin.ErrNotFound
		}
		return nil, fmt.Errorf("scan plugin: %w", err)
	}

	p.Status = plugin.Status(status)
	p.Enabled = enabled == 1
	p.CreatedAt, _ = time.Parse(time.RFC3339, installedAt)

	return &p, nil
}
