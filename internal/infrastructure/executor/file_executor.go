package executor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileExecutor struct{}

func NewFileExecutor() *FileExecutor {
	return &FileExecutor{}
}

type EditOp struct {
	OldStr string `json:"old_str"`
	NewStr string `json:"new_str"`
}

type EditResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

func (e *FileExecutor) Read(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("abs path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("read: %w", err)
	}

	return string(data), nil
}

func (e *FileExecutor) Write(path, content string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("abs path: %w", err)
	}

	dir := filepath.Dir(absPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	if err := os.WriteFile(absPath, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

func (e *FileExecutor) Edit(path string, edits []EditOp) (*EditResult, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return &EditResult{Success: false, Message: err.Error(), Path: path}, nil
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return &EditResult{Success: false, Message: fmt.Sprintf("read: %v", err), Path: path}, nil
	}

	content := string(data)

	for _, edit := range edits {
		if edit.OldStr == "" {
			return &EditResult{Success: false, Message: "old_str cannot be empty", Path: path}, nil
		}
		n := strings.Count(content, edit.OldStr)
		if n == 0 {
			return &EditResult{Success: false, Message: fmt.Sprintf("'%s' not found in file", edit.OldStr), Path: path}, nil
		}
		if n > 1 {
			return &EditResult{Success: false, Message: fmt.Sprintf("'%s' found %d times; be more specific", edit.OldStr, n), Path: path}, nil
		}
		content = strings.Replace(content, edit.OldStr, edit.NewStr, 1)
	}

	if err := os.WriteFile(absPath, []byte(content), 0o644); err != nil {
		return &EditResult{Success: false, Message: fmt.Sprintf("write: %v", err), Path: path}, nil
	}

	return &EditResult{Success: true, Message: "file edited", Path: path}, nil
}

func (e *FileExecutor) Delete(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("abs path: %w", err)
	}
	return os.Remove(absPath)
}
