package offline

import (
	"sync"
	"time"
)

type CacheEntry struct {
	Data      string
	CreatedAt time.Time
	TTL       time.Duration
}

type Cache struct {
	mu    sync.RWMutex
	store map[string]CacheEntry
}

func New() *Cache {
	return &Cache{
		store: make(map[string]CacheEntry),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.store[key]
	if !ok {
		return "", false
	}
	if entry.TTL > 0 && time.Since(entry.CreatedAt) > entry.TTL {
		delete(c.store, key)
		return "", false
	}
	return entry.Data, true
}

func (c *Cache) Set(key, data string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = CacheEntry{
		Data:      data,
		CreatedAt: time.Now(),
		TTL:       ttl,
	}
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store = make(map[string]CacheEntry)
}

func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.store)
}

type DegradationState int

const (
	DegradationNone DegradationState = iota
	DegradationPartial
	DegradationFull
)

type GracefulDegradation struct {
	mu    sync.RWMutex
	state DegradationState
}

func NewDegradation() *GracefulDegradation {
	return &GracefulDegradation{state: DegradationNone}
}

func (d *GracefulDegradation) SetState(s DegradationState) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.state = s
}

func (d *GracefulDegradation) State() DegradationState {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.state
}

func (d *GracefulDegradation) IsDegraded() bool {
	return d.State() != DegradationNone
}
