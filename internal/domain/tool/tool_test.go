package tool_test

import (
	"context"
	"testing"

	"termcode/internal/domain/tool"
)

func TestTool_New(t *testing.T) {
	tr := tool.New("bash", "Execute shell commands", map[string]any{"type": "object"}, tool.CatTerminal, tool.CapExecute)
	if tr.Name != "bash" {
		t.Errorf("expected Name=bash, got %q", tr.Name)
	}
	if tr.Category != tool.CatTerminal {
		t.Errorf("expected Category=terminal, got %q", tr.Category)
	}
	if len(tr.Capabilities) != 1 || tr.Capabilities[0] != tool.CapExecute {
		t.Errorf("expected Capabilities=[execute]")
	}
	if tr.Version != "1.0.0" {
		t.Errorf("expected Version=1.0.0, got %q", tr.Version)
	}
	if tr.DefaultTimeout != 30000 {
		t.Errorf("expected DefaultTimeout=30000, got %d", tr.DefaultTimeout)
	}
	if tr.Source != "built-in" {
		t.Errorf("expected Source=built-in, got %q", tr.Source)
	}
}

func TestTool_MatchName(t *testing.T) {
	tr := tool.New("read", "Read file", nil, tool.CatFileSystem)
	tr.Aliases = []string{"read_file", "cat"}
	if !tr.MatchName("read") {
		t.Error("expected MatchName(read)=true")
	}
	if !tr.MatchName("read_file") {
		t.Error("expected MatchName(read_file)=true")
	}
	if tr.MatchName("write") {
		t.Error("expected MatchName(write)=false")
	}
}

func TestResult_Defaults(t *testing.T) {
	r := tool.Result{
		Tool:   "bash",
		Input:  "ls",
		Output: "file1.txt",
		Status: tool.StatusSuccess,
	}
	if r.Tool != "bash" {
		t.Errorf("expected Tool=bash, got %q", r.Tool)
	}
	if r.Output != "file1.txt" {
		t.Errorf("expected Output=file1.txt, got %q", r.Output)
	}
	if r.Status != tool.StatusSuccess {
		t.Errorf("expected Status=success, got %q", r.Status)
	}
}

func TestResult_Truncate(t *testing.T) {
	r := tool.Result{Output: "hello world hello world"}
	r.Truncate(10)
	if r.Output != "hello worl... (truncated)" {
		t.Errorf("unexpected truncated output: %q", r.Output)
	}
	if !r.Truncated {
		t.Error("expected Truncated=true")
	}
	if r.RawSize != 23 {
		t.Errorf("expected RawSize=23, got %d", r.RawSize)
	}
}

func TestResult_Error(t *testing.T) {
	r := tool.Result{
		Tool:   "bash",
		Input:  "invalid",
		Error:  "command not found",
		Status: tool.StatusFailed,
	}
	if r.Error != "command not found" {
		t.Errorf("expected Error=command not found, got %q", r.Error)
	}
	if r.Status != tool.StatusFailed {
		t.Errorf("expected Status=failed, got %q", r.Status)
	}
}

func TestStatusValues(t *testing.T) {
	statuses := []tool.Status{
		tool.StatusPending,
		tool.StatusRunning,
		tool.StatusApproved,
		tool.StatusDenied,
		tool.StatusSuccess,
		tool.StatusFailed,
		tool.StatusTimeout,
		tool.StatusAborted,
	}
	for _, s := range statuses {
		t.Run(string(s), func(t *testing.T) {
			if s == "" {
				t.Error("expected non-empty status")
			}
		})
	}
}

func TestCategoryValues(t *testing.T) {
	cats := []tool.Category{
		tool.CatFileSystem,
		tool.CatDevelopment,
		tool.CatGit,
		tool.CatSearch,
		tool.CatNetwork,
		tool.CatCode,
		tool.CatTerminal,
		tool.CatUtility,
		tool.CatPlugin,
		tool.CatDatabase,
		tool.CatBrowser,
		tool.CatDangerous,
	}
	for _, c := range cats {
		t.Run(string(c), func(t *testing.T) {
			if c == "" {
				t.Error("expected non-empty category")
			}
		})
	}
}

func TestCapabilityValues(t *testing.T) {
	caps := []tool.Capability{
		tool.CapRead,
		tool.CapWrite,
		tool.CapExecute,
		tool.CapNetwork,
		tool.CapFiles,
		tool.CapSearch,
		tool.CapGit,
		tool.CapDatabase,
		tool.CapBrowser,
		tool.CapDanger,
	}
	for _, c := range caps {
		t.Run(string(c), func(t *testing.T) {
			if c == "" {
				t.Error("expected non-empty capability")
			}
		})
	}
}

func TestRegistry_RegisterAndLookup(t *testing.T) {
	reg := tool.NewRegistry()
	t1 := tool.New("read", "Read file", nil, tool.CatFileSystem)
	t1.Aliases = []string{"read_file"}
	t2 := tool.New("write", "Write file", nil, tool.CatFileSystem)

	if err := reg.Register(t1); err != nil {
		t.Fatalf("Register t1: %v", err)
	}
	if err := reg.Register(t2); err != nil {
		t.Fatalf("Register t2: %v", err)
	}

	if reg.Count() != 2 {
		t.Errorf("expected Count=2, got %d", reg.Count())
	}

	found, ok := reg.Lookup("read")
	if !ok || found.Name != "read" {
		t.Error("expected to find 'read'")
	}

	found, ok = reg.Lookup("read_file")
	if !ok || found.Name != "read" {
		t.Error("expected to find 'read' via alias 'read_file'")
	}

	_, ok = reg.Lookup("nonexistent")
	if ok {
		t.Error("expected not to find 'nonexistent'")
	}
}

func TestRegistry_Remove(t *testing.T) {
	reg := tool.NewRegistry()
	t1 := tool.New("read", "Read file", nil, tool.CatFileSystem)
	t1.Aliases = []string{"read_file"}
	if err := reg.Register(t1); err != nil {
		t.Fatalf("Register t1: %v", err)
	}
	if err := reg.Register(tool.New("write", "Write file", nil, tool.CatFileSystem)); err != nil {
		t.Fatalf("Register write: %v", err)
	}

	reg.Remove("read")
	if reg.Count() != 1 {
		t.Errorf("expected Count=1 after remove, got %d", reg.Count())
	}

	_, ok := reg.Lookup("read_file")
	if ok {
		t.Error("expected alias 'read_file' to be removed too")
	}
}

func TestRegistry_FilterByCategory(t *testing.T) {
	reg := tool.NewRegistry()
	reg.Register(tool.New("read", "", nil, tool.CatFileSystem))
	reg.Register(tool.New("bash", "", nil, tool.CatTerminal))
	reg.Register(tool.New("write", "", nil, tool.CatFileSystem))

	fs := reg.FilterByCategory(tool.CatFileSystem)
	if len(fs) != 2 {
		t.Errorf("expected 2 filesystem tools, got %d", len(fs))
	}
}

func TestValidateArgs_Required(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"path":    map[string]any{"type": "string"},
			"content": map[string]any{"type": "string"},
		},
		"required": []any{"path", "content"},
	}

	// Missing required
	r := tool.ValidateArgs(schema, map[string]any{"path": "/tmp/f"})
	if r.Valid {
		t.Error("expected validation to fail for missing 'content'")
	}
	if len(r.Errors) != 1 || r.Errors[0].Field != "content" {
		t.Errorf("expected error for 'content', got %v", r.Errors)
	}

	// Valid
	r = tool.ValidateArgs(schema, map[string]any{"path": "/tmp/f", "content": "hello"})
	if !r.Valid {
		t.Errorf("expected valid, got errors: %v", r.Errors)
	}
}

func TestValidateArgs_TypeCheck(t *testing.T) {
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"count": map[string]any{"type": "integer"},
			"name":  map[string]any{"type": "string"},
		},
		"required": []any{"count"},
	}

	r := tool.ValidateArgs(schema, map[string]any{"count": "not-a-number"})
	if r.Valid {
		t.Error("expected type validation to fail")
	}

	r = tool.ValidateArgs(schema, map[string]any{"count": 5, "name": "test"})
	if !r.Valid {
		t.Errorf("expected valid, got: %v", r.Errors)
	}
}

func TestRegistry_RegisterAll(t *testing.T) {
	reg := tool.NewRegistry()
	if err := reg.RegisterAll([]tool.Tool{
		tool.New("a", "", nil, tool.CatUtility),
		tool.New("b", "", nil, tool.CatUtility),
		tool.New("c", "", nil, tool.CatUtility),
	}); err != nil {
		t.Fatalf("RegisterAll: %v", err)
	}
	if reg.Count() != 3 {
		t.Errorf("expected Count=3, got %d", reg.Count())
	}
}

func TestRegistry_Conflict(t *testing.T) {
	reg := tool.NewRegistry()
	t1 := tool.New("read", "", nil, tool.CatFileSystem)
	t1.Source = "built-in"
	if err := reg.Register(t1); err != nil {
		t.Fatalf("Register t1: %v", err)
	}
	t2 := tool.New("read", "", nil, tool.CatFileSystem)
	t2.Source = "mcp"
	err := reg.Register(t2)
	if err == nil {
		t.Fatal("expected conflict error for same name from different source")
	}
}

func TestHooks_Executed(t *testing.T) {
	var calls []string
	t1 := tool.New("test", "", nil, tool.CatUtility)
	t1.Hooks = tool.ToolHooks{
		OnBefore: func(ctx context.Context, args map[string]any) error {
			calls = append(calls, "before")
			return nil
		},
		OnStart: func(ctx context.Context, args map[string]any) {
			calls = append(calls, "start")
		},
		OnEnd: func(ctx context.Context, result *tool.Result) {
			calls = append(calls, "end")
		},
		OnError: func(ctx context.Context, err error) {
			calls = append(calls, "error")
		},
	}

	_ = t1.Hooks.OnBefore(context.Background(), nil)
	t1.Hooks.OnStart(context.Background(), nil)
	t1.Hooks.OnEnd(context.Background(), &tool.Result{})

	expected := []string{"before", "start", "end"}
	for i, v := range expected {
		if i >= len(calls) || calls[i] != v {
			t.Errorf("expected call[%d]=%q, got %v", i, v, calls)
		}
	}
}
