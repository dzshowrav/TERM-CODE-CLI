package cache

import (
	"context"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type Service struct {
	cache *gocache.Cache
}

func NewService(defaultExpiration, cleanupInterval time.Duration) *Service {
	return &Service{
		cache: gocache.New(defaultExpiration, cleanupInterval),
	}
}

func (s *Service) Get(ctx context.Context, key string) (any, bool) {
	return s.cache.Get(key)
}

func (s *Service) Set(ctx context.Context, key string, value any, ttl time.Duration) {
	s.cache.Set(key, value, ttl)
}

func (s *Service) Delete(ctx context.Context, key string) {
	s.cache.Delete(key)
}

func (s *Service) GetOrSet(ctx context.Context, key string, fn func() (any, error), ttl time.Duration) (any, error) {
	if val, ok := s.cache.Get(key); ok {
		return val, nil
	}

	val, err := fn()
	if err != nil {
		return nil, err
	}

	s.cache.Set(key, val, ttl)
	return val, nil
}

func (s *Service) Flush(ctx context.Context) {
	s.cache.Flush()
}

func (s *Service) ItemCount(ctx context.Context) int {
	return s.cache.ItemCount()
}
