package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"termcode/internal/domain/session"
)

type BranchRepo struct {
	db *sql.DB
}

func NewBranchRepo(db *sql.DB) *BranchRepo {
	return &BranchRepo{db: db}
}

type BranchRecord struct {
	ID        string
	SessionID string
	Name      string
	History   []session.Message
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *BranchRepo) Create(ctx context.Context, b *BranchRecord) error {
	historyJSON, err := json.Marshal(b.History)
	if err != nil {
		return fmt.Errorf("marshal history: %w", err)
	}
	_, err = r.db.ExecContext(ctx, `
		INSERT INTO branches (id, session_id, name, history, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)`,
		b.ID, b.SessionID, b.Name, string(historyJSON),
		b.CreatedAt.UTC().Format(time.RFC3339),
		b.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create branch: %w", err)
	}
	return nil
}

func (r *BranchRepo) Update(ctx context.Context, b *BranchRecord) error {
	historyJSON, err := json.Marshal(b.History)
	if err != nil {
		return fmt.Errorf("marshal history: %w", err)
	}
	_, err = r.db.ExecContext(ctx, `
		UPDATE branches SET name = ?, history = ?, updated_at = ? WHERE id = ?`,
		b.Name, string(historyJSON), time.Now().UTC().Format(time.RFC3339), b.ID,
	)
	if err != nil {
		return fmt.Errorf("update branch: %w", err)
	}
	return nil
}

func (r *BranchRepo) ListBySession(ctx context.Context, sessionID string) ([]*BranchRecord, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, session_id, name, history, created_at, updated_at
		FROM branches WHERE session_id = ? ORDER BY created_at ASC`, sessionID)
	if err != nil {
		return nil, fmt.Errorf("list branches: %w", err)
	}
	defer rows.Close()
	var branches []*BranchRecord
	for rows.Next() {
		b, err := scanBranch(rows)
		if err != nil {
			return nil, err
		}
		branches = append(branches, b)
	}
	return branches, rows.Err()
}

func (r *BranchRepo) DeleteBySession(ctx context.Context, sessionID string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM branches WHERE session_id = ?`, sessionID)
	return err
}

func (r *BranchRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM branches WHERE id = ?`, id)
	return err
}

func scanBranch(s interface {
	Scan(dest ...any) error
}) (*BranchRecord, error) {
	var b BranchRecord
	var historyStr, createdAt, updatedAt string
	err := s.Scan(&b.ID, &b.SessionID, &b.Name, &historyStr, &createdAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("scan branch: %w", err)
	}
	if err := json.Unmarshal([]byte(historyStr), &b.History); err != nil {
		return nil, fmt.Errorf("unmarshal history: %w", err)
	}
	b.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	b.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
	return &b, nil
}
