package token_test

import (
	"testing"

	"termcode/internal/domain/token"
)

func TestCalculateCost(t *testing.T) {
	tests := []struct {
		name          string
		inputTokens   int
		outputTokens  int
		pricingInput  float64
		pricingOutput float64
		want          float64
	}{
		{name: "no tokens", inputTokens: 0, outputTokens: 0, pricingInput: 10, pricingOutput: 30, want: 0},
		{name: "only input", inputTokens: 1000, outputTokens: 0, pricingInput: 10, pricingOutput: 30, want: 10},
		{name: "only output", inputTokens: 0, outputTokens: 500, pricingInput: 10, pricingOutput: 30, want: 15},
		{name: "both tokens", inputTokens: 2000, outputTokens: 1000, pricingInput: 5, pricingOutput: 15, want: 25},
		{name: "partial tokens", inputTokens: 500, outputTokens: 200, pricingInput: 10, pricingOutput: 30, want: 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := token.CalculateCost(tt.inputTokens, tt.outputTokens, tt.pricingInput, tt.pricingOutput)
			if got != tt.want {
				t.Errorf("CalculateCost(%d,%d,%f,%f) = %f, want %f",
					tt.inputTokens, tt.outputTokens, tt.pricingInput, tt.pricingOutput, got, tt.want)
			}
		})
	}
}

func TestCalculateCost_ZeroPricing(t *testing.T) {
	got := token.CalculateCost(1000, 1000, 0, 0)
	if got != 0 {
		t.Errorf("expected 0 with zero pricing, got %f", got)
	}
}
