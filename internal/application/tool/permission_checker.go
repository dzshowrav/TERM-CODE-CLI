package tool

import (
	"sync"

	"termcode/internal/domain/permission"
)

type StorePermissionChecker struct {
	mu          sync.RWMutex
	entries     map[string]permission.Level
	requestFunc func(toolName, args string, resultCh chan<- string)
	trustMgr    *TrustManager
}

func NewPermissionChecker() *StorePermissionChecker {
	return &StorePermissionChecker{
		entries: make(map[string]permission.Level),
		requestFunc: func(toolName, args string, resultCh chan<- string) {
			resultCh <- "deny"
		},
	}
}

func (c *StorePermissionChecker) SetTrustManager(tm *TrustManager) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.trustMgr = tm
}

func (c *StorePermissionChecker) SetRequestFunc(f func(toolName, args string, resultCh chan<- string)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.requestFunc = f
}

func (c *StorePermissionChecker) Set(toolName string, level permission.Level) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[toolName] = level
}

func (c *StorePermissionChecker) IsAllowed(toolName string) bool {
	c.mu.RLock()
	level, ok := c.entries[toolName]
	tm := c.trustMgr
	c.mu.RUnlock()

	if ok && (level == permission.LevelAlwaysAllow || level == permission.LevelAllowOnce) {
		return true
	}
	if tm != nil && tm.GetDecision(toolName) == "always_allow" {
		return true
	}
	return false
}

func (c *StorePermissionChecker) IsDenied(toolName string) bool {
	c.mu.RLock()
	level, ok := c.entries[toolName]
	tm := c.trustMgr
	c.mu.RUnlock()

	if ok && level == permission.LevelDeny {
		return true
	}
	if tm != nil && tm.GetDecision(toolName) == "deny" {
		return true
	}
	return false
}

func (c *StorePermissionChecker) Request(toolName, args string) string {
	resultCh := make(chan string, 1)

	c.mu.RLock()
	f := c.requestFunc
	c.mu.RUnlock()

	f(toolName, args, resultCh)

	return <-resultCh
}
