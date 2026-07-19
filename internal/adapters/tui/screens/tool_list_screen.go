package screens

import (
	"fmt"
	"strings"
	"unicode"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
	"termcode/internal/domain/tool"
)

var (
	toolCatStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Bold(true)
	toolNameSt    = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
	toolDescSt    = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	toolDetailSt  = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	toolTagDang   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	toolTagSafe   = lipgloss.NewStyle().Foreground(lipgloss.Color("83"))
	toolTagPlugin = lipgloss.NewStyle().Foreground(lipgloss.Color("141"))
	toolKeyword   = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
)

type ToolListItem struct {
	Tool tool.Tool
}

type ToolListScreen struct {
	width    int
	height   int
	tools    []tool.Tool
	filtered []tool.Tool
	search   string
	cursor   int
	scroll   int
	expanded int // index into filtered that is expanded (-1 = none)
	done     bool
	result   string
}

func NewToolListScreen(tools []tool.Tool) *ToolListScreen {
	return &ToolListScreen{
		width:    80,
		height:   24,
		tools:    tools,
		expanded: -1,
	}
}

func (s *ToolListScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ToolListScreen) Done() bool {
	return s.done
}

func (s *ToolListScreen) Result() string {
	return s.result
}

func (s *ToolListScreen) applyFilter() {
	s.filtered = make([]tool.Tool, 0, len(s.tools))
	if s.search == "" {
		s.filtered = append(s.filtered, s.tools...)
	} else {
		lower := strings.ToLower(s.search)
		for _, t := range s.tools {
			if strings.Contains(strings.ToLower(t.Name), lower) ||
				strings.Contains(strings.ToLower(t.Description), lower) ||
				strings.Contains(strings.ToLower(string(t.Category)), lower) {
				s.filtered = append(s.filtered, t)
				continue
			}
			for _, a := range t.Aliases {
				if strings.Contains(strings.ToLower(a), lower) {
					s.filtered = append(s.filtered, t)
					break
				}
			}
		}
	}
	if s.cursor >= len(s.filtered) {
		s.cursor = max(0, len(s.filtered)-1)
	}
	if s.cursor < 0 {
		s.cursor = 0
	}
	if s.expanded >= len(s.filtered) {
		s.expanded = -1
	}
}

func (s *ToolListScreen) ensureVisible() {
	if s.cursor < s.scroll {
		s.scroll = s.cursor
		return
	}

	innerW := s.width - 2
	for s.scroll < s.cursor {
		availHeight := s.height - 9
		if availHeight < 1 {
			availHeight = 1
		}
		if s.scroll > 0 {
			availHeight--
		}

		usedHeight := 0
		fits := false
		for i := s.scroll; i <= s.cursor; i++ {
			t := s.filtered[i]
			isCursor := i == s.cursor
			isExpanded := i == s.expanded
			linesCount := strings.Count(s.renderToolLine(t, innerW, isCursor, isExpanded), "\n") + 1

			needed := linesCount
			if i < len(s.filtered)-1 {
				needed++
			}

			if usedHeight+needed <= availHeight {
				usedHeight += linesCount
				if i == s.cursor {
					fits = true
					break
				}
			} else {
				if i == s.cursor && i == len(s.filtered)-1 && usedHeight+linesCount <= availHeight {
					fits = true
				}
				break
			}
		}

		if fits {
			break
		}
		s.scroll++
	}
}

func (s *ToolListScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
			s.result = ""
			return s, nil
		case "enter":
			if len(s.filtered) > 0 && s.cursor >= 0 && s.cursor < len(s.filtered) {
				if s.expanded == s.cursor {
					s.expanded = -1
				} else {
					s.expanded = s.cursor
				}
			}
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
				s.ensureVisible()
			}
		case "down", "j":
			if s.cursor < len(s.filtered)-1 {
				s.cursor++
				s.ensureVisible()
			}
		case "backspace":
			if len(s.search) > 0 {
				s.search = s.search[:len(s.search)-1]
				s.cursor = 0
				s.scroll = 0
				s.expanded = -1
				s.applyFilter()
			}
		default:
			r := []rune(msg.String())
			if len(r) == 1 && !unicode.IsControl(r[0]) {
				s.search += string(r)
				s.cursor = 0
				s.scroll = 0
				s.expanded = -1
				s.applyFilter()
			}
		}
	case tea.PasteMsg:
		s.search += msg.String()
		s.cursor = 0
		s.scroll = 0
		s.expanded = -1
		s.applyFilter()
	}
	return s, nil
}

func (s *ToolListScreen) View() string {
	innerW := s.width - 2
	if s.filtered == nil {
		s.applyFilter()
	}

	title := fmt.Sprintf("%-*s%s", innerW-4, "Tools", "esc")

	searchText := s.search
	if searchText == "" {
		searchText = styles.HintStyle.Render("search tools")
	}
	searchLine := fmt.Sprintf("Search: %s", searchText)

	availHeight := s.height - 9
	if availHeight < 1 {
		availHeight = 1
	}

	showTopScroll := s.scroll > 0
	if showTopScroll {
		availHeight--
	}

	var bodyLines []string
	bodyLines = append(bodyLines, title)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, searchLine)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))

	// Show count
	countText := fmt.Sprintf("%d tools", len(s.filtered))
	if len(s.filtered) != len(s.tools) {
		countText = fmt.Sprintf("%d of %d tools", len(s.filtered), len(s.tools))
	}
	bodyLines = append(bodyLines, styles.HintStyle.Render("  "+countText))

	var listLines []string
	showBottomScroll := false
	remainingCount := 0

	for i := s.scroll; i < len(s.filtered); i++ {
		t := s.filtered[i]
		isCursor := i == s.cursor
		isExpanded := i == s.expanded
		itemText := s.renderToolLine(t, innerW, isCursor, isExpanded)
		itemLines := strings.Split(itemText, "\n")
		linesCount := len(itemLines)

		needed := linesCount
		if i < len(s.filtered)-1 {
			needed++
		}

		if len(listLines)+needed <= availHeight {
			listLines = append(listLines, itemLines...)
		} else {
			if i == len(s.filtered)-1 && len(listLines)+linesCount <= availHeight {
				listLines = append(listLines, itemLines...)
			} else {
				showBottomScroll = true
				remainingCount = len(s.filtered) - i
				break
			}
		}
	}

	if showTopScroll {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll))))
	}

	bodyLines = append(bodyLines, listLines...)

	if showBottomScroll {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↓ %d more", remainingCount))))
	}

	bodyLines = append(bodyLines, "")
	hintW := innerW
	hintText := styles.HintStyle.Render("esc: Close  Enter: Expand/collapse  Type: Search")
	hintPad := hintW - lipgloss.Width(hintText)
	if hintPad < 0 {
		hintPad = 0
	}
	bodyLines = append(bodyLines, fmt.Sprintf("%s%s", strings.Repeat(" ", hintPad), hintText))

	body := strings.Join(bodyLines, "\n")
	return styles.DialogBox(s.width, body)
}

func (s *ToolListScreen) renderToolLine(t tool.Tool, width int, isCursor, isExpanded bool) string {
	var lines []string

	cursor := " "
	if isCursor {
		cursor = styles.Active.Render(">")
	}

	label := t.Name
	if t.DisplayName != "" && t.DisplayName != t.Name {
		label = t.DisplayName
	}
	name := toolNameSt.Render(label)
	cat := styles.HintStyle.Render(fmt.Sprintf("[%s]", t.Category))

	// Tags
	var tags []string
	if t.Dangerous {
		tags = append(tags, toolTagDang.Render("!danger"))
	}
	if t.Author != "" && t.Author != "built-in" {
		tags = append(tags, toolTagPlugin.Render(t.Author))
	}
	if len(t.Aliases) > 0 {
		tags = append(tags, styles.HintStyle.Render(strings.Join(t.Aliases, ",")))
	}
	tagStr := ""
	if len(tags) > 0 {
		tagStr = " " + strings.Join(tags, " ")
	}

	desc := t.Description
	if len(desc) > 60 {
		desc = desc[:57] + "..."
	}

	mainText := fmt.Sprintf("%s %s %s  %s%s", cursor, name, cat, toolDescSt.Render(desc), tagStr)
	mainW := lipgloss.Width(mainText)
	if mainW > width {
		mainText = lipgloss.NewStyle().Width(width).Render(mainText)
	}
	lines = append(lines, mainText)

	if isExpanded {
		lines = append(lines, s.renderToolDetails(t, width))
	}

	return strings.Join(lines, "\n")
}

func (s *ToolListScreen) renderToolDetails(t tool.Tool, width int) string {
	indent := "    "
	var details []string

	details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Version:"), toolDetailSt.Render(t.Version)))

	source := t.Source
	if source == "" {
		source = "built-in"
	}
	details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Source:"), toolDetailSt.Render(source)))

	if t.PermissionLevel != "" {
		details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Permission:"), toolDetailSt.Render(string(t.PermissionLevel))))
	}

	if len(t.AllowedContexts) > 0 {
		ctxs := make([]string, len(t.AllowedContexts))
		for i, c := range t.AllowedContexts {
			ctxs[i] = string(c)
		}
		details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Contexts:"), toolDetailSt.Render(strings.Join(ctxs, ", "))))
	}

	if t.RateLimit > 0 {
		details = append(details, fmt.Sprintf("%s%s %d/min", indent, toolKeyword.Render("Rate limit:"), t.RateLimit))
	}

	if t.DefaultTimeout > 0 {
		details = append(details, fmt.Sprintf("%s%s %dms", indent, toolKeyword.Render("Timeout:"), t.DefaultTimeout))
	}

	if len(t.Aliases) > 0 {
		details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Aliases:"), toolDetailSt.Render(strings.Join(t.Aliases, ", "))))
	}

	if len(t.Capabilities) > 0 {
		caps := make([]string, len(t.Capabilities))
		for i, c := range t.Capabilities {
			caps[i] = string(c)
		}
		details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Capabilities:"), toolDetailSt.Render(strings.Join(caps, ", "))))
	}

	if t.DefaultTimeout > 0 {
		details = append(details, fmt.Sprintf("%s%s %dms", indent, toolKeyword.Render("Timeout:"), t.DefaultTimeout))
	}

	if t.Dangerous {
		details = append(details, fmt.Sprintf("%s%s", indent, toolTagDang.Render("Dangerous - requires permission confirmation")))
	}

	hooks := t.Hooks
	var hookNames []string
	if hooks.OnBefore != nil {
		hookNames = append(hookNames, "before")
	}
	if hooks.OnStart != nil {
		hookNames = append(hookNames, "start")
	}
	if hooks.OnDelta != nil {
		hookNames = append(hookNames, "delta")
	}
	if hooks.OnEnd != nil {
		hookNames = append(hookNames, "end")
	}
	if hooks.OnError != nil {
		hookNames = append(hookNames, "error")
	}
	if hooks.OnAbort != nil {
		hookNames = append(hookNames, "abort")
	}
	if len(hookNames) > 0 {
		details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Hooks:"), toolDetailSt.Render(strings.Join(hookNames, ", "))))
	}

	// Show input schema summary
	if schema, ok := t.InputSchema.(map[string]any); ok {
		if props, ok := schema["properties"].(map[string]any); ok {
			var paramNames []string
			for k := range props {
				paramNames = append(paramNames, k)
			}
			if required, ok := schema["required"].([]any); ok {
				reqSet := make(map[string]bool, len(required))
				for _, r := range required {
					reqSet[fmt.Sprintf("%v", r)] = true
				}
				var reqNames []string
				var optNames []string
				for _, n := range paramNames {
					if reqSet[n] {
						reqNames = append(reqNames, n)
					} else {
						optNames = append(optNames, n)
					}
				}
				if len(reqNames) > 0 {
					details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Required:"), toolDetailSt.Render(strings.Join(reqNames, ", "))))
				}
				if len(optNames) > 0 {
					details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Optional:"), toolDetailSt.Render(strings.Join(optNames, ", "))))
				}
			} else {
				details = append(details, fmt.Sprintf("%s%s %s", indent, toolKeyword.Render("Params:"), toolDetailSt.Render(strings.Join(paramNames, ", "))))
			}
		}
	}

	return strings.Join(details, "\n")
}
