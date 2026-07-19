package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"termcode/internal/application/util"
)

type EditOp struct {
	OldStr string `json:"old_str"`
	NewStr string `json:"new_str"`
}

type EditResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Changes int    `json:"changes"`
}

type EditService struct{}

func NewEditService() *EditService {
	return &EditService{}
}

func (s *EditService) Edit(path string, edits []EditOp) (*EditResult, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}

	data, err := os.ReadFile(abs)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	content := string(data)
	totalChanges := 0

	for _, edit := range edits {
		if edit.OldStr == "" {
			continue
		}
		count := strings.Count(content, edit.OldStr)
		if count == 0 {
			return &EditResult{
				Success: false,
				Message: fmt.Sprintf("text not found in %s: %q", path, util.Truncate(edit.OldStr, 50)),
			}, nil
		}
		content = strings.ReplaceAll(content, edit.OldStr, edit.NewStr)
		totalChanges += count
	}

	if err := os.WriteFile(abs, []byte(content), 0o644); err != nil {
		return nil, fmt.Errorf("write file: %w", err)
	}

	return &EditResult{
		Success: true,
		Message: fmt.Sprintf("Applied %d changes to %s", totalChanges, abs),
		Changes: totalChanges,
	}, nil
}
