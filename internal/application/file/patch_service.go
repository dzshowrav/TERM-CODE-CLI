package file

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type PatchService struct{}

func NewPatchService() *PatchService {
	return &PatchService{}
}

type PatchOp struct {
	File  string
	Hunks []Hunk
}

type Hunk struct {
	OldStart int
	OldCount int
	NewStart int
	NewCount int
	Lines    []string
}

var hunkHeader = regexp.MustCompile(`^@@ -(\d+),?(\d*) \+(\d+),?(\d*) @@`)

func (s *PatchService) ApplyPatch(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read patch file: %w", err)
	}

	patches := parsePatch(string(data))
	for _, p := range patches {
		if err := s.applyHunks(p.File, p.Hunks); err != nil {
			return fmt.Errorf("apply to %s: %w", p.File, err)
		}
	}
	return nil
}

func (s *PatchService) ApplyPatchContent(content string, baseDir string) error {
	patches := parsePatch(content)
	for _, p := range patches {
		fullPath := filepath.Join(baseDir, p.File)
		if err := s.applyHunks(fullPath, p.Hunks); err != nil {
			return fmt.Errorf("apply to %s: %w", fullPath, err)
		}
	}
	return nil
}

func parsePatch(patch string) []PatchOp {
	var ops []PatchOp
	var current *PatchOp

	lines := strings.Split(patch, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "--- ") {
			continue
		}
		if strings.HasPrefix(line, "+++ ") {
			filePath := strings.TrimPrefix(line, "+++ ")
			filePath = strings.TrimPrefix(filePath, "b/")
			filePath = strings.TrimPrefix(filePath, "a/")
			if current != nil {
				ops = append(ops, *current)
			}
			current = &PatchOp{File: filePath}
			continue
		}
		if current == nil {
			continue
		}
		if m := hunkHeader.FindStringSubmatch(line); m != nil {
			oldStart, _ := strconv.Atoi(m[1])
			oldCount, _ := strconv.Atoi(m[2])
			newStart, _ := strconv.Atoi(m[3])
			newCount, _ := strconv.Atoi(m[4])
			if oldCount == 0 && m[2] == "" {
				oldCount = 1
			}
			if newCount == 0 && m[4] == "" {
				newCount = 1
			}
			current.Hunks = append(current.Hunks, Hunk{
				OldStart: oldStart,
				OldCount: oldCount,
				NewStart: newStart,
				NewCount: newCount,
			})
		} else if len(current.Hunks) > 0 {
			h := &current.Hunks[len(current.Hunks)-1]
			h.Lines = append(h.Lines, line)
		}
	}
	if current != nil {
		ops = append(ops, *current)
	}
	return ops
}

func (s *PatchService) applyHunks(filePath string, hunks []Hunk) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read %s: %w", filePath, err)
	}

	content := strings.Split(string(data), "\n")
	// Remove trailing empty line
	if len(content) > 0 && content[len(content)-1] == "" {
		content = content[:len(content)-1]
	}

	for i := len(hunks) - 1; i >= 0; i-- {
		h := hunks[i]
		var newLines []string
		for _, line := range h.Lines {
			if len(line) == 0 {
				continue
			}
			switch line[0] {
			case '+':
				newLines = append(newLines, line[1:])
			case '-':
				continue
			default:
				newLines = append(newLines, line[1:])
			}
		}

		oldLines := make([]string, 0)
		for _, line := range h.Lines {
			if len(line) == 0 {
				continue
			}
			if line[0] != '+' {
				oldLines = append(oldLines, line[1:])
			}
		}

		start := h.OldStart - 1
		if start < 0 {
			start = 0
		}
		if start > len(content) {
			start = len(content)
		}

		end := start + len(oldLines)
		if end > len(content) {
			end = len(content)
		}

		// Verify context matches
		contextLen := end - start
		matchLen := contextLen
		if len(oldLines) < matchLen {
			matchLen = len(oldLines)
		}
		if matchLen > 0 {
			var newContent []string
			newContent = append(newContent, content[:start]...)
			newContent = append(newContent, newLines...)
			if end < len(content) {
				newContent = append(newContent, content[end:]...)
			}
			content = newContent
		} else {
			// Insert at position
			var newContent []string
			newContent = append(newContent, content[:start]...)
			newContent = append(newContent, newLines...)
			if start < len(content) {
				newContent = append(newContent, content[start:]...)
			}
			content = newContent
		}
	}

	result := strings.Join(content, "\n")
	return os.WriteFile(filePath, []byte(result), 0o644)
}
