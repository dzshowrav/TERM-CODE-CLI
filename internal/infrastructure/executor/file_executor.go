package executor

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type UndoOpType string

const (
	OpWrite   UndoOpType = "write"
	OpEdit    UndoOpType = "edit"
	OpReplace UndoOpType = "replace"
	OpDelete  UndoOpType = "delete"
	OpMove    UndoOpType = "move"
	OpCopy    UndoOpType = "copy"
	OpCreate  UndoOpType = "create"
)

type UndoEntry struct {
	Type      UndoOpType
	Path      string
	OldPath   string
	OldData   []byte
	NewData   []byte
	Timestamp time.Time
}

type Hunk struct {
	OldStart int
	OldCount int
	NewStart int
	NewCount int
	Lines    []string
}

type PatchOp struct {
	File  string
	Hunks []Hunk
}

var hunkHeader = regexp.MustCompile(`^@@ -(\d+),?(\d*) \+(\d+),?(\d*) @@`)

func ParsePatchContent(patch string) []PatchOp {
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

type UndoService struct {
	mu        sync.Mutex
	stack     []UndoEntry
	redoStack []UndoEntry
	maxSize   int
}

func NewUndoService() *UndoService {
	return &UndoService{maxSize: 100}
}

func (u *UndoService) Record(entry UndoEntry) {
	u.mu.Lock()
	defer u.mu.Unlock()
	entry.Timestamp = time.Now()
	u.stack = append(u.stack, entry)
	if len(u.stack) > u.maxSize {
		copy(u.stack, u.stack[1:])
		u.stack = u.stack[:u.maxSize]
	}
	u.redoStack = nil
}

func (u *UndoService) PopUndo() (UndoEntry, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	if len(u.stack) == 0 {
		return UndoEntry{}, false
	}
	entry := u.stack[len(u.stack)-1]
	u.stack = u.stack[:len(u.stack)-1]
	return entry, true
}

func (u *UndoService) PeekUndo() (UndoEntry, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	if len(u.stack) == 0 {
		return UndoEntry{}, false
	}
	return u.stack[len(u.stack)-1], true
}

func (u *UndoService) PopRedo() (UndoEntry, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	if len(u.redoStack) == 0 {
		return UndoEntry{}, false
	}
	entry := u.redoStack[len(u.redoStack)-1]
	u.redoStack = u.redoStack[:len(u.redoStack)-1]
	return entry, true
}

func (u *UndoService) PushRedo(entry UndoEntry) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.redoStack = append(u.redoStack, entry)
	if len(u.redoStack) > u.maxSize {
		copy(u.redoStack, u.redoStack[1:])
		u.redoStack = u.redoStack[:u.maxSize]
	}
}

func (u *UndoService) CanUndo() bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return len(u.stack) > 0
}

func (u *UndoService) CanRedo() bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return len(u.redoStack) > 0
}

func (u *UndoService) PushUndo(entry UndoEntry) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.stack = append(u.stack, entry)
	if len(u.stack) > u.maxSize {
		copy(u.stack, u.stack[1:])
		u.stack = u.stack[:u.maxSize]
	}
}

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

func (e *FileExecutor) Replace(path string, ops []ReplaceOp) (*ReplaceResult, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return &ReplaceResult{Success: false, Message: err.Error()}, nil
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return &ReplaceResult{Success: false, Message: fmt.Sprintf("read: %v", err)}, nil
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
				return &ReplaceResult{Success: false, Message: fmt.Sprintf("invalid regex %q: %s", op.Pattern, err)}, nil
			}
			matches := re.FindAllString(content, -1)
			if len(matches) == 0 {
				return &ReplaceResult{Success: false, Message: fmt.Sprintf("regex %q not found", op.Pattern)}, nil
			}
			content = re.ReplaceAllString(content, op.NewText)
			totalChanges += len(matches)
		} else {
			count := strings.Count(content, op.Pattern)
			if count == 0 {
				return &ReplaceResult{Success: false, Message: fmt.Sprintf("text %q not found", op.Pattern)}, nil
			}
			content = strings.ReplaceAll(content, op.Pattern, op.NewText)
			totalChanges += count
		}
	}

	if err := os.WriteFile(absPath, []byte(content), 0o644); err != nil {
		return &ReplaceResult{Success: false, Message: fmt.Sprintf("write: %v", err)}, nil
	}

	return &ReplaceResult{Success: true, Message: fmt.Sprintf("replaced %d occurrences", totalChanges), Changes: totalChanges}, nil
}

func (e *FileExecutor) Rename(oldPath, newPath string) error {
	absOld, err := filepath.Abs(oldPath)
	if err != nil {
		return fmt.Errorf("abs old: %w", err)
	}
	absNew, err := filepath.Abs(newPath)
	if err != nil {
		return fmt.Errorf("abs new: %w", err)
	}

	dir := filepath.Dir(absNew)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	return os.Rename(absOld, absNew)
}

func (e *FileExecutor) Move(source, dest string) error {
	return e.Rename(source, dest)
}

func (e *FileExecutor) Copy(source, dest string) error {
	absSrc, err := filepath.Abs(source)
	if err != nil {
		return fmt.Errorf("abs src: %w", err)
	}
	absDst, err := filepath.Abs(dest)
	if err != nil {
		return fmt.Errorf("abs dst: %w", err)
	}

	srcInfo, err := os.Stat(absSrc)
	if err != nil {
		return fmt.Errorf("stat src: %w", err)
	}

	dir := filepath.Dir(absDst)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	if srcInfo.IsDir() {
		return copyDir(absSrc, absDst)
	}
	return copyFile(absSrc, absDst)
}

func copyFile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open src: %w", err)
	}
	defer r.Close()

	w, err := os.OpenFile(dst, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("create dst: %w", err)
	}
	defer w.Close()

	if _, err := io.Copy(w, r); err != nil {
		return fmt.Errorf("copy: %w", err)
	}
	return nil
}

func copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if info.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *FileExecutor) CheckLargeFile(path string, maxSize int64) (bool, int64, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false, 0, fmt.Errorf("abs path: %w", err)
	}
	info, err := os.Stat(absPath)
	if err != nil {
		return false, 0, fmt.Errorf("stat: %w", err)
	}
	if info.Size() > maxSize {
		return true, info.Size(), nil
	}
	return false, info.Size(), nil
}

func (e *FileExecutor) ReadLinesRange(path string, startLine, endLine int) ([]string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("abs path: %w", err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	lines := strings.Split(string(data), "\n")
	if startLine < 0 {
		startLine = 0
	}
	if endLine > len(lines) {
		endLine = len(lines)
	}
	if startLine >= len(lines) {
		return nil, nil
	}
	return lines[startLine:endLine], nil
}

func (e *FileExecutor) ReadHead(path string, maxLines int) ([]string, error) {
	return e.ReadLinesRange(path, 0, maxLines)
}

func (e *FileExecutor) ReadTail(path string, maxLines int) ([]string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("abs path: %w", err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) <= maxLines {
		return lines, nil
	}
	return lines[len(lines)-maxLines:], nil
}
