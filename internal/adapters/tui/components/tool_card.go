package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/charmbracelet/lipgloss"
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
	toolBorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240"))
	toolBorderRunning = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("39"))
	toolBorderFailed = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("196"))
	toolNameStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Bold(true)
	toolArgStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Italic(true)
	toolOutputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	toolErrorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	toolDurationSty = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	toolLabelStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	toolActionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Underline(true)

	spinnerFrames = []string{"◐", "◓", "◑", "◒"}
)

func toolStatusIcon(status ToolStatus, animFrame int) string {
	switch status {
	case ToolQueued:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Render("○")
	case ToolInitializing:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("141")).Render("⚙")
	case ToolConnecting:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("↗")
	case ToolRunning:
		frame := spinnerFrames[animFrame%len(spinnerFrames)]
		return lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render(frame)
	case ToolCompleted:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Render("✓")
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

func toolStatusColor(status ToolStatus) string {
	switch status {
	case ToolCompleted:
		return "83"
	case ToolRunning:
		return "39"
	case ToolWaiting, ToolQueued:
		return "221"
	case ToolInitializing:
		return "141"
	case ToolConnecting:
		return "39"
	case ToolFailed:
		return "196"
	case ToolCancelled:
		return "245"
	default:
		return "245"
	}
}

type ToolCard struct {
	name        string
	status      ToolStatus
	args        string
	output      string
	errMsg      string
	durationMs  int64
	progressPct int
	expanded    bool
	width       int
	label       string
	focused     bool
	animFrame   int
	showAll     bool
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

func (c *ToolCard) SetProgress(pct int) {
	c.progressPct = pct
}

func (c *ToolCard) SetExpanded(v bool) {
	c.expanded = v
	if !v {
		c.showAll = false
	}
}

func (c *ToolCard) SetFocused(v bool) {
	c.focused = v
}

func (c *ToolCard) SetAnimFrame(n int) {
	c.animFrame = n
}

func (c *ToolCard) AnimFrame() int {
	return c.animFrame
}

func (c *ToolCard) SetShowAll(v bool) {
	c.showAll = v
}

func (c *ToolCard) ToggleShowAll() {
	c.showAll = !c.showAll
}

func (c *ToolCard) ToggleExpanded() {
	c.expanded = !c.expanded
	if !c.expanded {
		c.showAll = false
	}
}

func (c *ToolCard) Expanded() bool {
	return c.expanded
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
	innerW := c.width - 4
	if innerW < 10 {
		innerW = 10
	}

	icon := toolStatusIcon(c.status, c.animFrame)
	color := toolStatusColor(c.status)
	statusStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

	expandIcon := "▶"
	if c.expanded {
		expandIcon = "▼"
	}

	focusIndicator := ""
	if c.focused {
		focusIndicator = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("▸") + " "
	}
	header := fmt.Sprintf("%s%s %s %s  %s",
		focusIndicator,
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(expandIcon),
		icon,
		toolNameStyle.Render(c.name),
		statusStyle.Render(string(c.status)),
	)

	if c.durationMs > 0 {
		header += fmt.Sprintf("  %s", toolDurationSty.Render(c.durationStr()))
	}

	var lines []string
	lines = append(lines, header)

	if c.expanded {
		if c.args != "" && len(c.args) < 200 {
			lines = append(lines, "")
			lines = append(lines, toolLabelStyle.Render("Arguments")+" "+toolArgStyle.Render(truncateStr(c.args, innerW-12)))
		}

		if c.status == ToolRunning && c.progressPct > 0 {
			lines = append(lines, "")
			lines = append(lines, c.renderProgress(innerW))
		}

		if c.output != "" {
			lines = append(lines, "")
			lines = append(lines, toolLabelStyle.Render("Output"))
			totalLines := c.OutputLines()
			rawOutput := c.output
			pretty := prettyPrintJSON(rawOutput)
			if pretty != rawOutput {
				rawOutput = pretty
				totalLines = len(strings.Split(rawOutput, "\n"))
			}
			formatted := formatToolOutput(rawOutput)
			var displayOutput string
			if c.showAll {
				displayOutput = formatted
			} else {
				displayOutput = truncateOutput(formatted, innerW, 15)
			}
			for _, line := range strings.Split(displayOutput, "\n") {
				if strings.HasPrefix(line, "\033") {
					lines = append(lines, "  "+line)
				} else {
					lines = append(lines, toolOutputStyle.Render("  "+line))
				}
			}
			if totalLines > 15 {
				if c.showAll {
					lines = append(lines, toolActionStyle.Render("  ▼ Show Less"))
				} else {
					lines = append(lines, toolActionStyle.Render("  ▶ Show All"))
				}
			}
		}

		if c.errMsg != "" {
			lines = append(lines, "")
			lines = append(lines, toolLabelStyle.Render("Error"))
			stackLines, logLines, suggestionLines := categorizeError(c.errMsg)
			if len(stackLines) > 0 || len(logLines) > 0 || len(suggestionLines) > 0 {
				if len(logLines) > 0 {
					lines = append(lines, toolLabelStyle.Render("  Logs"))
					for _, l := range logLines {
						lines = append(lines, toolErrorStyle.Render("    "+l))
					}
				}
				if len(stackLines) > 0 {
					lines = append(lines, toolLabelStyle.Render("  Stack"))
					for _, l := range stackLines {
						lines = append(lines, toolErrorStyle.Render("    "+l))
					}
				}
				if len(suggestionLines) > 0 {
					lines = append(lines, toolLabelStyle.Render("  Suggestion"))
					for _, l := range suggestionLines {
						lines = append(lines, lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Render("    "+l))
					}
				}
			} else {
				lines = append(lines, toolErrorStyle.Render("  "+c.errMsg))
			}
		}

		if c.status == ToolFailed {
			lines = append(lines, "")
			lines = append(lines, toolActionStyle.Render("Retry"))
		}
	}

	rendered := strings.Join(lines, "\n")

	var border lipgloss.Style
	switch c.status {
	case ToolRunning:
		border = toolBorderRunning
	case ToolFailed:
		border = toolBorderFailed
	default:
		border = toolBorderStyle
	}

	if c.focused {
		border = border.BorderForeground(lipgloss.Color("39"))
	}

	result := border.Width(innerW).Render(rendered)
	return result
}

func (c *ToolCard) renderProgress(w int) string {
	barWidth := w - 20
	if barWidth < 5 {
		barWidth = 5
	}
	filled := (c.progressPct * barWidth) / 100
	if filled > barWidth {
		filled = barWidth
	}
	empty := barWidth - filled
	bar := lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render(strings.Repeat("■", filled)) +
		lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(strings.Repeat("■", empty))
	return fmt.Sprintf("  %s  %d%%", bar, c.progressPct)
}

func (c *ToolCard) Lines() []string {
	return strings.Split(c.View(), "\n")
}

func (c *ToolCard) OutputLines() int {
	if c.output == "" {
		return 0
	}
	return len(strings.Split(c.output, "\n"))
}

func truncateOutput(s string, maxW, maxLines int) string {
	lines := strings.Split(s, "\n")
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}
	for i, line := range lines {
		if len(line) > maxW {
			lines[i] = line[:maxW-3] + "..."
		}
	}
	return strings.Join(lines, "\n")
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

func colorizeDiff(line string) string {
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

func isDiffLine(line string) bool {
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

func isTableRow(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return false
	}
	if trimmed[0] != '|' {
		return false
	}
	pipes := strings.Count(trimmed, "|")
	return pipes >= 2
}

func formatTable(content string) string {
	lines := strings.Split(content, "\n")
	if len(lines) < 2 {
		return content
	}

	var rows [][]string
	var separators []int
	for i, line := range lines {
		if !isTableRow(line) {
			continue
		}
		cells := splitTableRow(line)
		if i == 1 && strings.Count(strings.TrimSpace(lines[1]), "-") > 0 {
			separators = make([]int, len(cells))
			for j, c := range cells {
				separators[j] = len(c)
			}
			continue
		}
		rows = append(rows, cells)
	}

	if len(rows) == 0 {
		return content
	}

	numCols := 0
	for _, row := range rows {
		if len(row) > numCols {
			numCols = len(row)
		}
	}
	if numCols == 0 {
		return content
	}

	if separators == nil || len(separators) != numCols {
		separators = make([]int, numCols)
		for _, row := range rows {
			for j, cell := range row {
				clean := strings.TrimSpace(cell)
				if len(clean) > separators[j] {
					separators[j] = len(clean)
				}
			}
		}
	}

	var buf bytes.Buffer
	for _, row := range rows {
		for j, cell := range row {
			if j >= numCols {
				break
			}
			clean := strings.TrimSpace(cell)
			width := separators[j]
			if j == len(row)-1 {
				buf.WriteString(fmt.Sprintf("  %s", clean))
			} else {
				buf.WriteString(fmt.Sprintf("  %-*s", width, clean))
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func splitTableRow(line string) []string {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) < 2 {
		return nil
	}
	inner := trimmed
	if inner[0] == '|' {
		inner = inner[1:]
	}
	if len(inner) > 0 && inner[len(inner)-1] == '|' {
		inner = inner[:len(inner)-1]
	}
	var cells []string
	current := strings.Builder{}
	for _, ch := range inner {
		if ch == '|' {
			cells = append(cells, current.String())
			current.Reset()
		} else {
			current.WriteRune(ch)
		}
	}
	if current.Len() > 0 {
		cells = append(cells, current.String())
	}
	return cells
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

	if strings.HasPrefix(trimmed, "---") || strings.HasPrefix(trimmed, "***") || strings.HasPrefix(trimmed, "___") {
		sepStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
		return sepStyle.Render(strings.Repeat("─", 40))
	}

	if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
		bullet := lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("•")
		return "  " + bullet + " " + strings.TrimPrefix(strings.TrimPrefix(trimmed, "- "), "* ")
	}

	if strings.HasPrefix(trimmed, "1.") || strings.HasPrefix(trimmed, "1)") {
		numStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
		return "  " + numStyle.Render("1.") + " " + strings.TrimSpace(trimmed[2:])
	}

	return line
}

func formatToolOutput(s string) string {
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

	diffLines := 0
	tableLines := 0

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

		if isDiffLine(line) {
			diffLines++
			if diffLines >= 3 {
				buf.WriteString(colorizeDiff(line))
				buf.WriteString("\n")
				continue
			}
		} else {
			diffLines = 0
		}

		if isTableRow(line) {
			tableLines++
		} else {
			if tableLines > 0 {
				tableLines = 0
			}
		}

		formatted := formatMarkdownLine(line)
		buf.WriteString(formatted)
		buf.WriteString("\n")
	}

	flushCode()
	result := buf.String()

	if tableLines > 0 {
		result = formatTable(result)
	}

	return result
}

func categorizeError(err string) (stack, logs, suggestions []string) {
	lines := strings.Split(err, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		low := strings.ToLower(trimmed)
		if strings.Contains(low, "suggestion:") || strings.Contains(low, "hint:") || strings.Contains(low, "try ") {
			suggestions = append(suggestions, trimmed)
		} else if matched, _ := fmt.Sscanf(trimmed, "%s.go:%d", new(string), new(int)); matched == 2 || strings.Contains(trimmed, ".go:") {
			stack = append(stack, trimmed)
		} else {
			logs = append(logs, trimmed)
		}
	}
	if len(suggestions) == 0 && len(stack) == 0 && len(logs) == 0 {
		logs = append(logs, err)
	}
	return
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

func truncateStr(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen-3]) + "..."
}
