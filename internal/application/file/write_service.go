package file

import (
	"fmt"
	"os"
	"path/filepath"
)

type WriteService struct{}

func NewWriteService() *WriteService {
	return &WriteService{}
}

func (s *WriteService) Write(path, content string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("resolve path: %w", err)
	}

	dir := filepath.Dir(abs)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create directories: %w", err)
	}

	if err := os.WriteFile(abs, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
