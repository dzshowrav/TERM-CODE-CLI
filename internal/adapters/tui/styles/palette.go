package styles

import (
	"sync"

	"github.com/charmbracelet/lipgloss"
)

type Palette struct {
	mu     sync.RWMutex
	isDark bool
	colors map[string]lipgloss.Color
}

type PaletteName string

const (
	PaletteDefault PaletteName = "default"
	PaletteDark    PaletteName = "dark"
	PaletteLight   PaletteName = "light"
	PaletteOcean   PaletteName = "ocean"
	PaletteForest  PaletteName = "forest"
	PaletteMonokai PaletteName = "monokai"
)

var builtinPalettes = map[PaletteName]map[string]string{
	PaletteDefault: {
		"bg":         "236",
		"fg":         "255",
		"accent":     "39",
		"success":    "42",
		"warning":    "214",
		"error":      "196",
		"muted":      "240",
		"subtle":     "245",
		"border":     "240",
		"selection":  "33",
		"highlight":  "226",
		"link":       "45",
		"code":       "43",
		"heading":    "39",
		"blockquote": "242",
	},
	PaletteDark: {
		"bg":         "233",
		"fg":         "255",
		"accent":     "75",
		"success":    "77",
		"warning":    "220",
		"error":      "161",
		"muted":      "238",
		"subtle":     "243",
		"border":     "236",
		"selection":  "24",
		"highlight":  "228",
		"link":       "81",
		"code":       "79",
		"heading":    "75",
		"blockquote": "241",
	},
	PaletteLight: {
		"bg":         "255",
		"fg":         "232",
		"accent":     "27",
		"success":    "28",
		"warning":    "214",
		"error":      "160",
		"muted":      "248",
		"subtle":     "243",
		"border":     "250",
		"selection":  "153",
		"highlight":  "226",
		"link":       "33",
		"code":       "23",
		"heading":    "27",
		"blockquote": "247",
	},
	PaletteOcean: {
		"bg":         "17",
		"fg":         "255",
		"accent":     "39",
		"success":    "42",
		"warning":    "214",
		"error":      "196",
		"muted":      "240",
		"subtle":     "245",
		"border":     "239",
		"selection":  "33",
		"highlight":  "226",
		"link":       "45",
		"code":       "43",
		"heading":    "39",
		"blockquote": "242",
	},
	PaletteForest: {
		"bg":         "22",
		"fg":         "255",
		"accent":     "71",
		"success":    "77",
		"warning":    "214",
		"error":      "160",
		"muted":      "237",
		"subtle":     "243",
		"border":     "235",
		"selection":  "65",
		"highlight":  "228",
		"link":       "79",
		"code":       "43",
		"heading":    "71",
		"blockquote": "242",
	},
	PaletteMonokai: {
		"bg":         "234",
		"fg":         "255",
		"accent":     "141",
		"success":    "83",
		"warning":    "221",
		"error":      "197",
		"muted":      "236",
		"subtle":     "244",
		"border":     "235",
		"selection":  "62",
		"highlight":  "228",
		"link":       "117",
		"code":       "83",
		"heading":    "141",
		"blockquote": "244",
	},
}

var globalPalette *Palette

func init() {
	globalPalette = NewPalette(true)
	globalPalette.Load(PaletteDefault)
}

func NewPalette(dark bool) *Palette {
	return &Palette{
		isDark: dark,
		colors: make(map[string]lipgloss.Color),
	}
}

func (p *Palette) Load(name PaletteName) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if palette, ok := builtinPalettes[name]; ok {
		p.colors = make(map[string]lipgloss.Color, len(palette))
		for k, v := range palette {
			p.colors[k] = lipgloss.Color(v)
		}
	}
}

func (p *Palette) Set(key string, color lipgloss.Color) {
	p.mu.Lock()
	p.colors[key] = color
	p.mu.Unlock()
}

func (p *Palette) Color(name string) lipgloss.Color {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if c, ok := p.colors[name]; ok {
		return c
	}
	return lipgloss.Color("255")
}

func (p *Palette) IsDark() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.isDark
}

func (p *Palette) SetDark(dark bool) {
	p.mu.Lock()
	p.isDark = dark
	p.mu.Unlock()
}

func (p *Palette) Style(name string) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(p.Color(name))
}

func (p *Palette) Bg(name string) lipgloss.Color {
	return p.Color(name)
}

func (p *Palette) Fg(name string) lipgloss.Color {
	return p.Color(name)
}

func (p *Palette) Accent() lipgloss.Color  { return p.Color("accent") }
func (p *Palette) Muted() lipgloss.Color   { return p.Color("muted") }
func (p *Palette) Subtle() lipgloss.Color  { return p.Color("subtle") }
func (p *Palette) Success() lipgloss.Color { return p.Color("success") }
func (p *Palette) Warning() lipgloss.Color { return p.Color("warning") }
func (p *Palette) Error() lipgloss.Color   { return p.Color("error") }
func (p *Palette) Border() lipgloss.Color  { return p.Color("border") }

func SetTheme(name PaletteName) {
	globalPalette.Load(name)
}

func GetPalette() *Palette {
	return globalPalette
}

func Color(name string) lipgloss.Color {
	return globalPalette.Color(name)
}

func PaletteStyle(name string) lipgloss.Style {
	return globalPalette.Style(name)
}
