package config_test

import (
	"testing"

	"termcode/internal/domain/config"
)

func TestDefaults(t *testing.T) {
	c := config.Defaults()
	if c == nil {
		t.Fatal("expected non-nil config")
	}
	if c.Temperature != 0.7 {
		t.Errorf("expected Temperature=0.7, got %f", c.Temperature)
	}
	if c.MaxTokens != 4096 {
		t.Errorf("expected MaxTokens=4096, got %d", c.MaxTokens)
	}
	if !c.Streaming {
		t.Error("expected Streaming=true")
	}
	if !c.AutoSave {
		t.Error("expected AutoSave=true")
	}
	if c.AutoSaveInterval != 30 {
		t.Errorf("expected AutoSaveInterval=30, got %d", c.AutoSaveInterval)
	}
	if !c.EncryptKeys {
		t.Error("expected EncryptKeys=true")
	}
	if c.DefaultPerm != "ask" {
		t.Errorf("expected DefaultPerm=ask, got %q", c.DefaultPerm)
	}
	if c.Timeout != 60 {
		t.Errorf("expected Timeout=60, got %d", c.Timeout)
	}
	if c.RetryCount != 3 {
		t.Errorf("expected RetryCount=3, got %d", c.RetryCount)
	}
}

func TestDefaults_FalseValues(t *testing.T) {
	c := config.Defaults()
	if c.CompactMode {
		t.Error("expected CompactMode=false")
	}
	if c.ShowTokenGauge != true {
		t.Error("expected ShowTokenGauge=true")
	}
}

func TestDefaults_EmptyStrings(t *testing.T) {
	c := config.Defaults()
	if c.DefaultProvider != "" {
		t.Errorf("expected empty DefaultProvider, got %q", c.DefaultProvider)
	}
	if c.DefaultModel != "" {
		t.Errorf("expected empty DefaultModel, got %q", c.DefaultModel)
	}
	if c.Theme != "" {
		t.Errorf("expected empty Theme, got %q", c.Theme)
	}
}
