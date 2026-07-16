# Theming

## Overview

The theming system controls the visual appearance of the agent — colors, typography, spacing, and visual style. Themes are fully customizable and can be switched at runtime.

## Theme Definition

A theme is a collection of color mappings:

### Color Palette
- **Primary** — main accent color (used for highlights, active elements)
- **Secondary** — secondary accent (used for less prominent elements)
- **Background** — main background color
- **Surface** — secondary background (cards, dialogs)
- **Text** — primary text color
- **Text muted** — dimmed/secondary text
- **Border** — border and separator colors
- **Success** — green/positive indicators
- **Warning** — yellow/caution indicators
- **Error** — red/error indicators
- **Info** — blue/informational indicators

### Semantic Colors (Code Syntax)
- **Keyword** — language keywords (function, if, return)
- **String** — string literals
- **Number** — numeric literals
- **Comment** — comments
- **Type** — type names
- **Function** — function names
- **Variable** — variable names
- **Operator** — operators
- **Punctuation** — brackets, semicolons, commas

### UI Element Colors
- **Status bar** — background, text, indicators
- **Input area** — background, text, cursor
- **Selection** — selected text background
- **Scrollbar** — scrollbar track and thumb
- **Dialog** — dialog background, border, overlay
- **Link** — link text color

### Visual Properties
- **Border style** — rounded, straight, double, none
- **Spacing** — compact, normal, relaxed
- **Animation speed** — none, slow, normal, fast
- **Cursor style** — block, underline, bar, none

## Built-in Themes

The agent ships with a curated set of themes:
- **Tokyo Night** (primary/default) — dark blue-based theme with vibrant accents
- **Catppuccin** (mocha, latte, frappe, macchiato)
- **Dracula** — dark purple-based
- **Nord** — arctic blue-gray
- **Solarized** (dark and light)
- **Gruvbox** (dark and light)
- **One Dark** — Atom editor style
- **Monokai** — classic Monokai
- **GitHub** (dark and light)
- **Terminal default** — uses the terminal emulator's own colors

Total: 27+ themes covering dark, light, and high-contrast variants.

## Terminal Color Integration

The agent can read the terminal emulator's color scheme:

### Termux-Style Integration
- Reads the terminal's color configuration file
- Parses color values (0-15 ANSI colors, background, foreground)
- Maps terminal colors to theme colors
- Creates a "terminal-native" theme that matches the user's terminal exactly
- Updates in real-time if the terminal configuration changes

### Generic Terminal Detection
- Reads environment variables for terminal color info
- Detects light vs dark background
- Adapts theme contrast automatically
- Falls back gracefully if terminal colors are unavailable

## Dynamic Theme Loading

Themes are loaded:
1. **Built-in** — compiled into the agent
2. **File-based** — loaded from a themes directory
3. **URL-based** — loaded from a remote URL
4. **Inline** — defined in configuration

Themes can be:
- **Switched** at runtime without restart
- **Previewed** before applying
- **Customized** and saved as a new theme
- **Shared** by exporting the theme definition

## Theme Creation

Users can create themes:
- By modifying an existing theme in configuration
- By providing a full theme definition file
- By using the theme creation wizard (/themes create)
- By importing from popular editor themes

Theme files are plain-text, human-readable format.

## Accessibility

- All built-in themes meet WCAG contrast ratios
- A high-contrast theme is available for low-vision users
- Themes never rely solely on color to convey information
- Focus indicators are always visible regardless of theme
- The "terminal default" theme uses the user's familiar terminal colors

## Key Design Decisions

- Themes are purely cosmetic — they never affect functionality
- Tokyo Night is the default because it's optimized for terminal use (good contrast, reduced eye strain)
- Terminal color integration lets the agent blend into the user's existing setup
- All 27+ themes are equally maintained — none are second-class
- Theme switching is instant — no reload or restart needed
- Theme files are portable — shareable between users and installations
- Semantic color names ensure themes work with all UI elements
