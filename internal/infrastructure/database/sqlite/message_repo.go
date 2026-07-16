package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"termcode/internal/domain/session"
)

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) Create(ctx context.Context, m *session.Message) error {
	_, err := r.db.ExecContext(
		ctx, `
		INSERT INTO messages (id, session_id, role, content, tool_call, tool_result, tokens_in, tokens_out, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		m.ID, m.SessionID, string(m.Role), m.Content, m.ToolCall, m.ToolRes, m.TokenIn, m.TokenOut,
		m.CreatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create message: %w", err)
	}
	return nil
}

func (r *MessageRepo) ListBySession(ctx context.Context, sessionID string) ([]*session.Message, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, session_id, role, content, tool_call, tool_result, tokens_in, tokens_out, created_at
		FROM messages WHERE session_id = ? ORDER BY created_at ASC`, sessionID)
	if err != nil {
		return nil, fmt.Errorf("list messages: %w", err)
	}
	defer rows.Close()

	var messages []*session.Message
	for rows.Next() {
		m, err := scanMessage(rows)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, rows.Err()
}

func (r *MessageRepo) DeleteBySession(ctx context.Context, sessionID string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM messages WHERE session_id = ?`, sessionID)
	return err
}

func (r *MessageRepo) CountBySession(ctx context.Context, sessionID string) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM messages WHERE session_id = ?`, sessionID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count messages: %w", err)
	}
	return count, nil
}

func (r *MessageRepo) GetLastBySession(ctx context.Context, sessionID string) (*session.Message, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, session_id, role, content, tool_call, tool_result, tokens_in, tokens_out, created_at
		FROM messages WHERE session_id = ? ORDER BY created_at DESC LIMIT 1`, sessionID)

	return scanMessage(row)
}

func scanMessage(s interface {
	Scan(dest ...any) error
},
) (*session.Message, error) {
	var msg session.Message
	var role, createdAt string

	err := s.Scan(&msg.ID, &msg.SessionID, &role, &msg.Content, &msg.ToolCall, &msg.ToolRes,
		&msg.TokenIn, &msg.TokenOut, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, session.ErrNotFound
		}
		return nil, fmt.Errorf("scan message: %w", err)
	}

	msg.Role = session.Role(role)
	msg.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)

	return &msg, nil
}
