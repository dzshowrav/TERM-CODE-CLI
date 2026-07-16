package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"termcode/internal/domain/provider"
)

type ProviderRepo struct {
	db *sql.DB
}

func NewProviderRepo(db *sql.DB) *ProviderRepo {
	return &ProviderRepo{db: db}
}

func (r *ProviderRepo) Create(ctx context.Context, p *provider.Provider) error {
	_, err := r.db.ExecContext(
		ctx, `
		INSERT INTO providers (id, name, base_url, api_key, description, status, latency_ms, priority, is_default, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.ID, p.Name, p.BaseURL, p.APIKey, p.Description, string(p.Status), p.Latency, p.Priority, boolToInt(p.IsDefault), p.CreatedAt.UTC().Format(time.RFC3339), p.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create provider: %w", err)
	}
	return nil
}

func (r *ProviderRepo) GetByID(ctx context.Context, id string) (*provider.Provider, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, base_url, api_key, description, status, latency_ms, priority, is_default, created_at, updated_at
		FROM providers WHERE id = ?`, id)

	return scanProvider(row)
}

func (r *ProviderRepo) GetByName(ctx context.Context, name string) (*provider.Provider, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, base_url, api_key, description, status, latency_ms, priority, is_default, created_at, updated_at
		FROM providers WHERE name = ?`, name)

	return scanProvider(row)
}

func (r *ProviderRepo) List(ctx context.Context) ([]*provider.Provider, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, base_url, api_key, description, status, latency_ms, priority, is_default, created_at, updated_at
		FROM providers ORDER BY priority ASC, name ASC`)
	if err != nil {
		return nil, fmt.Errorf("list providers: %w", err)
	}
	defer rows.Close()

	var providers []*provider.Provider
	for rows.Next() {
		p, err := scanProvider(rows)
		if err != nil {
			return nil, err
		}
		providers = append(providers, p)
	}
	return providers, rows.Err()
}

func (r *ProviderRepo) Update(ctx context.Context, p *provider.Provider) error {
	p.UpdatedAt = time.Now()
	res, err := r.db.ExecContext(
		ctx, `
		UPDATE providers SET name=?, base_url=?, api_key=?, description=?, status=?, latency_ms=?, priority=?, is_default=?, updated_at=?
		WHERE id=?`,
		p.Name, p.BaseURL, p.APIKey, p.Description, string(p.Status), p.Latency, p.Priority, boolToInt(p.IsDefault), p.UpdatedAt.UTC().Format(time.RFC3339), p.ID,
	)
	if err != nil {
		return fmt.Errorf("update provider: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return provider.ErrNotFound
	}
	return nil
}

func (r *ProviderRepo) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM providers WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete provider: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return provider.ErrNotFound
	}
	return nil
}

func (r *ProviderRepo) SetDefault(ctx context.Context, id string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `UPDATE providers SET is_default = 0`); err != nil {
		return fmt.Errorf("clear defaults: %w", err)
	}
	if _, err := tx.ExecContext(ctx, `UPDATE providers SET is_default = 1 WHERE id = ?`, id); err != nil {
		return fmt.Errorf("set default: %w", err)
	}

	return tx.Commit()
}

func (r *ProviderRepo) GetDefault(ctx context.Context) (*provider.Provider, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, base_url, api_key, description, status, latency_ms, priority, is_default, created_at, updated_at
		FROM providers WHERE is_default = 1 LIMIT 1`)

	p, err := scanProvider(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, provider.ErrNoDefault
		}
		return nil, err
	}
	return p, nil
}

func scanProvider(s interface {
	Scan(dest ...any) error
},
) (*provider.Provider, error) {
	var p provider.Provider
	var status, createdAt, updatedAt string
	var isDefault int

	err := s.Scan(&p.ID, &p.Name, &p.BaseURL, &p.APIKey, &p.Description, &status, &p.Latency, &p.Priority, &isDefault, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, provider.ErrNotFound
		}
		return nil, fmt.Errorf("scan provider: %w", err)
	}

	p.Status = provider.Status(status)
	p.IsDefault = isDefault == 1
	p.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	p.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &p, nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
