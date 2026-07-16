package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type CacheRepo struct {
	db *sql.DB
}

func NewCacheRepo(db *sql.DB) *CacheRepo {
	return &CacheRepo{db: db}
}

func (r *CacheRepo) Get(ctx context.Context, key string) (string, error) {
	var value string
	err := r.db.QueryRowContext(ctx, `SELECT value FROM cache WHERE key = ? AND (expires_at IS NULL OR expires_at > datetime('now'))`, key).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", fmt.Errorf("get cache: %w", err)
	}
	return value, nil
}

func (r *CacheRepo) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	var expiresAt *string
	if ttl > 0 {
		e := time.Now().Add(ttl).UTC().Format(time.RFC3339)
		expiresAt = &e
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO cache (key, value, expires_at, created_at) VALUES (?, ?, ?, ?)
		ON CONFLICT(key) DO UPDATE SET value=excluded.value, expires_at=excluded.expires_at`,
		key, value, expiresAt, time.Now().UTC().Format(time.RFC3339))
	return err
}

func (r *CacheRepo) Delete(ctx context.Context, key string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM cache WHERE key = ?`, key)
	return err
}

func (r *CacheRepo) Clear(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM cache`)
	return err
}

func (r *CacheRepo) CleanExpired(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM cache WHERE expires_at IS NOT NULL AND expires_at < datetime('now')`)
	return err
}
