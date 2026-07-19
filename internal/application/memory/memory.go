package memory

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

type Fact struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	SessionID string    `json:"session_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AccessCnt int       `json:"access_cnt"`
}

type SearchResult struct {
	Fact  Fact
	Score float64
}

type Store struct {
	mu    sync.RWMutex
	facts map[string]Fact
}

func New() *Store {
	return &Store{
		facts: make(map[string]Fact),
	}
}

func (s *Store) Add(fact Fact) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if fact.ID == "" {
		return fmt.Errorf("fact ID is required")
	}
	now := time.Now()
	fact.CreatedAt = now
	fact.UpdatedAt = now
	s.facts[fact.ID] = fact
	return nil
}

func (s *Store) Get(id string) (Fact, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f, ok := s.facts[id]
	if ok {
		f.AccessCnt++
		s.facts[id] = f
	}
	return f, ok
}

func (s *Store) Update(fact Fact) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.facts[fact.ID]; !ok {
		return fmt.Errorf("fact %s not found", fact.ID)
	}
	fact.UpdatedAt = time.Now()
	s.facts[fact.ID] = fact
	return nil
}

func (s *Store) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.facts, id)
}

func (s *Store) Search(query string, limit int) []SearchResult {
	s.mu.RLock()
	defer s.mu.RUnlock()

	query = strings.ToLower(query)
	terms := strings.Fields(query)

	var results []SearchResult
	for _, f := range s.facts {
		score := 0.0
		content := strings.ToLower(f.Content)

		for _, term := range terms {
			if strings.Contains(content, term) {
				score += 1.0
			}
			for _, tag := range f.Tags {
				if strings.Contains(strings.ToLower(tag), term) {
					score += 0.5
				}
			}
		}

		if score > 0 {
			score += float64(f.AccessCnt) * 0.01
			results = append(results, SearchResult{Fact: f, Score: score})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if limit > 0 && len(results) > limit {
		results = results[:limit]
	}

	return results
}

func (s *Store) ListBySession(sessionID string) []Fact {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []Fact
	for _, f := range s.facts {
		if f.SessionID == sessionID {
			result = append(result, f)
		}
	}
	return result
}

func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.facts)
}

func (s *Store) Export() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	facts := make([]Fact, 0, len(s.facts))
	for _, f := range s.facts {
		facts = append(facts, f)
	}
	return json.Marshal(facts)
}

func (s *Store) Import(data []byte) error {
	var facts []Fact
	if err := json.Unmarshal(data, &facts); err != nil {
		return fmt.Errorf("unmarshal facts: %w", err)
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, f := range facts {
		s.facts[f.ID] = f
	}
	return nil
}
