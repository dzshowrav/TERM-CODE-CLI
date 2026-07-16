package file

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type ReadService struct{}

func NewReadService() *ReadService {
	return &ReadService{}
}

func (s *ReadService) Read(ctx context.Context, path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("resolve path: %w", err)
	}

	data, err := os.ReadFile(abs)
	if err != nil {
		return "", fmt.Errorf("read file: %w", err)
	}

	return string(data), nil
}

func (s *ReadService) ReadLines(ctx context.Context, path string, start, end int) ([]string, error) {
	data, err := s.Read(ctx, path)
	if err != nil {
		return nil, err
	}

	lines := splitLines(data)
	if start < 0 {
		start = 0
	}
	if end > len(lines) {
		end = len(lines)
	}
	if start >= len(lines) {
		return nil, nil
	}

	return lines[start:end], nil
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}
