package cost

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type ModelPricing struct {
	InputPricePer1K  float64
	OutputPricePer1K float64
	Currency         string
}

type Record struct {
	Model     string
	InputTok  int
	OutputTok int
	Cost      float64
	Currency  string
	Timestamp time.Time
}

type Engine struct {
	mu       sync.RWMutex
	pricing  map[string]ModelPricing
	records  []Record
	budget   float64
	currency string
}

func New() *Engine {
	return &Engine{
		pricing: map[string]ModelPricing{
			"gpt-4":           {InputPricePer1K: 0.03, OutputPricePer1K: 0.06, Currency: "USD"},
			"gpt-4-turbo":     {InputPricePer1K: 0.01, OutputPricePer1K: 0.03, Currency: "USD"},
			"gpt-3.5-turbo":   {InputPricePer1K: 0.001, OutputPricePer1K: 0.002, Currency: "USD"},
			"claude-3-opus":   {InputPricePer1K: 0.015, OutputPricePer1K: 0.075, Currency: "USD"},
			"claude-3-sonnet": {InputPricePer1K: 0.003, OutputPricePer1K: 0.015, Currency: "USD"},
			"claude-3-haiku":  {InputPricePer1K: 0.00025, OutputPricePer1K: 0.00125, Currency: "USD"},
			"llama3":          {InputPricePer1K: 0.0, OutputPricePer1K: 0.0, Currency: "USD"},
			"codestral":       {InputPricePer1K: 0.001, OutputPricePer1K: 0.003, Currency: "USD"},
			"deepseek-coder":  {InputPricePer1K: 0.0014, OutputPricePer1K: 0.0028, Currency: "USD"},
		},
		currency: "USD",
	}
}

func (e *Engine) SetPricing(model string, inputPrice, outputPrice float64, currency string) {
	e.mu.Lock()
	e.pricing[model] = ModelPricing{
		InputPricePer1K:  inputPrice,
		OutputPricePer1K: outputPrice,
		Currency:         currency,
	}
	e.mu.Unlock()
}

func (e *Engine) SetBudget(b float64) {
	e.mu.Lock()
	e.budget = b
	e.mu.Unlock()
}

func (e *Engine) Calculate(model string, inputTokens, outputTokens int) (float64, string) {
	e.mu.RLock()
	p, ok := e.pricing[model]
	e.mu.RUnlock()

	if !ok {
		return 0, "USD"
	}

	inputCost := float64(inputTokens) / 1000.0 * p.InputPricePer1K
	outputCost := float64(outputTokens) / 1000.0 * p.OutputPricePer1K
	cost := inputCost + outputCost
	cost = math.Round(cost*100000) / 100000

	return cost, p.Currency
}

func (e *Engine) Record(model string, inputTokens, outputTokens int) Record {
	cost, currency := e.Calculate(model, inputTokens, outputTokens)
	r := Record{
		Model:     model,
		InputTok:  inputTokens,
		OutputTok: outputTokens,
		Cost:      cost,
		Currency:  currency,
		Timestamp: time.Now(),
	}
	e.mu.Lock()
	e.records = append(e.records, r)
	e.mu.Unlock()
	return r
}

func (e *Engine) TotalCost() float64 {
	e.mu.RLock()
	defer e.mu.RUnlock()
	var total float64
	for _, r := range e.records {
		total += r.Cost
	}
	return math.Round(total*100000) / 100000
}

func (e *Engine) BudgetRemaining() float64 {
	e.mu.RLock()
	defer e.mu.RUnlock()
	remaining := e.budget - e.TotalCost()
	if remaining < 0 {
		return 0
	}
	return remaining
}

func (e *Engine) IsOverBudget() bool {
	if e.budget <= 0 {
		return false
	}
	return e.TotalCost() >= e.budget
}

func (e *Engine) History() []Record {
	e.mu.RLock()
	defer e.mu.RUnlock()
	h := make([]Record, len(e.records))
	copy(h, e.records)
	return h
}

func (e *Engine) Format(cost float64, currency string) string {
	if cost == 0 {
		return "Free"
	}
	return fmt.Sprintf("%.5f %s", cost, currency)
}
