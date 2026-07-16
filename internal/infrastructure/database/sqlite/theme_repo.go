package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"termcode/internal/domain/theme"
)

type ThemeRepo struct {
	db *sql.DB
}

func NewThemeRepo(db *sql.DB) *ThemeRepo {
	return &ThemeRepo{db: db}
}

func (r *ThemeRepo) Create(ctx context.Context, t *theme.Theme) error {
	_, err := r.db.ExecContext(
		ctx, `
		INSERT INTO themes (id, name, author, version, is_dark, palette, is_active, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		t.ID, t.Name, t.Author, t.Version, boolToInt(t.IsDark), t.Palette, boolToInt(t.IsActive),
		t.CreatedAt.UTC().Format(time.RFC3339), t.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create theme: %w", err)
	}
	return nil
}

func (r *ThemeRepo) GetByID(ctx context.Context, id string) (*theme.Theme, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, author, version, is_dark, palette, is_active, created_at, updated_at
		FROM themes WHERE id = ?`, id)
	return scanTheme(row)
}

func (r *ThemeRepo) List(ctx context.Context) ([]*theme.Theme, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, author, version, is_dark, palette, is_active, created_at, updated_at
		FROM themes ORDER BY name`)
	if err != nil {
		return nil, fmt.Errorf("list themes: %w", err)
	}
	defer rows.Close()

	var themes []*theme.Theme
	for rows.Next() {
		t, err := scanTheme(rows)
		if err != nil {
			return nil, err
		}
		themes = append(themes, t)
	}
	return themes, rows.Err()
}

func (r *ThemeRepo) Update(ctx context.Context, t *theme.Theme) error {
	t.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		UPDATE themes SET name=?, author=?, version=?, is_dark=?, palette=?, is_active=?, updated_at=? WHERE id=?`,
		t.Name, t.Author, t.Version, boolToInt(t.IsDark), t.Palette, boolToInt(t.IsActive),
		t.UpdatedAt.UTC().Format(time.RFC3339), t.ID)
	return err
}

func (r *ThemeRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM themes WHERE id = ?`, id)
	return err
}

func scanTheme(s interface{ Scan(dest ...any) error }) (*theme.Theme, error) {
	var t theme.Theme
	var createdAt, updatedAt string
	var isDark, isActive int

	err := s.Scan(&t.ID, &t.Name, &t.Author, &t.Version, &isDark, &t.Palette, &isActive, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, theme.ErrNotFound
		}
		return nil, fmt.Errorf("scan theme: %w", err)
	}

	t.IsDark = isDark == 1
	t.IsActive = isActive == 1
	t.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	t.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &t, nil
}
