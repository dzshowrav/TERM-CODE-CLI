package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	domainplugin "termcode/internal/domain/plugin"
)

type Manager struct {
	mu      sync.RWMutex
	plugins map[string]*domainplugin.Plugin
	dir     string
}

func NewManager(dir string) *Manager {
	return &Manager{
		plugins: make(map[string]*domainplugin.Plugin),
		dir:     dir,
	}
}

func (m *Manager) Load(path string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	p := &domainplugin.Plugin{
		Name:     filepath.Base(path),
		Version:  "1.0.0",
		FilePath: path,
	}

	m.plugins[p.Name] = p
	return nil
}

func (m *Manager) Unload(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.plugins[name]; !ok {
		return fmt.Errorf("plugin not loaded: %s", name)
	}
	delete(m.plugins, name)
	return nil
}

func (m *Manager) Get(name string) (*domainplugin.Plugin, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	p, ok := m.plugins[name]
	return p, ok
}

func (m *Manager) List() []*domainplugin.Plugin {
	m.mu.RLock()
	defer m.mu.RUnlock()
	result := make([]*domainplugin.Plugin, 0, len(m.plugins))
	for _, p := range m.plugins {
		result = append(result, p)
	}
	return result
}

func (m *Manager) Execute(name string, args []string) (string, error) {
	plugin, ok := m.Get(name)
	if !ok {
		return "", fmt.Errorf("plugin not found: %s", name)
	}

	cmd := exec.Command(plugin.FilePath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("execute plugin %s: %w", name, err)
	}

	return string(output), nil
}

func (m *Manager) LoadAll() error {
	entries, err := os.ReadDir(m.dir)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(m.dir, 0o755)
		}
		return fmt.Errorf("read plugin dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(m.dir, entry.Name())
		if err := m.Load(path); err != nil {
			return fmt.Errorf("load plugin %s: %w", entry.Name(), err)
		}
	}

	return nil
}
