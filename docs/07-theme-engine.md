# 07-theme-engine.md

# Theme Engine Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Theme Engine?

A **Theme Engine** is the subsystem responsible for controlling the visual appearance of a Terminal User Interface (TUI).

It manages colors, typography styles, borders, icons, spacing, highlights, syntax colors, and component styling without changing the application logic.

The rendering system asks the Theme Engine **how** something should look, while the Theme Engine never decides **what** should be displayed.

---

# Why Theme Engine?

Without Theme Engine

```
Renderer

↓

Hardcoded Colors

↓

Impossible To Customize
```

Example

```
print(RED)

print(GREEN)

print(BLUE)
```

Problems

- Hard to maintain
- Duplicate colors
- No theme switching
- Poor accessibility
- Difficult branding

---

With Theme Engine

```
Renderer

↓

Theme Engine

↓

Current Theme

↓

ANSI Colors

↓

Terminal
```

---

# Goals

A production Theme Engine should provide

- Theme switching
- Color palettes
- Semantic colors
- Component styling
- Syntax highlighting
- Accessibility
- High contrast mode
- User customization
- Live reload
- Theme inheritance

---

# High-Level Architecture

```
                Settings
                    │
                    ▼
             Theme Manager
                    │
      ┌─────────────┼──────────────┐
      ▼             ▼              ▼
 Theme Loader   Theme Registry   Theme Cache
      │             │              │
      └──────┬──────┴──────┬───────┘
             ▼             ▼
      Theme Resolver   ANSI Generator
             │
             ▼
        Render Pipeline
             │
             ▼
          Terminal
```

---

# Folder Structure

```
themes/

    dracula.json

    github.json

    catppuccin.json

    nord.json

    everforest.json

src/

theme/

    ThemeManager.ts

    ThemeLoader.ts

    ThemeRegistry.ts

    ThemeResolver.ts

    ThemeCache.ts

    ThemeValidator.ts

    ThemeEvents.ts

    ThemePalette.ts

    ThemeTokens.ts

    ThemeGenerator.ts

    ANSIConverter.ts

    Accessibility.ts
```

---

# Core Components

## Theme Manager

Central controller.

Responsibilities

- Load theme
- Switch theme
- Notify renderer
- Manage cache

---

## Theme Registry

Stores

```
Installed Themes

Names

Versions

Authors

Metadata
```

---

## Theme Loader

Reads themes from disk.

Responsibilities

- Parse JSON
- Validate schema
- Register palettes
- Cache themes

---

## Theme Resolver

Converts semantic tokens into actual colors.

Example

```
Primary

↓

#61AFEF

↓

ANSI 39
```

---

## ANSI Converter

Converts theme colors into terminal-compatible ANSI escape codes.

Supports

```
16 Color

256 Color

True Color (24-bit)
```

---

## Theme Cache

Stores active theme in memory.

Benefits

- Faster rendering
- Less file access
- Instant switching

---

## Accessibility Manager

Provides

- High contrast
- Reduced color confusion
- Readability checks

---

# Theme Lifecycle

```
Application Start

↓

Load Settings

↓

Read Theme

↓

Validate

↓

Register

↓

Activate

↓

Render
```

---

# Theme Loading Flow

```
Theme File

↓

Loader

↓

Validator

↓

Registry

↓

Resolver

↓

Renderer
```

---

# Theme Structure

```
theme/

    manifest.json

    colors.json

    syntax.json

    components.json

    icons.json

    README.md
```

---

# Theme Manifest

Contains

```
Name

Version

Author

Description

License

Compatibility
```

---

# Color Palette

Example

```
Background

Foreground

Primary

Secondary

Accent

Border

Selection

Highlight

Muted

Disabled
```

---

# Semantic Color Tokens

Never use raw colors in components.

Use tokens

```
text.primary

text.secondary

button.primary

button.secondary

panel.background

status.success

status.warning

status.error
```

Renderer only understands tokens.

---

# UI Components

Theme may define styles for

```
Header

Footer

Sidebar

Panel

Button

Input

Status Bar

Dialog

Menu

Progress Bar

Spinner

Table

Tooltip
```

---

# Syntax Highlighting

Theme may define

```
Keyword

String

Comment

Function

Variable

Number

Operator

Type

Constant
```

Useful for code blocks.

---

# Icon Theme

Icons may define

```
Folder

File

Git

Error

Warning

Success

Info

Loading
```

Icons remain separate from colors.

---

# Typography

Terminal typography options

```
Bold

Italic

Underline

Dim

Reverse

Hidden

Blink
```

Availability depends on terminal support.

---

# Theme Inheritance

Example

```
Dark Theme

↓

Custom Theme

↓

Override Only Needed Values
```

Reduces duplication.

---

# Live Theme Switching

```
User Changes Theme

↓

Theme Manager

↓

Reload Theme

↓

Emit Event

↓

Renderer Updates

↓

Done
```

No restart required.

---

# Event Bus Integration

Common events

```
theme:load

theme:change

theme:reload

theme:error
```

Renderer subscribes to these events.

---

# Renderer Integration

```
Render Request

↓

Theme Resolver

↓

ANSI Converter

↓

Terminal Output
```

Renderer never stores colors.

---

# Plugin Integration

Plugins may

- Register themes
- Extend palettes
- Add icons
- Add syntax definitions

---

# Skills Integration

Skills may request

```
Markdown Theme

Code Theme

Diff Theme
```

The Theme Engine resolves them automatically.

---

# Configuration

Users may customize

```
Theme

Accent Color

Border Style

Icon Set

Syntax Theme

High Contrast

Transparency (if supported)
```

---

# Theme Validation

Always check

- Required tokens
- Missing colors
- Invalid values
- Compatibility
- Duplicate identifiers

---

# Performance Optimizations

Use

- Theme cache
- Token lookup table
- Lazy loading
- ANSI caching
- Immutable palettes

Avoid

- Parsing JSON every frame
- Hardcoded colors
- Rebuilding palettes repeatedly

---

# Security

Always

- Validate theme files
- Restrict file paths
- Ignore unsupported fields
- Use safe defaults

Never

- Execute code from themes
- Trust malformed JSON
- Allow arbitrary scripts

---

# Best Practices

Always

- Use semantic tokens
- Separate colors from components
- Support dark themes
- Provide accessibility options
- Cache resolved values
- Keep themes immutable

Never

- Hardcode ANSI values in components
- Mix business logic with styling
- Duplicate palettes
- Assume terminal capabilities

---

# Common Mistakes

Bad

```
Button

↓

Foreground = Blue

Background = Black
```

Color is fixed forever.

---

Good

```
Button

↓

button.primary

↓

Theme Resolver

↓

Actual Color
```

Theme controls appearance.

---

# Testing Checklist

- Theme loading
- Validation
- Live switching
- ANSI conversion
- Accessibility mode
- Syntax highlighting
- Component styling
- Theme inheritance
- Missing tokens
- Error recovery

---

# Example Theme Collection

Examples

- Dracula
- Catppuccin
- Nord
- GitHub
- Everforest
- Tokyo Night
- One Dark
- Gruvbox
- Solarized
- Flexoki
- Carbonfox
- Nightfox

---

# Advantages

- Easy customization
- Better maintainability
- Accessibility support
- Consistent UI
- Brand flexibility
- Cleaner rendering
- Community theme ecosystem
- Runtime switching

---

# Disadvantages

- Requires token management
- Theme compatibility must be maintained
- More validation logic
- Terminal capability differences

---

# Used In

- OpenCode
- Antigravity CLI
- VS Code
- Neovim
- Helix
- Warp Terminal
- WezTerm
- Ghostty
- Lazygit
- GitUI

---

# Summary

The **Theme Engine** is responsible for separating visual presentation from application logic in a terminal-based AI Coding Agent.

A production-grade Theme Engine should include:

- Theme Manager
- Loader
- Registry
- Resolver
- ANSI Converter
- Theme Cache
- Semantic Color Tokens
- Syntax Highlighting
- Accessibility Support
- Event Bus Integration

By using semantic tokens, runtime theme switching, and ANSI color abstraction, the application becomes highly customizable, visually consistent, and easy to maintain while supporting a growing ecosystem of community-created themes similar to OpenCode, Antigravity CLI, and modern terminal applications.