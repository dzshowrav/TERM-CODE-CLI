package sqlite

import (
	"database/sql"
	"fmt"
)

const schemaVersion = 1

var migrations = []struct {
	version int
	sql     string
}{
	{
		version: 1,
		sql: `
		CREATE TABLE IF NOT EXISTS providers (
			id            TEXT PRIMARY KEY,
			name          TEXT NOT NULL UNIQUE,
			base_url      TEXT NOT NULL,
			api_key       TEXT NOT NULL DEFAULT '',
			description   TEXT NOT NULL DEFAULT '',
			status        TEXT NOT NULL DEFAULT 'disconnected',
			latency_ms    INTEGER NOT NULL DEFAULT 0,
			priority      INTEGER NOT NULL DEFAULT 0,
			is_default    INTEGER NOT NULL DEFAULT 0,
			created_at    TEXT NOT NULL,
			updated_at    TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS models (
			id             TEXT PRIMARY KEY,
			provider_id    TEXT NOT NULL REFERENCES providers(id) ON DELETE CASCADE,
			model_id       TEXT NOT NULL,
			display_name   TEXT NOT NULL DEFAULT '',
			description    TEXT NOT NULL DEFAULT '',
			category       TEXT NOT NULL DEFAULT 'general',
			capabilities   TEXT NOT NULL DEFAULT '{}',
			max_context    INTEGER NOT NULL DEFAULT 4096,
			max_output     INTEGER NOT NULL DEFAULT 4096,
			pricing_input  REAL NOT NULL DEFAULT 0.0,
			pricing_output REAL NOT NULL DEFAULT 0.0,
			is_local       INTEGER NOT NULL DEFAULT 0,
			is_favorite    INTEGER NOT NULL DEFAULT 0,
			enabled        INTEGER NOT NULL DEFAULT 1,
			created_at     TEXT NOT NULL,
			updated_at     TEXT NOT NULL,
			UNIQUE(provider_id, model_id)
		);

		CREATE TABLE IF NOT EXISTS sessions (
			id            TEXT PRIMARY KEY,
			name          TEXT NOT NULL DEFAULT '',
			provider_id   TEXT NOT NULL,
			model_id      TEXT NOT NULL,
			agent_id      TEXT NOT NULL DEFAULT '',
			workspace     TEXT NOT NULL DEFAULT '',
			status        TEXT NOT NULL DEFAULT 'active',
			message_count INTEGER NOT NULL DEFAULT 0,
			tokens_in     INTEGER NOT NULL DEFAULT 0,
			tokens_out    INTEGER NOT NULL DEFAULT 0,
			created_at    TEXT NOT NULL,
			updated_at    TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS messages (
			id          TEXT PRIMARY KEY,
			session_id  TEXT NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
			role        TEXT NOT NULL,
			content     TEXT NOT NULL DEFAULT '',
			tool_call   TEXT NOT NULL DEFAULT '',
			tool_result TEXT NOT NULL DEFAULT '',
			tokens_in   INTEGER NOT NULL DEFAULT 0,
			tokens_out  INTEGER NOT NULL DEFAULT 0,
			created_at  TEXT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_messages_session ON messages(session_id, created_at);

		CREATE TABLE IF NOT EXISTS workspaces (
			id          TEXT PRIMARY KEY,
			name        TEXT NOT NULL UNIQUE,
			path        TEXT NOT NULL,
			is_default  INTEGER NOT NULL DEFAULT 0,
			created_at  TEXT NOT NULL,
			updated_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS agents (
			id            TEXT PRIMARY KEY,
			name          TEXT NOT NULL UNIQUE,
			description   TEXT NOT NULL DEFAULT '',
			system_prompt TEXT NOT NULL DEFAULT '',
			model_id      TEXT NOT NULL DEFAULT '',
			tools         TEXT NOT NULL DEFAULT '[]',
			is_default    INTEGER NOT NULL DEFAULT 0,
			created_at    TEXT NOT NULL,
			updated_at    TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS conversations (
			id          TEXT PRIMARY KEY,
			session_id  TEXT NOT NULL,
			summary     TEXT NOT NULL DEFAULT '',
			status      TEXT NOT NULL DEFAULT 'active',
			message_count INTEGER NOT NULL DEFAULT 0,
			tokens_in   INTEGER NOT NULL DEFAULT 0,
			tokens_out  INTEGER NOT NULL DEFAULT 0,
			created_at  TEXT NOT NULL,
			updated_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS skills (
			id          TEXT PRIMARY KEY,
			name        TEXT NOT NULL UNIQUE,
			version     TEXT NOT NULL DEFAULT '1.0',
			category    TEXT NOT NULL DEFAULT 'general',
			description TEXT NOT NULL DEFAULT '',
			path        TEXT NOT NULL DEFAULT '',
			enabled     INTEGER NOT NULL DEFAULT 1,
			is_builtin  INTEGER NOT NULL DEFAULT 0,
			created_at  TEXT NOT NULL,
			updated_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS mcp_servers (
			id          TEXT PRIMARY KEY,
			name        TEXT NOT NULL UNIQUE,
			transport   TEXT NOT NULL CHECK(transport IN ('stdio','sse','websocket')),
			command     TEXT NOT NULL DEFAULT '',
			args        TEXT NOT NULL DEFAULT '[]',
			url         TEXT NOT NULL DEFAULT '',
			env         TEXT NOT NULL DEFAULT '[]',
			status      TEXT NOT NULL DEFAULT 'disconnected',
			enabled     INTEGER NOT NULL DEFAULT 1,
			created_at  TEXT NOT NULL,
			updated_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS plugins (
			id          TEXT PRIMARY KEY,
			name        TEXT NOT NULL UNIQUE,
			version     TEXT NOT NULL DEFAULT '1.0',
			author      TEXT NOT NULL DEFAULT '',
			description TEXT NOT NULL DEFAULT '',
			status      TEXT NOT NULL DEFAULT 'inactive',
			enabled     INTEGER NOT NULL DEFAULT 1,
			installed_at TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS themes (
			id          TEXT PRIMARY KEY,
			name        TEXT NOT NULL UNIQUE,
			author      TEXT NOT NULL DEFAULT '',
			version     TEXT NOT NULL DEFAULT '1.0',
			is_dark     INTEGER NOT NULL DEFAULT 1,
			palette     TEXT NOT NULL DEFAULT '{}',
			is_active   INTEGER NOT NULL DEFAULT 0,
			created_at  TEXT NOT NULL,
			updated_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS settings (
			key         TEXT PRIMARY KEY,
			value       TEXT NOT NULL,
			updated_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS permissions (
			tool_name   TEXT PRIMARY KEY,
			permission  TEXT NOT NULL CHECK(permission IN ('always_allow','allow_once','ask','deny')),
			updated_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS cache (
			key         TEXT PRIMARY KEY,
			value       TEXT NOT NULL,
			expires_at  TEXT,
			created_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS statistics (
			id          TEXT PRIMARY KEY,
			metric      TEXT NOT NULL,
			value       REAL NOT NULL,
			recorded_at TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tool_logs (
			id          TEXT PRIMARY KEY,
			session_id  TEXT,
			tool_name   TEXT NOT NULL,
			arguments   TEXT NOT NULL DEFAULT '{}',
			result      TEXT NOT NULL DEFAULT '',
			status      TEXT NOT NULL DEFAULT 'completed',
			duration_ms INTEGER NOT NULL DEFAULT 0,
			created_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS history (
			id          TEXT PRIMARY KEY,
			type        TEXT NOT NULL,
			target_id   TEXT NOT NULL,
			target_name TEXT NOT NULL DEFAULT '',
			used_at     TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS favorites (
			id          TEXT PRIMARY KEY,
			type        TEXT NOT NULL CHECK(type IN ('model','session','command','skill','agent','file')),
			target_id   TEXT NOT NULL,
			created_at  TEXT NOT NULL,
			UNIQUE(type, target_id)
		);

		CREATE TABLE IF NOT EXISTS bookmarks (
			id          TEXT PRIMARY KEY,
			session_id  TEXT,
			message_id  TEXT,
			label       TEXT NOT NULL DEFAULT '',
			created_at  TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS attachments (
			id          TEXT PRIMARY KEY,
			session_id  TEXT,
			message_id  TEXT,
			filename    TEXT NOT NULL,
			path        TEXT NOT NULL,
			mime_type   TEXT NOT NULL DEFAULT 'text/plain',
			size_bytes  INTEGER NOT NULL DEFAULT 0,
			created_at  TEXT NOT NULL
		);

		CREATE INDEX IF NOT EXISTS idx_conversations_session ON conversations(session_id);
		CREATE INDEX IF NOT EXISTS idx_tool_logs_session ON tool_logs(session_id);
		CREATE INDEX IF NOT EXISTS idx_history_type ON history(type);
		CREATE INDEX IF NOT EXISTS idx_history_used ON history(used_at);
		CREATE INDEX IF NOT EXISTS idx_statistics_metric ON statistics(metric);
		CREATE INDEX IF NOT EXISTS idx_cache_expires ON cache(expires_at);
		CREATE INDEX IF NOT EXISTS idx_favorites_type ON favorites(type);

		CREATE TABLE IF NOT EXISTS schema_meta (
			version INTEGER PRIMARY KEY,
			applied_at TEXT NOT NULL
		);
		`,
	},
}

func RunMigrations(db *sql.DB) error {
	var currentVersion int
	err := db.QueryRow("SELECT COALESCE(MAX(version), 0) FROM schema_meta").Scan(&currentVersion)
	if err != nil {
		currentVersion = 0
	}

	for _, m := range migrations {
		if m.version > currentVersion {
			tx, err := db.Begin()
			if err != nil {
				return fmt.Errorf("begin migration %d: %w", m.version, err)
			}

			if _, err := tx.Exec(m.sql); err != nil {
				tx.Rollback()
				return fmt.Errorf("migration %d: %w", m.version, err)
			}

			if _, err := tx.Exec("INSERT INTO schema_meta (version, applied_at) VALUES (?, datetime('now'))", m.version); err != nil {
				tx.Rollback()
				return fmt.Errorf("record migration %d: %w", m.version, err)
			}

			if err := tx.Commit(); err != nil {
				return fmt.Errorf("commit migration %d: %w", m.version, err)
			}
		}
	}

	return nil
}
