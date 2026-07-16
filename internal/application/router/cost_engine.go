package router

import "termcode/internal/domain/model"

type CostEngine struct{}

func NewCostEngine() *CostEngine {
	return &CostEngine{}
}

type CostEstimate struct {
	InputCost  float64 `json:"input_cost"`
	OutputCost float64 `json:"output_cost"`
	TotalCost  float64 `json:"total_cost"`
}

func (e *CostEngine) Estimate(mod *model.Model, inputTokens, outputTokens int) *CostEstimate {
	inputCost := (float64(inputTokens) / 1000.0) * mod.PricingInput
	outputCost := (float64(outputTokens) / 1000.0) * mod.PricingOut

	return &CostEstimate{
		InputCost:  inputCost,
		OutputCost: outputCost,
		TotalCost:  inputCost + outputCost,
	}
}

func (e *CostEngine) CheaperThan(mod *model.Model, other *model.Model, inputTokens, outputTokens int) bool {
	return e.Estimate(mod, inputTokens, outputTokens).TotalCost < e.Estimate(other, inputTokens, outputTokens).TotalCost
}
