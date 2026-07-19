package review

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
	SeverityInfo    Severity = "info"
)

type Finding struct {
	File       string   `json:"file"`
	Line       int      `json:"line"`
	Column     int      `json:"column"`
	Severity   Severity `json:"severity"`
	Message    string   `json:"message"`
	Category   string   `json:"category"`
	Suggestion string   `json:"suggestion,omitempty"`
}

type Review struct {
	ID        string    `json:"id"`
	Target    string    `json:"target"`
	Status    string    `json:"status"`
	Findings  []Finding `json:"findings"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Engine struct {
	mu      sync.Mutex
	reviews map[string]*Review
}

func NewEngine() *Engine {
	return &Engine{
		reviews: make(map[string]*Review),
	}
}

func (e *Engine) CreateReview(target string) *Review {
	e.mu.Lock()
	defer e.mu.Unlock()

	r := &Review{
		ID:        fmt.Sprintf("review_%d", time.Now().UnixNano()),
		Target:    target,
		Status:    "pending",
		Findings:  make([]Finding, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	e.reviews[r.ID] = r
	return r
}

func (e *Engine) GetReview(id string) (*Review, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	r, ok := e.reviews[id]
	return r, ok
}

func (e *Engine) AddFinding(reviewID string, finding Finding) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	r, ok := e.reviews[reviewID]
	if !ok {
		return fmt.Errorf("review not found: %s", reviewID)
	}
	r.Findings = append(r.Findings, finding)
	r.UpdatedAt = time.Now()
	return nil
}

func (e *Engine) SetSummary(reviewID, summary string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	r, ok := e.reviews[reviewID]
	if !ok {
		return fmt.Errorf("review not found: %s", reviewID)
	}
	r.Summary = summary
	r.UpdatedAt = time.Now()
	return nil
}

func (e *Engine) SetStatus(reviewID, status string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	r, ok := e.reviews[reviewID]
	if !ok {
		return fmt.Errorf("review not found: %s", reviewID)
	}
	r.Status = status
	r.UpdatedAt = time.Now()
	return nil
}

func (e *Engine) ListReviews() []*Review {
	e.mu.Lock()
	defer e.mu.Unlock()
	result := make([]*Review, 0, len(e.reviews))
	for _, r := range e.reviews {
		result = append(result, r)
	}
	return result
}

func (r *Review) SummaryText() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Review: %s\n", r.Target))
	b.WriteString(fmt.Sprintf("Status: %s\n", r.Status))
	b.WriteString(fmt.Sprintf("Findings: %d\n", len(r.Findings)))
	b.WriteString(strings.Repeat("-", 40) + "\n")

	errs := 0
	warns := 0
	infos := 0
	for _, f := range r.Findings {
		switch f.Severity {
		case SeverityError:
			errs++
		case SeverityWarning:
			warns++
		case SeverityInfo:
			infos++
		}
	}
	b.WriteString(fmt.Sprintf("Errors: %d  Warnings: %d  Info: %d\n", errs, warns, infos))

	if r.Summary != "" {
		b.WriteString("\nSummary:\n" + r.Summary + "\n")
	}

	return b.String()
}

func (r *Review) FindingsByCategory(category string) []Finding {
	var result []Finding
	for _, f := range r.Findings {
		if f.Category == category {
			result = append(result, f)
		}
	}
	return result
}
