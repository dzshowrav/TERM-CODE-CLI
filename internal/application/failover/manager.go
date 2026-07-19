package failover

import (
	"math"
	"sync"
	"time"
)

type ProviderHealth struct {
	Name      string
	Latency   time.Duration
	LastCheck time.Time
	FailCount int
	Healthy   bool
}

type FailoverManager struct {
	mu            sync.RWMutex
	providers     map[string]*ProviderHealth
	threshold     int
	cooldown      time.Duration
	checkInterval time.Duration
}

func New() *FailoverManager {
	return &FailoverManager{
		providers:     make(map[string]*ProviderHealth),
		threshold:     3,
		cooldown:      30 * time.Second,
		checkInterval: 60 * time.Second,
	}
}

func (m *FailoverManager) SetThreshold(n int)               { m.threshold = n }
func (m *FailoverManager) SetCooldown(d time.Duration)      { m.cooldown = d }
func (m *FailoverManager) SetCheckInterval(d time.Duration) { m.checkInterval = d }

func (m *FailoverManager) Register(name string) {
	m.mu.Lock()
	if _, ok := m.providers[name]; !ok {
		m.providers[name] = &ProviderHealth{
			Name:    name,
			Healthy: true,
		}
	}
	m.mu.Unlock()
}

func (m *FailoverManager) RecordSuccess(name string, latency time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	h, ok := m.providers[name]
	if !ok {
		return
	}
	h.FailCount = 0
	h.Latency = latency
	h.LastCheck = time.Now()
	h.Healthy = true
}

func (m *FailoverManager) RecordFailure(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	h, ok := m.providers[name]
	if !ok {
		return
	}
	h.FailCount++
	h.LastCheck = time.Now()
	if h.FailCount >= m.threshold {
		h.Healthy = false
	}
}

func (m *FailoverManager) IsHealthy(name string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	h, ok := m.providers[name]
	if !ok {
		return true
	}

	if !h.Healthy && time.Since(h.LastCheck) > m.cooldown {
		return true
	}

	return h.Healthy
}

func (m *FailoverManager) Next(names []string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var best string
	bestLatency := time.Duration(math.MaxInt64)

	for _, name := range names {
		h, ok := m.providers[name]
		if !ok || !h.Healthy {
			continue
		}
		if best == "" || h.Latency < bestLatency {
			best = name
			bestLatency = h.Latency
		}
	}

	if best == "" && len(names) > 0 {
		return names[0]
	}

	return best
}

func (m *FailoverManager) Health() []ProviderHealth {
	m.mu.RLock()
	defer m.mu.RUnlock()

	health := make([]ProviderHealth, 0, len(m.providers))
	for _, h := range m.providers {
		health = append(health, *h)
	}
	return health
}

func (m *FailoverManager) Recover(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if h, ok := m.providers[name]; ok {
		h.Healthy = true
		h.FailCount = 0
	}
}
