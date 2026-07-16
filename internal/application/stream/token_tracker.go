package stream

import (
	"sync"
	"time"
)

type Usage struct {
	InputTokens      int     `json:"input_tokens"`
	OutputTokens     int     `json:"output_tokens"`
	TotalTokens      int     `json:"total_tokens"`
	ContextWindow    int     `json:"context_window"`
	RemainingContext int     `json:"remaining_context"`
	ContextPercent   float64 `json:"context_percent"`
	EstimatedCost    float64 `json:"estimated_cost"`
}

type SessionTotals struct {
	TotalInputTokens  int           `json:"total_input"`
	TotalOutputTokens int           `json:"total_output"`
	TotalCost         float64       `json:"total_cost"`
	RequestCount      int           `json:"request_count"`
	SessionDuration   time.Duration `json:"session_duration"`
}

type Tracker struct {
	mu           sync.RWMutex
	usage        Usage
	session      SessionTotals
	sessionStart time.Time
	pricingIn    float64
	pricingOut   float64
}

func NewTracker(contextWindow int, pricingIn, pricingOut float64) *Tracker {
	return &Tracker{
		usage: Usage{
			ContextWindow:    contextWindow,
			RemainingContext: contextWindow,
		},
		sessionStart: time.Now(),
		pricingIn:    pricingIn,
		pricingOut:   pricingOut,
	}
}

func (t *Tracker) TrackInput(tokens int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.usage.InputTokens += tokens
	t.usage.TotalTokens = t.usage.InputTokens + t.usage.OutputTokens
	t.usage.RemainingContext = t.usage.ContextWindow - t.usage.TotalTokens
	if t.usage.RemainingContext < 0 {
		t.usage.RemainingContext = 0
	}
	t.usage.ContextPercent = float64(t.usage.TotalTokens) / float64(t.usage.ContextWindow)
	if t.usage.ContextPercent > 1.0 {
		t.usage.ContextPercent = 1.0
	}
	t.usage.EstimatedCost = calculateCost(t.usage.InputTokens, t.usage.OutputTokens, t.pricingIn, t.pricingOut)

	t.session.TotalInputTokens += tokens
	t.session.RequestCount++
}

func (t *Tracker) TrackOutput(delta int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.usage.OutputTokens += delta
	t.usage.TotalTokens = t.usage.InputTokens + t.usage.OutputTokens
	t.usage.RemainingContext = t.usage.ContextWindow - t.usage.TotalTokens
	if t.usage.RemainingContext < 0 {
		t.usage.RemainingContext = 0
	}
	t.usage.ContextPercent = float64(t.usage.TotalTokens) / float64(t.usage.ContextWindow)
	if t.usage.ContextPercent > 1.0 {
		t.usage.ContextPercent = 1.0
	}
	t.usage.EstimatedCost = calculateCost(t.usage.InputTokens, t.usage.OutputTokens, t.pricingIn, t.pricingOut)

	t.session.TotalOutputTokens += delta
}

func (t *Tracker) GetUsage() Usage {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.usage
}

func (t *Tracker) GetSession() SessionTotals {
	t.mu.RLock()
	defer t.mu.RUnlock()

	s := t.session
	s.SessionDuration = time.Since(t.sessionStart)
	return s
}

func (t *Tracker) Reset() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.usage = Usage{
		ContextWindow:    t.usage.ContextWindow,
		RemainingContext: t.usage.ContextWindow,
	}
	t.sessionStart = time.Now()
}

func (t *Tracker) ContextLevel() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	pct := t.usage.ContextPercent
	switch {
	case pct < 0.7:
		return 0
	case pct < 0.85:
		return 1
	default:
		return 2
	}
}

func calculateCost(inputTokens, outputTokens int, pricingIn, pricingOut float64) float64 {
	inputCost := (float64(inputTokens) / 1000.0) * pricingIn
	outputCost := (float64(outputTokens) / 1000.0) * pricingOut
	return inputCost + outputCost
}
