package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Read(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("abs path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("read file: %w", err)
	}

	return string(data), nil
}

func (r *Reader) ReadLines(path string, start, end int) ([]string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("abs path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	content := string(data)
	scanner := make([]string, 0)

	line := ""
	for _, ch := range content {
		if ch == '\n' {
			scanner = append(scanner, line)
			line = ""
		} else {
			line += string(ch)
		}
	}
	if line != "" {
		scanner = append(scanner, line)
	}

	if start < 0 {
		start = 0
	}
	if end > len(scanner) {
		end = len(scanner)
	}
	if end <= start {
		return []string{}, nil
	}

	return scanner[start:end], nil
}

func (r *Reader) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (r *Reader) Size(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
