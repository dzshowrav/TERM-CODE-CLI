package executor

import (
	"os/exec"
	"strings"
)

type FileSearcher struct{}

func NewFileSearcher() *FileSearcher {
	return &FileSearcher{}
}

func (s *FileSearcher) Search(root, pattern string) ([]string, error) {
	cmd := exec.Command("rg", "-l", pattern, root)
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return []string{}, nil
		}
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return []string{}, nil
	}
	return lines, nil
}
