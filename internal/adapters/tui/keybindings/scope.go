package keybindings

type Scope string

const (
	ScopeGlobal     Scope = "global"
	ScopeChat       Scope = "chat"
	ScopeDiff       Scope = "diff"
	ScopeSearch     Scope = "search"
	ScopeSettings   Scope = "settings"
	ScopeDialog     Scope = "dialog"
	ScopeList       Scope = "list"
	ScopeArtifact   Scope = "artifact"
	ScopeTaskDetail Scope = "task_detail"
	ScopeSubagent   Scope = "subagent"
)

type ModelID string

const (
	ModelRoot        ModelID = "root"
	ModelChat        ModelID = "chat"
	ModelHome        ModelID = "home"
	ModelDiff        ModelID = "diff"
	ModelSearch      ModelID = "search"
	ModelSettings    ModelID = "settings"
	ModelHelp        ModelID = "help"
	ModelArtifact    ModelID = "artifact"
	ModelAuth        ModelID = "auth"
	ModelMcpAuth     ModelID = "mcp_auth"
	ModelTasks       ModelID = "tasks"
	ModelSubagent    ModelID = "subagent"
	ModelToolConfirm ModelID = "tool_confirm"
	ModelPrompt      ModelID = "prompt"
	ModelFeedback    ModelID = "feedback"
)

type ModelPanelConfig struct {
	ModelID  ModelID
	Scope    Scope
	Priority int
}
