package tool

import (
	"fmt"
	"strings"
	"sync"
)

type RegistryConflict struct {
	Name     string
	Existing string
	Incoming string
}

type Registry struct {
	mu    sync.RWMutex
	tools []Tool
	index map[string]int
}

func NewRegistry() *Registry {
	return &Registry{
		index: make(map[string]int),
	}
}

func (r *Registry) Register(t Tool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if idx, ok := r.index[t.Name]; ok {
		existing := r.tools[idx]
		if existing.Source != t.Source {
			return fmt.Errorf("conflict: tool %q already registered from source %q (incoming: %q)",
				t.Name, existing.Source, t.Source)
		}
		r.tools[idx] = t
	} else {
		r.index[t.Name] = len(r.tools)
		r.tools = append(r.tools, t)
	}

	for _, alias := range t.Aliases {
		r.index[alias] = r.index[t.Name]
	}
	return nil
}

func (r *Registry) RegisterAll(ts []Tool) error {
	var errs []string
	for _, t := range ts {
		if err := r.Register(t); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("registration errors:\n%s", strings.Join(errs, "\n"))
	}
	return nil
}

func (r *Registry) Remove(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if idx, ok := r.index[name]; ok {
		primary := r.tools[idx].Name
		delete(r.index, primary)
		for _, alias := range r.tools[idx].Aliases {
			delete(r.index, alias)
		}
		r.tools = append(r.tools[:idx], r.tools[idx+1:]...)
		for i := idx; i < len(r.tools); i++ {
			r.index[r.tools[i].Name] = i
			for _, alias := range r.tools[i].Aliases {
				r.index[alias] = i
			}
		}
	}
}

func (r *Registry) Lookup(name string) (Tool, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if idx, ok := r.index[name]; ok {
		return r.tools[idx], true
	}
	return Tool{}, false
}

func (r *Registry) List() []Tool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]Tool, len(r.tools))
	copy(result, r.tools)
	return result
}

func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.tools)
}

func (r *Registry) FilterByCategory(cat Category) []Tool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []Tool
	for _, t := range r.tools {
		if t.Category == cat {
			result = append(result, t)
		}
	}
	return result
}

func (r *Registry) FilterByCapability(cap Capability) []Tool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []Tool
	for _, t := range r.tools {
		for _, c := range t.Capabilities {
			if c == cap {
				result = append(result, t)
				break
			}
		}
	}
	return result
}
