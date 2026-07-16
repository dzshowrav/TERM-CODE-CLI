package token

import "time"

type Usage struct {
	InputTokens      int     `json:"input_tokens"`
	OutputTokens     int     `json:"output_tokens"`
	TotalTokens      int     `json:"total_tokens"`
	ContextWindow    int     `json:"context_window"`
	RemainingContext int     `json:"remaining_context"`
	ContextPercent   float64 `json:"context_percent"`
	EstimatedCost    float64 `json:"estimated_cost"`
	RequestCount     int     `json:"request_count"`
}

type SessionTotals struct {
	TotalInputTokens  int           `json:"total_input"`
	TotalOutputTokens int           `json:"total_output"`
	TotalCost         float64       `json:"total_cost"`
	RequestCount      int           `json:"request_count"`
	SessionDuration   time.Duration `json:"session_duration"`
}

type ContextLevel int

const (
	ContextNormal  ContextLevel = 0
	ContextWarning ContextLevel = 1
	ContextFull    ContextLevel = 2
)

func CalculateCost(inputTokens, outputTokens int, pricingInput, pricingOutput float64) float64 {
	inputCost := (float64(inputTokens) / 1000.0) * pricingInput
	outputCost := (float64(outputTokens) / 1000.0) * pricingOutput
	return inputCost + outputCost
}
