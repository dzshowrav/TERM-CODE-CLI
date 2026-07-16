package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"termcode/internal/domain/model"
)

type ModelRepo struct {
	db *sql.DB
}

func NewModelRepo(db *sql.DB) *ModelRepo {
	return &ModelRepo{db: db}
}

func (r *ModelRepo) Create(ctx context.Context, m *model.Model) error {
	caps, err := json.Marshal(m.Capabilities)
	if err != nil {
		return fmt.Errorf("marshal capabilities: %w", err)
	}

	_, err = r.db.ExecContext(
		ctx, `
		INSERT INTO models (id, provider_id, model_id, display_name, description, category, capabilities, max_context, max_output, pricing_input, pricing_output, is_local, is_favorite, enabled, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(provider_id, model_id) DO UPDATE SET
			display_name=excluded.display_name, description=excluded.description, category=excluded.category,
			capabilities=excluded.capabilities, max_context=excluded.max_context, max_output=excluded.max_output,
			pricing_input=excluded.pricing_input, pricing_output=excluded.pricing_output, enabled=excluded.enabled,
			updated_at=excluded.updated_at`,
		m.ID, m.ProviderID, m.ModelID, m.DisplayName, m.Description, string(m.Category), string(caps),
		m.MaxContext, m.MaxOutput, m.PricingInput, m.PricingOut, boolToInt(m.IsLocal), boolToInt(m.IsFavorite),
		boolToInt(m.Enabled), m.CreatedAt.UTC().Format(time.RFC3339), m.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create model: %w", err)
	}
	return nil
}

func (r *ModelRepo) GetByID(ctx context.Context, id string) (*model.Model, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, provider_id, model_id, display_name, description, category, capabilities, max_context, max_output, pricing_input, pricing_output, is_local, is_favorite, enabled, created_at, updated_at
		FROM models WHERE id = ?`, id)

	return scanModel(row)
}

func (r *ModelRepo) List(ctx context.Context) ([]*model.Model, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, provider_id, model_id, display_name, description, category, capabilities, max_context, max_output, pricing_input, pricing_output, is_local, is_favorite, enabled, created_at, updated_at
		FROM models ORDER BY category, display_name`)
	if err != nil {
		return nil, fmt.Errorf("list models: %w", err)
	}
	defer rows.Close()

	return scanModels(rows)
}

func (r *ModelRepo) ListByProvider(ctx context.Context, providerID string) ([]*model.Model, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, provider_id, model_id, display_name, description, category, capabilities, max_context, max_output, pricing_input, pricing_output, is_local, is_favorite, enabled, created_at, updated_at
		FROM models WHERE provider_id = ? ORDER BY display_name`, providerID)
	if err != nil {
		return nil, fmt.Errorf("list models by provider: %w", err)
	}
	defer rows.Close()

	return scanModels(rows)
}

func (r *ModelRepo) DeleteByProvider(ctx context.Context, providerID string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM models WHERE provider_id = ?`, providerID)
	return err
}

func (r *ModelRepo) Update(ctx context.Context, m *model.Model) error {
	caps, err := json.Marshal(m.Capabilities)
	if err != nil {
		return fmt.Errorf("marshal capabilities: %w", err)
	}

	m.UpdatedAt = time.Now()
	res, err := r.db.ExecContext(
		ctx, `
		UPDATE models SET display_name=?, description=?, category=?, capabilities=?, max_context=?, max_output=?, pricing_input=?, pricing_output=?, is_favorite=?, enabled=?, updated_at=?
		WHERE id=?`,
		m.DisplayName, m.Description, string(m.Category), string(caps), m.MaxContext, m.MaxOutput,
		m.PricingInput, m.PricingOut, boolToInt(m.IsFavorite), boolToInt(m.Enabled),
		m.UpdatedAt.UTC().Format(time.RFC3339), m.ID,
	)
	if err != nil {
		return fmt.Errorf("update model: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return model.ErrNotFound
	}
	return nil
}

func (r *ModelRepo) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM models WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete model: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return model.ErrNotFound
	}
	return nil
}

func (r *ModelRepo) SetFavorite(ctx context.Context, id string, favorite bool) error {
	res, err := r.db.ExecContext(ctx, `UPDATE models SET is_favorite = ? WHERE id = ?`, boolToInt(favorite), id)
	if err != nil {
		return fmt.Errorf("set favorite: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return model.ErrNotFound
	}
	return nil
}

func scanModel(s interface {
	Scan(dest ...any) error
},
) (*model.Model, error) {
	var m model.Model
	var category, capsJSON, createdAt, updatedAt string
	var isLocal, isFav, enabled int

	err := s.Scan(&m.ID, &m.ProviderID, &m.ModelID, &m.DisplayName, &m.Description, &category, &capsJSON,
		&m.MaxContext, &m.MaxOutput, &m.PricingInput, &m.PricingOut, &isLocal, &isFav, &enabled,
		&createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrNotFound
		}
		return nil, fmt.Errorf("scan model: %w", err)
	}

	m.Category = model.Category(category)
	m.IsLocal = isLocal == 1
	m.IsFavorite = isFav == 1
	m.Enabled = enabled == 1

	if capsJSON != "" && capsJSON != "{}" {
		json.Unmarshal([]byte(capsJSON), &m.Capabilities)
	}

	m.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	m.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &m, nil
}

func scanModels(rows *sql.Rows) ([]*model.Model, error) {
	var models []*model.Model
	for rows.Next() {
		m, err := scanModel(rows)
		if err != nil {
			return nil, err
		}
		models = append(models, m)
	}
	return models, rows.Err()
}
