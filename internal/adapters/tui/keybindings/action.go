package keybindings

type Action string

const (
	ActionNone      Action = ""
	ActionSubmit    Action = "submit"
	ActionReject    Action = "reject"
	ActionDelete    Action = "delete"
	ActionEscape    Action = "escape"
	ActionBackspace Action = "backspace"
	ActionTab       Action = "tab"
	ActionShiftTab  Action = "shift_tab"
	ActionUp        Action = "up"
	ActionDown      Action = "down"
	ActionLeft      Action = "left"
	ActionRight     Action = "right"
	ActionPageUp    Action = "page_up"
	ActionPageDown  Action = "page_down"
	ActionHome      Action = "home"
	ActionEnd       Action = "end"
	ActionCtrlC     Action = "ctrl_c"
	ActionCtrlD     Action = "ctrl_d"
	ActionCtrlU     Action = "ctrl_u"
	ActionCtrlF     Action = "ctrl_f"
	ActionCtrlB     Action = "ctrl_b"
	ActionCtrlP     Action = "ctrl_p"
	ActionEnter     Action = "enter"
	ActionSpace     Action = "space"

	ActionDiffNextFile Action = "diff_next_file"
	ActionDiffPrevFile Action = "diff_prev_file"
	ActionSearch       Action = "search"
	ActionZoomIn       Action = "zoom_in"
	ActionZoomOut      Action = "zoom_out"
	ActionCycleMermaid Action = "cycle_mermaid"

	ActionBestOfNCancel      Action = "best_of_n_cancel"
	ActionBestOfNSelectFocus Action = "best_of_n_select_focus"
	ActionBestOfNSwitchFocus Action = "best_of_n_switch_focus"
	ActionDownShift          Action = "down_shift"
)
