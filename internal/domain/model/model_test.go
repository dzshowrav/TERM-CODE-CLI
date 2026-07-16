package model_test

import (
	"testing"

	"termcode/internal/domain/model"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		providerID  string
		modelID     string
		displayName string
		category    model.Category
	}{
		{
			name:        "creates coding model",
			providerID:  "prov-1",
			modelID:     "gpt-4",
			displayName: "GPT-4",
			category:    model.CategoryCoding,
		},
		{
			name:        "creates general model",
			providerID:  "prov-2",
			modelID:     "gpt-3.5-turbo",
			displayName: "GPT-3.5",
			category:    model.CategoryGeneral,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := model.New(tt.providerID, tt.modelID, tt.displayName, tt.category)
			if m.ID == "" {
				t.Error("expected non-empty ID")
			}
			if m.ProviderID != tt.providerID {
				t.Errorf("expected ProviderID=%q", tt.providerID)
			}
			if m.ModelID != tt.modelID {
				t.Errorf("expected ModelID=%q", tt.modelID)
			}
			if m.DisplayName != tt.displayName {
				t.Errorf("expected DisplayName=%q", tt.displayName)
			}
			if m.Category != tt.category {
				t.Errorf("expected Category=%v", tt.category)
			}
		})
	}
}

func TestNew_Defaults(t *testing.T) {
	m := model.New("p1", "m1", "Test", model.CategoryGeneral)
	if !m.Capabilities.Streaming {
		t.Error("expected Streaming=true by default")
	}
	if !m.Capabilities.SystemPrompt {
		t.Error("expected SystemPrompt=true by default")
	}
	if m.MaxContext != 4096 {
		t.Errorf("expected MaxContext=4096, got %d", m.MaxContext)
	}
	if m.MaxOutput != 4096 {
		t.Errorf("expected MaxOutput=4096, got %d", m.MaxOutput)
	}
	if !m.Enabled {
		t.Error("expected Enabled=true")
	}
}

func TestNew_AllCategories(t *testing.T) {
	categories := []model.Category{
		model.CategoryGeneral,
		model.CategoryCoding,
		model.CategoryReasoning,
		model.CategoryVision,
		model.CategoryEmbedding,
		model.CategoryAudio,
		model.CategoryExperimental,
		model.CategoryCustom,
	}
	for _, cat := range categories {
		t.Run(string(cat), func(t *testing.T) {
			m := model.New("p1", "m1", "Test", cat)
			if m.Category != cat {
				t.Errorf("expected Category=%q, got %q", cat, m.Category)
			}
		})
	}
}
