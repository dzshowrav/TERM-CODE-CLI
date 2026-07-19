package keybindings

import tea "charm.land/bubbletea/v2"

type Keymap map[string]Action

var defaultKeymaps = map[Scope]Keymap{
	ScopeGlobal: {
		"ctrl+c": ActionCtrlC,
		"ctrl+p": ActionCtrlP,
		"tab":    ActionTab,
	},
	ScopeChat: {
		"enter":     ActionSubmit,
		"up":        ActionUp,
		"down":      ActionDown,
		"pgup":      ActionPageUp,
		"pgdown":    ActionPageDown,
		"home":      ActionHome,
		"end":       ActionEnd,
		"ctrl+u":    ActionCtrlU,
		"ctrl+d":    ActionCtrlD,
		"ctrl+f":    ActionCtrlF,
		"ctrl+b":    ActionCtrlB,
		"esc":       ActionEscape,
		"backspace": ActionBackspace,
		"shift+tab": ActionShiftTab,
	},
	ScopeDiff: {
		"right":  ActionDiffNextFile,
		"left":   ActionDiffPrevFile,
		"up":     ActionUp,
		"down":   ActionDown,
		"tab":    ActionTab,
		"esc":    ActionEscape,
		"ctrl+s": ActionSearch,
		"enter":  ActionZoomIn,
	},
	ScopeSearch: {
		"enter":     ActionSubmit,
		"esc":       ActionEscape,
		"up":        ActionUp,
		"down":      ActionDown,
		"tab":       ActionTab,
		"backspace": ActionBackspace,
	},
	ScopeDialog: {
		"enter":     ActionSubmit,
		"esc":       ActionEscape,
		"tab":       ActionTab,
		"shift+tab": ActionShiftTab,
		"up":        ActionUp,
		"down":      ActionDown,
	},
	ScopeList: {
		"enter":  ActionSubmit,
		"esc":    ActionEscape,
		"up":     ActionUp,
		"down":   ActionDown,
		"pgup":   ActionPageUp,
		"pgdown": ActionPageDown,
		"home":   ActionHome,
		"end":    ActionEnd,
	},
	ScopeSettings: {
		"right": ActionRight,
		"left":  ActionLeft,
		"enter": ActionSubmit,
		"esc":   ActionEscape,
		"tab":   ActionTab,
		"up":    ActionUp,
		"down":  ActionDown,
	},
}

type Manager struct {
	keymaps    map[Scope]Keymap
	models     map[ModelID]ModelPanelConfig
	scopeStack []Scope
}

func NewManager() *Manager {
	km := make(map[Scope]Keymap)
	for scope, kmap := range defaultKeymaps {
		k := make(Keymap)
		for key, action := range kmap {
			k[key] = action
		}
		km[scope] = k
	}
	return &Manager{
		keymaps: km,
		models:  make(map[ModelID]ModelPanelConfig),
	}
}

func (m *Manager) PushScope(scope Scope) {
	m.scopeStack = append(m.scopeStack, scope)
}

func (m *Manager) PopScope() {
	if len(m.scopeStack) > 0 {
		m.scopeStack = m.scopeStack[:len(m.scopeStack)-1]
	}
}

func (m *Manager) CurrentScope() Scope {
	if len(m.scopeStack) > 0 {
		return m.scopeStack[len(m.scopeStack)-1]
	}
	return ScopeGlobal
}

func (m *Manager) RegisterModel(modelID ModelID, config ModelPanelConfig) {
	m.models[modelID] = config
}

func (m *Manager) Resolve(msg tea.KeyMsg) Action {
	key := keyName(msg)
	if km, ok := m.keymaps[m.CurrentScope()]; ok {
		if action, ok := km[key]; ok {
			return action
		}
	}
	if km, ok := m.keymaps[ScopeGlobal]; ok {
		if action, ok := km[key]; ok {
			return action
		}
	}
	return ActionNone
}

func (m *Manager) Bind(scope Scope, key string, action Action) {
	if _, ok := m.keymaps[scope]; !ok {
		m.keymaps[scope] = make(Keymap)
	}
	m.keymaps[scope][key] = action
}

func (m *Manager) Unbind(scope Scope, key string) {
	if km, ok := m.keymaps[scope]; ok {
		delete(km, key)
	}
}

func keyName(msg tea.KeyMsg) string {
	s := msg.String()
	switch s {
	case "ctrl+c":
		return "ctrl+c"
	case "ctrl+p":
		return "ctrl+p"
	case "ctrl+u":
		return "ctrl+u"
	case "ctrl+d":
		return "ctrl+d"
	case "ctrl+f":
		return "ctrl+f"
	case "ctrl+b":
		return "ctrl+b"
	case "ctrl+s":
		return "ctrl+s"
	case "shift+tab":
		return "shift+tab"
	case "pgup":
		return "pgup"
	case "pgdown":
		return "pgdown"
	case " ":
		return "space"
	default:
		return s
	}
}
