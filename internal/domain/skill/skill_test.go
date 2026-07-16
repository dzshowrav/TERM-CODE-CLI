package skill_test

import (
	"testing"

	"termcode/internal/domain/skill"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		skillName   string
		description string
		category    skill.Category
	}{
		{
			name:        "creates coding skill",
			skillName:   "go-best-practices",
			description: "Go development conventions",
			category:    skill.CategoryCoding,
		},
		{
			name:      "creates general skill",
			skillName: "communication",
			category:  skill.CategoryGeneral,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := skill.New(tt.skillName, tt.description, tt.category)
			if s.ID == "" {
				t.Error("expected non-empty ID")
			}
			if s.Name != tt.skillName {
				t.Errorf("expected Name=%q", tt.skillName)
			}
			if s.Description != tt.description {
				t.Errorf("expected Description=%q", tt.description)
			}
			if s.Category != tt.category {
				t.Errorf("expected Category=%v", tt.category)
			}
		})
	}
}

func TestNew_Defaults(t *testing.T) {
	s := skill.New("test", "", skill.CategoryGeneral)
	if s.Version != "1.0" {
		t.Errorf("expected Version=1.0, got %q", s.Version)
	}
	if !s.Enabled {
		t.Error("expected Enabled=true")
	}
	if s.IsBuiltin {
		t.Error("expected IsBuiltin=false")
	}
}

func TestNew_AllCategories(t *testing.T) {
	categories := []skill.Category{
		skill.CategoryGeneral,
		skill.CategoryCoding,
		skill.CategoryReview,
		skill.CategoryTesting,
		skill.CategorySecurity,
		skill.CategoryCustom,
	}
	for _, cat := range categories {
		t.Run(string(cat), func(t *testing.T) {
			s := skill.New("test", "", cat)
			if s.Category != cat {
				t.Errorf("expected Category=%q, got %q", cat, s.Category)
			}
		})
	}
}
