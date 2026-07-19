package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/application/util"
)

type ToolStatus string

const (
	ToolQueued       ToolStatus = "queued"
	ToolInitializing ToolStatus = "initializing"
	ToolConnecting   ToolStatus = "connecting"
	ToolRunning      ToolStatus = "running"
	ToolCompleted    ToolStatus = "completed"
	ToolFailed       ToolStatus = "failed"
	ToolWaiting      ToolStatus = "waiting"
	ToolCancelled    ToolStatus = "cancelled"
)

var (
	toolNameStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Bold(true)
	toolArgStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	toolOutputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	toolErrorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	toolDurationSty = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	toolStatusSty   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	toolPathStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Underline(true)
	toolSepStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))

	spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
)

func toolStatusIcon(status ToolStatus, animFrame int) string {
	switch status {
	case ToolQueued:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Render("○")
	case ToolInitializing:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("141")).Render("●")
	case ToolConnecting:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("●")
	case ToolRunning:
		frame := spinnerFrames[animFrame%len(spinnerFrames)]
		return lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render(frame)
	case ToolCompleted:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Render("●")
	case ToolFailed:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("✗")
	case ToolWaiting:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Render("…")
	case ToolCancelled:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("⊘")
	default:
		return "○"
	}
}

type ToolCard struct {
	name       string
	status     ToolStatus
	args       string
	output     string
	errMsg     string
	durationMs int64
	width      int
	animFrame  int
	showAll    bool
	focused    bool
	collapsed  bool
}

func NewToolCard(name string) *ToolCard {
	return &ToolCard{
		name:   name,
		status: ToolQueued,
		width:  80,
	}
}

func (c *ToolCard) SetWidth(w int) {
	c.width = w
}

func (c *ToolCard) SetStatus(status ToolStatus) {
	c.status = status
}

func (c *ToolCard) SetArgs(args string) {
	c.args = args
}

func (c *ToolCard) SetOutput(output string) {
	c.output = output
}

func (c *ToolCard) SetError(err string) {
	c.errMsg = err
}

func (c *ToolCard) SetDuration(ms int64) {
	c.durationMs = ms
}

func (c *ToolCard) SetFocused(v bool) {
	c.focused = v
}

func (c *ToolCard) SetExpanded(v bool) {
	c.collapsed = !v
	if !v {
		c.showAll = false
	}
}

func (c *ToolCard) ToggleExpanded() {
	c.collapsed = !c.collapsed
	if c.collapsed {
		c.showAll = false
	}
}

func (c *ToolCard) ToggleShowAll() {
	c.showAll = !c.showAll
}

func (c *ToolCard) Expanded() bool {
	return !c.collapsed
}

func (c *ToolCard) SetAnimFrame(n int) {
	c.animFrame = n
}

func (c *ToolCard) AnimFrame() int {
	return c.animFrame
}

func (c *ToolCard) Name() string {
	return c.name
}

func (c *ToolCard) Status() ToolStatus {
	return c.status
}

func (c *ToolCard) Duration() int64 {
	return c.durationMs
}

func (c *ToolCard) IsRunning() bool {
	return c.status == ToolRunning || c.status == ToolInitializing || c.status == ToolConnecting || c.status == ToolQueued || c.status == ToolWaiting
}

func (c *ToolCard) durationStr() string {
	if c.durationMs < 1000 {
		return fmt.Sprintf("%dms", c.durationMs)
	}
	return fmt.Sprintf("%.1fs", float64(c.durationMs)/1000)
}

func (c *ToolCard) View() string {
	var lines []string

	icon := toolStatusIcon(c.status, c.animFrame)

	focusIndicator := ""
	if c.focused {
		focusIndicator = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("▸") + " "
	}

	name := c.name
	parts := strings.Split(name, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	name = strings.Join(parts, "")

	header := fmt.Sprintf("%s%s %s", focusIndicator, icon, toolNameStyle.Render(name))

	if c.args != "" {
		displayArgs := formatToolArgs(c.name, c.args)
		argsTrimmed := util.Truncate(displayArgs, c.width-lipgloss.Width(header)-10)
		header += fmt.Sprintf("(%s)", toolArgStyle.Render(argsTrimmed))
	}

	if c.durationMs > 0 {
		header += fmt.Sprintf("  %s", toolDurationSty.Render(c.durationStr()))
	}

	lines = append(lines, header)

	output := c.output
	if output == "" {
		if c.errMsg != "" {
			output = c.errMsg
		} else if c.collapsed {
			return strings.Join(lines, "\n")
		} else {
			return strings.Join(lines, "\n")
		}
	}

	pretty := prettyPrintJSON(output)
	if pretty != output {
		output = pretty
	}

	formatted := FormatToolOutput(output)
	outputLines := strings.Split(formatted, "\n")

	if c.collapsed {
		summary := c.generateSummary(outputLines)
		if summary != "" {
			truncated := util.Truncate(summary, c.width-6)
			lines = append(lines, "  "+lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Bold(true).Render("⎿")+"  "+toolOutputStyle.Render(truncated))
		}
		return strings.Join(lines, "\n")
	}

	total := len(outputLines)
	maxLines := 15
	if total > maxLines && !c.showAll {
		outputLines = outputLines[:maxLines]
	}

	indent := toolSepStyle.Render("  ")
	for _, line := range outputLines {
		rendered := toolOutputStyle.Render(indent + line)
		lines = append(lines, rendered)
	}

	if c.errMsg != "" && c.output == "" {
		for _, line := range strings.Split(c.errMsg, "\n") {
			lines = append(lines, toolErrorStyle.Render("  "+line))
		}
	}

	if total > maxLines {
		if c.showAll {
			lines = append(lines, toolArgStyle.Render("  \u25bc show less"))
		} else {
			lines = append(lines, toolArgStyle.Render(fmt.Sprintf("  \u25bc %d more lines", total-maxLines)))
		}
	}

	return strings.Join(lines, "\n")
}

func (c *ToolCard) generateSummary(lines []string) string {
	if c.errMsg != "" {
		return "Failed: " + c.firstMeaningfulLine(strings.Split(c.errMsg, "\n"))
	}
	if c.output == "" {
		return ""
	}

	name := strings.ToLower(c.name)
	if strings.Contains(name, "read") {
		return fmt.Sprintf("Read %d lines", len(lines))
	} else if strings.Contains(name, "list") && strings.Contains(name, "dir") {
		var files, dirs int
		var parsed []map[string]any
		if err := json.Unmarshal([]byte(c.output), &parsed); err == nil {
			for _, item := range parsed {
				if isDir, ok := item["is_dir"].(bool); ok && isDir {
					dirs++
				} else if typ, ok := item["type"].(string); ok && typ == "dir" {
					dirs++
				} else {
					files++
				}
			}
			return fmt.Sprintf("%d files, %d directories", files, dirs)
		}
		return fmt.Sprintf("%d items", len(lines))
	} else if strings.Contains(name, "search") || strings.Contains(name, "glob") || strings.Contains(name, "grep") {
		var parsed []map[string]any
		if err := json.Unmarshal([]byte(c.output), &parsed); err == nil {
			return fmt.Sprintf("Found %d results", len(parsed))
		}
		return fmt.Sprintf("Found %d results", len(lines))
	}

	return c.firstMeaningfulLine(lines)
}

func (c *ToolCard) firstMeaningfulLine(lines []string) string {
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			return trimmed
		}
	}
	return ""
}

func (c *ToolCard) Lines() []string {
	return strings.Split(c.View(), "\n")
}

func (c *ToolCard) Output() string {
	return c.output
}

func (c *ToolCard) OutputLines() int {
	if c.output == "" {
		return 0
	}
	return len(strings.Split(c.output, "\n"))
}

func FormatToolOutput(s string) string {
	lines := strings.Split(s, "\n")
	var buf bytes.Buffer
	inCode := false
	var codeLang string
	var codeBuf strings.Builder

	flushCode := func() {
		if codeBuf.Len() == 0 {
			return
		}
		highlighted := formatCodeBlock(codeBuf.String(), codeLang)
		buf.WriteString(highlighted)
		if !strings.HasSuffix(highlighted, "\n") {
			buf.WriteString("\n")
		}
		codeBuf.Reset()
		codeLang = ""
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "```") {
			if inCode {
				flushCode()
				buf.WriteString("\n")
			} else {
				flushCode()
				inCode = true
				codeLang = strings.TrimPrefix(trimmed, "```")
			}
			inCode = !inCode
			continue
		}

		if inCode {
			codeBuf.WriteString(line)
			codeBuf.WriteString("\n")
			continue
		}

		if IsDiffLine(line) {
			buf.WriteString(ColorizeDiff(line))
			buf.WriteString("\n")
			continue
		}

		formatted := formatMarkdownLine(line)
		buf.WriteString(formatted)
		buf.WriteString("\n")
	}

	flushCode()
	return buf.String()
}

func formatCodeBlock(code, lang string) string {
	if lang == "" {
		lang = "text"
	}
	var highlighted bytes.Buffer
	err := quick.Highlight(&highlighted, code, lang, "terminal256", "monokai")
	if err != nil || highlighted.Len() == 0 {
		return code
	}
	return highlighted.String()
}

func ColorizeDiff(line string) string {
	if len(line) == 0 {
		return line
	}
	switch line[0] {
	case '+':
		return lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Render(line)
	case '-':
		return lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(line)
	}
	if strings.HasPrefix(line, "@@") {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render(line)
	}
	if strings.HasPrefix(line, "--- ") || strings.HasPrefix(line, "+++ ") {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Render(line)
	}
	return line
}

func IsDiffLine(line string) bool {
	if len(line) == 0 {
		return false
	}
	if line[0] == '+' || line[0] == '-' {
		return true
	}
	if strings.HasPrefix(line, "@@") {
		return true
	}
	if strings.HasPrefix(line, "--- ") || strings.HasPrefix(line, "+++ ") {
		return true
	}
	return false
}

func formatMarkdownLine(line string) string {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return line
	}

	if strings.HasPrefix(trimmed, "#") {
		level := 0
		for _, ch := range trimmed {
			if ch == '#' {
				level++
			} else {
				break
			}
		}
		if level >= 1 && level <= 6 && len(trimmed) > level && trimmed[level] == ' ' {
			content := strings.TrimSpace(trimmed[level:])
			color := "39"
			switch level {
			case 1:
				color = "228"
			case 2:
				color = "221"
			case 3:
				color = "83"
			default:
				color = "39"
			}
			return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(color)).Render(content)
		}
	}

	if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
		bullet := lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("•")
		return "  " + bullet + " " + strings.TrimPrefix(strings.TrimPrefix(trimmed, "- "), "* ")
	}

	return line
}

func formatToolArgs(toolName, args string) string {
	if args == "" {
		return args
	}
	var parsed map[string]any
	if err := json.Unmarshal([]byte(args), &parsed); err != nil {
		return args
	}

	primaryKeys := []string{"path", "absolutepath", "query", "command", "file", "dir", "url", "targetfile", "searchpath"}

	// Fast path case-sensitive
	for _, k := range primaryKeys {
		if v, ok := parsed[k].(string); ok && v != "" {
			return v
		}
	}

	// Slow path case-insensitive
	for k, v := range parsed {
		kLower := strings.ToLower(k)
		for _, pk := range primaryKeys {
			if kLower == pk {
				if s, ok := v.(string); ok && s != "" {
					return s
				}
			}
		}
	}

	parts := make([]string, 0, len(parsed))
	for _, v := range parsed {
		if s, ok := v.(string); ok {
			parts = append(parts, s)
		}
	}
	if len(parts) == 0 {
		return args
	}
	return strings.Join(parts, ", ")
}

func prettyPrintJSON(s string) string {
	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return s
	}
	if trimmed[0] != '{' && trimmed[0] != '[' {
		return s
	}
	var parsed any
	if err := json.Unmarshal([]byte(trimmed), &parsed); err != nil {
		return s
	}
	formatted, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		return s
	}
	return string(formatted)
}
