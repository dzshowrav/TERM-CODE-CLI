package backup

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Service struct {
	dbPath    string
	backupDir string
}

func NewService(dbPath, backupDir string) *Service {
	return &Service{
		dbPath:    dbPath,
		backupDir: backupDir,
	}
}

func (s *Service) Create(ctx context.Context) (string, error) {
	if err := os.MkdirAll(s.backupDir, 0o755); err != nil {
		return "", fmt.Errorf("create backup dir: %w", err)
	}

	timestamp := time.Now().Format("20060102_150405")
	backupPath := filepath.Join(s.backupDir, fmt.Sprintf("tc_%s.db", timestamp))

	src, err := os.Open(s.dbPath)
	if err != nil {
		return "", fmt.Errorf("open source: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(backupPath)
	if err != nil {
		return "", fmt.Errorf("create backup: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("copy: %w", err)
	}

	return backupPath, nil
}

func (s *Service) List(ctx context.Context) ([]string, error) {
	entries, err := os.ReadDir(s.backupDir)
	if err != nil {
		return nil, fmt.Errorf("read backup dir: %w", err)
	}

	var backups []string
	for _, e := range entries {
		if !e.IsDir() && filepath.Ext(e.Name()) == ".db" {
			backups = append(backups, filepath.Join(s.backupDir, e.Name()))
		}
	}

	return backups, nil
}

func (s *Service) Restore(ctx context.Context, backupPath string) error {
	src, err := os.Open(backupPath)
	if err != nil {
		return fmt.Errorf("open backup: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(s.dbPath)
	if err != nil {
		return fmt.Errorf("create db: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	return nil
}
