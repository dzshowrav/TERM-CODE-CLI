package file

import (
	"fmt"
	"os"
	"path/filepath"
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

type UndoService struct {
	mu        sync.Mutex
	stack     []UndoEntry
	redoStack []UndoEntry
	maxSize   int
}

func NewUndoService() *UndoService {
	return &UndoService{
		maxSize: 100,
	}
}

func (u *UndoService) SetMaxSize(n int) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.maxSize = n
}

func (u *UndoService) Record(entry UndoEntry) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.stack = append(u.stack, entry)
	if len(u.stack) > u.maxSize {
		u.stack = u.stack[1:]
	}
	u.redoStack = nil
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

func (u *UndoService) Undo() (string, error) {
	u.mu.Lock()
	if len(u.stack) == 0 {
		u.mu.Unlock()
		return "", fmt.Errorf("nothing to undo")
	}
	entry := u.stack[len(u.stack)-1]
	u.stack = u.stack[:len(u.stack)-1]
	u.mu.Unlock()

	var err error
	switch entry.Type {
	case OpWrite:
		err = u.undoWrite(entry)
	case OpEdit:
		err = u.undoEdit(entry)
	case OpReplace:
		err = u.undoEdit(entry)
	case OpDelete:
		err = u.undoDelete(entry)
	case OpMove:
		err = u.undoMove(entry)
	case OpCreate:
		err = u.undoCreate(entry)
	case OpCopy:
		err = nil
	default:
		err = fmt.Errorf("unknown op type: %s", entry.Type)
	}

	if err != nil {
		return "", err
	}

	u.mu.Lock()
	u.redoStack = append(u.redoStack, entry)
	if len(u.redoStack) > u.maxSize {
		u.redoStack = u.redoStack[1:]
	}
	u.mu.Unlock()

	return fmt.Sprintf("Undone: %s %s", entry.Type, entry.Path), nil
}

func (u *UndoService) Redo() (string, error) {
	u.mu.Lock()
	if len(u.redoStack) == 0 {
		u.mu.Unlock()
		return "", fmt.Errorf("nothing to redo")
	}
	entry := u.redoStack[len(u.redoStack)-1]
	u.redoStack = u.redoStack[:len(u.redoStack)-1]
	u.mu.Unlock()

	var err error
	switch entry.Type {
	case OpWrite:
		err = os.WriteFile(entry.Path, entry.NewData, 0o644)
	case OpEdit, OpReplace:
		err = os.WriteFile(entry.Path, entry.NewData, 0o644)
	case OpDelete:
		err = os.WriteFile(entry.Path, entry.OldData, 0o644)
	case OpMove:
		err = os.Rename(entry.Path, entry.OldPath)
	case OpCreate:
		err = os.Remove(entry.Path)
	case OpCopy:
		err = os.Remove(entry.Path)
	}

	if err != nil {
		return "", err
	}

	u.mu.Lock()
	u.stack = append(u.stack, entry)
	if len(u.stack) > u.maxSize {
		u.stack = u.stack[1:]
	}
	u.mu.Unlock()

	return fmt.Sprintf("Redone: %s %s", entry.Type, entry.Path), nil
}

func (u *UndoService) undoWrite(entry UndoEntry) error {
	if len(entry.OldData) == 0 {
		return os.Remove(entry.Path)
	}
	return os.WriteFile(entry.Path, entry.OldData, 0o644)
}

func (u *UndoService) undoEdit(entry UndoEntry) error {
	return os.WriteFile(entry.Path, entry.OldData, 0o644)
}

func (u *UndoService) undoDelete(entry UndoEntry) error {
	dir := filepath.Dir(entry.Path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(entry.Path, entry.OldData, 0o644)
}

func (u *UndoService) undoMove(entry UndoEntry) error {
	dir := filepath.Dir(entry.OldPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return os.Rename(entry.Path, entry.OldPath)
}

func (u *UndoService) undoCreate(entry UndoEntry) error {
	return os.Remove(entry.Path)
}
