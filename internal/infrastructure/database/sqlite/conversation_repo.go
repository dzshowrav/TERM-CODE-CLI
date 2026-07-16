package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"termcode/internal/domain/conversation"
)

type ConversationRepo struct {
	db *sql.DB
}

func NewConversationRepo(db *sql.DB) *ConversationRepo {
	return &ConversationRepo{db: db}
}

func (r *ConversationRepo) Create(ctx context.Context, c *conversation.Conversation) error {
	_, err := r.db.ExecContext(
		ctx, `
		INSERT INTO conversations (id, session_id, summary, status, message_count, tokens_in, tokens_out, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		c.ID, c.SessionID, c.Summary, string(c.Status), c.MsgCount, c.TokenIn, c.TokenOut,
		c.CreatedAt.UTC().Format(time.RFC3339), c.UpdatedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("create conversation: %w", err)
	}
	return nil
}

func (r *ConversationRepo) GetBySession(ctx context.Context, sessionID string) (*conversation.Conversation, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, session_id, summary, status, message_count, tokens_in, tokens_out, created_at, updated_at
		FROM conversations WHERE session_id = ?`, sessionID)
	return scanConversation(row)
}

func (r *ConversationRepo) Update(ctx context.Context, c *conversation.Conversation) error {
	c.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, `
		UPDATE conversations SET summary=?, status=?, message_count=?, tokens_in=?, tokens_out=?, updated_at=?
		WHERE id=?`,
		c.Summary, string(c.Status), c.MsgCount, c.TokenIn, c.TokenOut,
		c.UpdatedAt.UTC().Format(time.RFC3339), c.ID)
	return err
}

func (r *ConversationRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM conversations WHERE id = ?`, id)
	return err
}

func scanConversation(s interface{ Scan(dest ...any) error }) (*conversation.Conversation, error) {
	var c conversation.Conversation
	var status, createdAt, updatedAt string

	err := s.Scan(&c.ID, &c.SessionID, &c.Summary, &status, &c.MsgCount, &c.TokenIn, &c.TokenOut, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, conversation.ErrNotFound
		}
		return nil, fmt.Errorf("scan conversation: %w", err)
	}

	c.Status = conversation.Status(status)
	c.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	c.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &c, nil
}
