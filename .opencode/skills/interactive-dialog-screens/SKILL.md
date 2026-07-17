---
name: interactive-dialog-screens
description: Creating interactive DialogScreen components in TermCode with search, scrolling, keyboard navigation, and styled rendering
license: MIT
compatibility: opencode
metadata:
  audience: developers
  framework: bubbletea
---

## DialogScreen Interface

Every reusable dialog in TermCode implements the `DialogScreen` interface (`screens/dialog.go`):

```go
type DialogScreen interface {
    SetSize(w, h int)
    Update(msg tea.Msg) (DialogScreen, tea.Cmd)
    View() string
    Done() bool
    Result() string
}
```

## File Location

Put new dialog screens in `internal/adapters/tui/screens/` (NOT `dialogs/`). The `dialogs/` package uses an older `tea.Model` pattern; new code uses the `DialogScreen` interface.

## Rendering with DialogBox

Use `styles.DialogBox(s.width, body)` for rounded border rendering (`╭─╮││╰─╯`). Inner content width = `s.width - 2`. Each body line is padded/truncated to inner width.

```go
func (s *MyScreen) View() string {
    innerW := s.width - 2
    lines := []string{
        title,
        styles.DialogSep(innerW), // separator line
        content,
    }
    body := strings.Join(lines, "\n")
    return styles.DialogBox(s.width, body)
}
```

- `styles.DialogSep(innerW)` renders a colored separator (`─` repeated, color 236)
- `styles.HintStyle` for hints/placeholders (color 240, italic)
- `styles.Subtitle` for section headers (color 250)
- `styles.Active` for cursor/active elements (color 39, bold)
- `styles.ValueStyle` for primary content (color 255)

## Search + Scroll Pattern

For interactive lists with search and scrolling:

```go
type MyScreen struct {
    width    int
    height   int
    items    []domainmodel.Model  // full list
    filtered []domainmodel.Model  // search-filtered
    search   string
    cursor   int
    scroll   int
    done     bool
    result   string
    onSelect func(id string) string
}
```

### Search Filtering

```go
func (s *MyScreen) applyFilter() {
    s.filtered = nil
    if s.search == "" {
        s.filtered = append(s.filtered, s.items...)
    } else {
        lower := strings.ToLower(s.search)
        for _, m := range s.items {
            if strings.Contains(strings.ToLower(m.DisplayName), lower) {
                s.filtered = append(s.filtered, m)
            }
        }
    }
    // clamp cursor after filtering
    if s.cursor >= len(s.filtered) {
        s.cursor = max(0, len(s.filtered)-1)
    }
}
```

### Update Handler

```go
case tea.KeyMsg:
    switch msg.String() {
    case "esc":
        s.done = true
    case "enter":
        // select current item
        m := s.filtered[s.cursor]
        s.result = s.onSelect(...)
        s.done = true
    case "up", "k":
        if s.cursor > 0 { s.cursor--; s.ensureVisible() }
    case "down", "j":
        if s.cursor < len(s.filtered)-1 { s.cursor++; s.ensureVisible() }
    case "backspace":
        if len(s.search) > 0 {
            s.search = s.search[:len(s.search)-1]
            s.cursor = 0; s.scroll = 0
            s.applyFilter()
        }
    default:
        r := []rune(msg.String())
        if len(r) == 1 && !unicode.IsControl(r[0]) {
            s.search += string(r)
            s.cursor = 0; s.scroll = 0
            s.applyFilter()
        }
    }
case tea.PasteMsg:
    s.search += msg.String() // NOT string(msg) — PasteMsg is a struct with .String()
    s.cursor = 0; s.scroll = 0
    s.applyFilter()
```

### Height Calculation for View

```go
func (s *MyScreen) entryBounds() (maxEntries int) {
    bodyFixed := 7   // title + sep + search + sep + header + blank + bottomBar
    scrollOverhead := 0
    if s.scroll > 0 { scrollOverhead++ }
    remaining := len(s.filtered) - s.scroll - maxEntries
    if remaining > 0 { scrollOverhead++ }
    avail := s.height - bodyFixed - 5 - scrollOverhead // 5 = borders(2) + AppModel overhead(3)
    if avail < 1 { return 0 }
    return avail
}
```

### Scroll Indicators

```go
if s.scroll > 0 {
    bodyLines = append(bodyLines, styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll)))
}
if end < len(s.filtered) {
    bodyLines = append(bodyLines, styles.HintStyle.Render(fmt.Sprintf("↓ %d more", len(s.filtered)-end)))
}
```

## Registration in commands.go

```go
r.register("mycommand", func(ctx context.Context, args []string) string {
    return r.cmdMyCommand(ctx, args, modelSvc)
})
```

In the handler, convert domain models (which come as `[]*Model`) to value types before passing to the screen:

```go
func (r *commandRegistry) cmdMyCommand(ctx context.Context, args []string, modelSvc *model.Service) string {
    items, err := modelSvc.List(ctx)
    if err != nil { return fmt.Sprintf("Error: %v", err) }
    models := make([]domainmodel.Model, len(items))
    for i, m := range items {
        models[i] = *m  // dereference pointer to value
    }
    s := screens.NewMyScreen(onSelect)
    s.SetModels(models)
    r.app.ShowDialog(s)
    return "__dialog__"
}
```

## Use Vertical Centering

The AppModel automatically centers dialogs vertically in `View()` when `m.activeDialog != nil`. No manual centering needed in the screen's `View()`.

## What Didn't Work

- **Using `dialogs/` package (older Pattern B)**: These dialogs return `tea.Model` directly and use `lipgloss.Place()` for centering. They're inconsistent with the newer `DialogScreen` pattern used by `screens/`. Always use `screens/` for new dialogs.
- **`string(msg)` for PasteMsg**: `tea.PasteMsg` is a struct with a `Content` field, not a `string`. Use `msg.String()` instead. A direct `string(msg)` cast causes a compile error.
- **Named return values with `:=` shadowing**: In a function with named returns like `func() (x int) { x = 1; return }`, using `x := 1` inside shadows the return variable. Use `x = 1` (plain assignment) to write to the named return.
- **Appending non-Go content after `Write()`**: When using the `Write` tool to create a .go file, ensure no trailing markdown or comments are included after the final `}`. Extra lines cause syntax errors. Use `head -n <line> file.go > clean.go` via bash to strip trailing content.
