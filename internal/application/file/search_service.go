package file

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type SearchService struct{}

func NewSearchService() *SearchService {
	return &SearchService{}
}

func (s *SearchService) Search(ctx context.Context, root, pattern string) ([]string, error) {
	cmd := exec.CommandContext(ctx, "rg", "--no-heading", "--line-number", pattern, root)
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if len(exitErr.Stderr) > 0 {
				return nil, fmt.Errorf("rg error: %s", string(exitErr.Stderr))
			}
		}
	}
	output := strings.TrimSpace(string(out))
	if output == "" {
		return nil, nil
	}
	return strings.Split(output, "\n"), nil
}

func (s *SearchService) FindFiles(ctx context.Context, root, pattern string) ([]string, error) {
	cmd := exec.CommandContext(ctx, "fd", pattern, root)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("fd error: %w", err)
	}
	output := strings.TrimSpace(string(out))
	if output == "" {
		return nil, nil
	}
	return strings.Split(output, "\n"), nil
}
