package tool

import (
	"context"
	"strings"
	"sync"

	"termcode/internal/infrastructure/database/sqlite"
)

type TrustManager struct {
	mu           sync.RWMutex
	settingsRepo *sqlite.SettingsRepo
	decisions    map[string]string
}

func NewTrustManager(settingsRepo *sqlite.SettingsRepo) *TrustManager {
	tm := &TrustManager{
		settingsRepo: settingsRepo,
		decisions:    make(map[string]string),
	}
	tm.LoadDecisions()
	return tm
}

func (tm *TrustManager) LoadDecisions() {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	ctx := context.Background()
	settings, err := tm.settingsRepo.List(ctx)
	if err != nil {
		return
	}
	for k, v := range settings {
		if strings.HasPrefix(k, "trust_") {
			tm.decisions[strings.TrimPrefix(k, "trust_")] = v
		}
	}
}

func (tm *TrustManager) SaveDecision(toolName, decision string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	ctx := context.Background()
	tm.decisions[toolName] = decision
	return tm.settingsRepo.Set(ctx, "trust_"+toolName, decision)
}

func (tm *TrustManager) GetDecision(toolName string) string {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.decisions[toolName]
}
