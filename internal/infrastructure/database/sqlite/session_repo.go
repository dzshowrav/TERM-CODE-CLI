package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"termcode/internal/domain/session"
)

type SessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepo {
	return &SessionRepo{db: db}
}

func (r *SessionRepo) Create(ctx context.Context, s *session.Session) error {
	_, err := r.db.ExecContext(
		ctx, `
		INSERT INTO sessions (id, name, provider_id, model_id, agent_id, workspace, status, message_count, tokens_in, tokens_out, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		s.ID, s.Name, s.ProviderID, s.ModelID, "", "", string(s.Status), s.MessageCnt, s.TokenIn, s.TokenOut,
		s.CreatedAt.UTC().Format(time.RFC3339), s.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create session: %w", err)
	}
	return nil
}

func (r *SessionRepo) GetByID(ctx context.Context, id string) (*session.Session, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, name, provider_id, model_id, agent_id, workspace, status, message_count, tokens_in, tokens_out, created_at, updated_at
		FROM sessions WHERE id = ?`, id)

	return scanSession(row)
}

func (r *SessionRepo) List(ctx context.Context) ([]*session.Session, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, provider_id, model_id, agent_id, workspace, status, message_count, tokens_in, tokens_out, created_at, updated_at
		FROM sessions ORDER BY updated_at DESC`)
	if err != nil {
		return nil, fmt.Errorf("list sessions: %w", err)
	}
	defer rows.Close()

	var sessions []*session.Session
	for rows.Next() {
		s, err := scanSession(rows)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}
	return sessions, rows.Err()
}

func (r *SessionRepo) ListActive(ctx context.Context) ([]*session.Session, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, provider_id, model_id, agent_id, workspace, status, message_count, tokens_in, tokens_out, created_at, updated_at
		FROM sessions WHERE status = 'active' ORDER BY updated_at DESC`)
	if err != nil {
		return nil, fmt.Errorf("list active sessions: %w", err)
	}
	defer rows.Close()

	var sessions []*session.Session
	for rows.Next() {
		s, err := scanSession(rows)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}
	return sessions, rows.Err()
}

func (r *SessionRepo) ListByStatus(ctx context.Context, status session.Status) ([]*session.Session, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, provider_id, model_id, agent_id, workspace, status, message_count, tokens_in, tokens_out, created_at, updated_at
		FROM sessions WHERE status = ? ORDER BY updated_at DESC`, string(status))
	if err != nil {
		return nil, fmt.Errorf("list sessions by status: %w", err)
	}
	defer rows.Close()

	var sessions []*session.Session
	for rows.Next() {
		s, err := scanSession(rows)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}
	return sessions, rows.Err()
}

func (r *SessionRepo) Update(ctx context.Context, s *session.Session) error {
	s.UpdatedAt = time.Now()
	res, err := r.db.ExecContext(
		ctx, `
		UPDATE sessions SET name=?, provider_id=?, model_id=?, agent_id=?, workspace=?, status=?, message_count=?, tokens_in=?, tokens_out=?, updated_at=?
		WHERE id=?`,
		s.Name, s.ProviderID, s.ModelID, "", "", string(s.Status), s.MessageCnt, s.TokenIn, s.TokenOut,
		s.UpdatedAt.UTC().Format(time.RFC3339), s.ID,
	)
	if err != nil {
		return fmt.Errorf("update session: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return session.ErrNotFound
	}
	return nil
}

func (r *SessionRepo) Archive(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `UPDATE sessions SET status = 'archived' WHERE id = ?`, id)
	return err
}

func (r *SessionRepo) ArchiveAll(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `UPDATE sessions SET status = 'archived' WHERE status = 'active'`)
	return err
}

func (r *SessionRepo) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM sessions WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return session.ErrNotFound
	}
	return nil
}

func scanSession(s interface {
	Scan(dest ...any) error
},
) (*session.Session, error) {
	var sess session.Session
	var status, createdAt, updatedAt, agentID, workspace string

	err := s.Scan(&sess.ID, &sess.Name, &sess.ProviderID, &sess.ModelID, &agentID, &workspace,
		&status, &sess.MessageCnt, &sess.TokenIn, &sess.TokenOut, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, session.ErrNotFound
		}
		return nil, fmt.Errorf("scan session: %w", err)
	}

	sess.Status = session.Status(status)
	sess.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	sess.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &sess, nil
}
