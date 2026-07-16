package file

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

type DiffService struct{}

func NewDiffService() *DiffService {
	return &DiffService{}
}

func (s *DiffService) DiffFiles(pathA, pathB string) (string, error) {
	absA, err := filepath.Abs(pathA)
	if err != nil {
		return "", fmt.Errorf("resolve %s: %w", pathA, err)
	}
	absB, err := filepath.Abs(pathB)
	if err != nil {
		return "", fmt.Errorf("resolve %s: %w", pathB, err)
	}

	dataA, err := os.ReadFile(absA)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", pathA, err)
	}
	dataB, err := os.ReadFile(absB)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", pathB, err)
	}

	edits := myers.ComputeEdits(span.URIFromPath(pathA), string(dataA), string(dataB))
	diff := gotextdiff.ToUnified(pathA, pathB, string(dataA), edits)
	return fmt.Sprint(diff), nil
}

func (s *DiffService) DiffTexts(nameA, nameB, textA, textB string) string {
	edits := myers.ComputeEdits(span.URIFromPath(nameA), textA, textB)
	diff := gotextdiff.ToUnified(nameA, nameB, textA, edits)
	return fmt.Sprint(diff)
}
