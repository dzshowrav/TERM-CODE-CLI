package workspace

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type TrustManager struct {
	mu      sync.RWMutex
	trusted map[string]bool
}

func NewTrustManager() *TrustManager {
	return &TrustManager{
		trusted: make(map[string]bool),
	}
}

func (tm *TrustManager) IsTrusted(path string) bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	abs, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	return tm.trusted[abs]
}

func (tm *TrustManager) MarkTrusted(path string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	abs, err := filepath.Abs(path)
	if err != nil {
		return
	}
	tm.trusted[abs] = true
}

func (tm *TrustManager) Revoke(path string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	abs, err := filepath.Abs(path)
	if err != nil {
		return
	}
	delete(tm.trusted, abs)
}

func (tm *TrustManager) IsDangerous(path string) bool {
	entries, err := os.ReadDir(path)
	if err != nil {
		return false
	}
	for _, entry := range entries {
		name := entry.Name()
		switch {
		case !entry.IsDir() && isExecutableExt(name):
			return true
		case strings.HasPrefix(name, ".git") && entry.IsDir():
			continue
		case !entry.IsDir() && isConfigFile(name):
			return true
		}
	}
	return false
}

func isExecutableExt(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".sh", ".bash", ".zsh", ".py", ".pl", ".rb", ".exe", ".bin", ".bat", ".cmd":
		return true
	}
	return false
}

func isConfigFile(name string) bool {
	name = strings.ToLower(name)
	switch name {
	case "ssh_config", "config", ".env", ".envrc", ".gitconfig":
		return true
	}
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".key", ".pem", ".crt", ".cert", ".secret":
		return true
	}
	return false
}
