package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"termcode/internal/domain/skill"
)

type SkillRepo struct {
	db *sql.DB
}

func NewSkillRepo(db *sql.DB) *SkillRepo {
	return &SkillRepo{db: db}
}

func (r *SkillRepo) Create(ctx context.Context, s *skill.Skill) error {
	_, err := r.db.ExecContext(
		ctx, `
		INSERT INTO skills (id, name, version, category, description, path, enabled, is_builtin, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		s.ID, s.Name, s.Version, string(s.Category), s.Description, s.Path, boolToInt(s.Enabled), boolToInt(s.IsBuiltin),
		s.CreatedAt.UTC().Format(time.RFC3339), s.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create skill: %w", err)
	}
	return nil
}

func (r *SkillRepo) GetByID(ctx context.Context, id string) (*skill.Skill, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, version, category, description, path, enabled, is_builtin, created_at, updated_at
		FROM skills WHERE id = ?`, id)
	return scanSkill(row)
}

func (r *SkillRepo) List(ctx context.Context) ([]*skill.Skill, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, version, category, description, path, enabled, is_builtin, created_at, updated_at
		FROM skills ORDER BY category, name`)
	if err != nil {
		return nil, fmt.Errorf("list skills: %w", err)
	}
	defer rows.Close()

	var skills []*skill.Skill
	for rows.Next() {
		s, err := scanSkill(rows)
		if err != nil {
			return nil, err
		}
		skills = append(skills, s)
	}
	return skills, rows.Err()
}

func (r *SkillRepo) Update(ctx context.Context, s *skill.Skill) error {
	s.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		UPDATE skills SET name=?, version=?, category=?, description=?, path=?, enabled=?, is_builtin=?, updated_at=?
		WHERE id=?`,
		s.Name, s.Version, string(s.Category), s.Description, s.Path, boolToInt(s.Enabled), boolToInt(s.IsBuiltin),
		s.UpdatedAt.UTC().Format(time.RFC3339), s.ID)
	return err
}

func (r *SkillRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM skills WHERE id = ?`, id)
	return err
}

func scanSkill(s interface{ Scan(dest ...any) error }) (*skill.Skill, error) {
	var sk skill.Skill
	var cat, createdAt, updatedAt string
	var enabled, isBuiltin int

	err := s.Scan(&sk.ID, &sk.Name, &sk.Version, &cat, &sk.Description, &sk.Path, &enabled, &isBuiltin, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, skill.ErrNotFound
		}
		return nil, fmt.Errorf("scan skill: %w", err)
	}

	sk.Category = skill.Category(cat)
	sk.Enabled = enabled == 1
	sk.IsBuiltin = isBuiltin == 1
	sk.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	sk.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &sk, nil
}
