package file

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ReplaceOp struct {
	Pattern string `json:"pattern"`
	NewText string `json:"new_text"`
	IsRegex bool   `json:"is_regex"`
}

type ReplaceResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Changes int    `json:"changes"`
}

type ReplaceService struct{}

func NewReplaceService() *ReplaceService {
	return &ReplaceService{}
}

func (s *ReplaceService) Replace(path string, ops []ReplaceOp) (*ReplaceResult, error) {
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

	for _, op := range ops {
		if op.Pattern == "" {
			continue
		}
		if op.IsRegex {
			re, err := regexp.Compile(op.Pattern)
			if err != nil {
				return &ReplaceResult{
					Success: false,
					Message: fmt.Sprintf("invalid regex %q: %s", op.Pattern, err),
				}, nil
			}
			matches := re.FindAllString(content, -1)
			if len(matches) == 0 {
				return &ReplaceResult{
					Success: false,
					Message: fmt.Sprintf("regex %q not found in %s", op.Pattern, path),
				}, nil
			}
			content = re.ReplaceAllString(content, op.NewText)
			totalChanges += len(matches)
		} else {
			count := strings.Count(content, op.Pattern)
			if count == 0 {
				return &ReplaceResult{
					Success: false,
					Message: fmt.Sprintf("text %q not found in %s", op.Pattern, path),
				}, nil
			}
			content = strings.ReplaceAll(content, op.Pattern, op.NewText)
			totalChanges += count
		}
	}

	if err := os.WriteFile(abs, []byte(content), 0o644); err != nil {
		return nil, fmt.Errorf("write file: %w", err)
	}

	return &ReplaceResult{
		Success: true,
		Message: fmt.Sprintf("Applied %d replacements to %s", totalChanges, abs),
		Changes: totalChanges,
	}, nil
}

func (s *ReplaceService) ReplaceInFiles(root string, pattern string, newText string, isRegex bool, filePattern string) (*ReplaceResult, error) {
	matches, err := filepath.Glob(filepath.Join(root, filePattern))
	if err != nil {
		return nil, fmt.Errorf("glob: %w", err)
	}

	if len(matches) == 0 {
		return &ReplaceResult{
			Success: true,
			Message: "No files matched pattern.",
		}, nil
	}

	totalChanges := 0
	var failed []string
	op := ReplaceOp{Pattern: pattern, NewText: newText, IsRegex: isRegex}

	for _, m := range matches {
		res, err := s.Replace(m, []ReplaceOp{op})
		if err != nil {
			failed = append(failed, fmt.Sprintf("%s: %s", m, err))
			continue
		}
		if !res.Success {
			continue
		}
		totalChanges += res.Changes
	}

	result := &ReplaceResult{
		Success: true,
		Message: fmt.Sprintf("Applied %d replacements across files", totalChanges),
		Changes: totalChanges,
	}
	if len(failed) > 0 {
		result.Message += "; failures: " + strings.Join(failed, ", ")
	}
	return result, nil
}
