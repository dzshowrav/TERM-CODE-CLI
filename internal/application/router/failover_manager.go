package router

import (
	"context"
	"log/slog"
	"time"
)

type FailoverManager struct {
	logger *slog.Logger
}

func NewFailoverManager(logger *slog.Logger) *FailoverManager {
	return &FailoverManager{logger: logger.With("svc", "failover")}
}

type FailoverStrategy int

const (
	StrategySequential FailoverStrategy = iota
	StrategyPriority
	StrategyRandom
)

func (m *FailoverManager) SelectProvider(ctx context.Context, providers []ProviderInfo, strategy FailoverStrategy) (*ProviderInfo, error) {
	if len(providers) == 0 {
		return nil, nil
	}

	switch strategy {
	case StrategyPriority:
		best := providers[0]
		return &best, nil
	case StrategyRandom:
		return &providers[0], nil
	default:
		return &providers[0], nil
	}
}

func (m *FailoverManager) WithRetry(ctx context.Context, fn func(ctx context.Context) error, attempts int) error {
	var lastErr error
	for i := 0; i < attempts; i++ {
		if err := fn(ctx); err != nil {
			lastErr = err
			m.logger.Warn("retry failed", "attempt", i+1, "error", err)
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}
		return nil
	}
	return lastErr
}
