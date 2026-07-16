# Git TUI & CLI Tools — Complete A-to-Z Reference

> Essential Git tools from the Awesome CLI Apps ecosystem (81 tools).
> Source: toolleeo/awesome-cli-apps-in-a-csv

---

## 1. Lazygit — Terminal Git UI

**GitHub**: https://github.com/jesseduffield/lazygit  
**Language**: Go  
**Stars**: 55K+  
**Purpose**: Simple terminal UI for git commands

### Installation
```bash
# macOS
brew install lazygit
# Linux (Debian/Ubuntu)
apt install lazygit
# Go
go install github.com/jesseduffield/lazygit@latest
```

### Key Features
- **Full git workflow** in a TUI
- **Staging interface** — stage/unstage individual lines
- **Branch management** — create, checkout, merge
- **Rebase** — interactive rebase with easy UI
- **Stash management**
- **Diff viewer** — side-by-side or inline
- **Remote operations** — push, pull, fetch
- **Commit message** editor with emoji support
- **Worktree** management
- **Blame** viewer
- **Patch** management

### Keybindings (Main Panel)
| Key | Action |
|-----|--------|
| `1-5` | Switch panels |
| `tab` | Next panel |
| `space` | Stage/unstage |
| `c` | Commit |
| `C` | Commit using editor |
| `p` | Push |
| `P` | Pull |
| `s` | Stash |
| `S` | Stash staged |
| `d` | Diff |
| `g` | Git flow menu |
| `r` | Refresh |
| `+` | Next commit |
| `_` | Previous commit |
| `q` | Quit |
| `?` | Keybindings help |

### Configuration
```yaml
# ~/.config/lazygit/config.yml
gui:
  theme:
    lightTheme: false
    activeBorderColor:
      - green
      - bold
    inactiveBorderColor:
      - white
    selectedLineBgColor:
      - blue
  language: 'en'
git:
  paging:
    colorArg: always
    pager: delta
  commitPrefixes:
    main: { pattern: "^(\\w+\\/\\d+)", replace: "[$1] " }
```

### Key Differentiators
- **Most popular git TUI** (55K+ stars)
- **Zero-config** — works out of the box
- **Keyboard-first** — no mouse needed
- **Fast** — written in Go
- **Great for beginners** and experts
- **Customizable** themes and keybindings

---

## 2. GitUI — Rust Git TUI

**GitHub**: https://github.com/extrawurst/gitui  
**Language**: Rust  
**Stars**: 19K+  
**Purpose**: Blazing-fast terminal UI for git

### Installation
```bash
# macOS
brew install gitui
# cargo
cargo install gitui
# Linux (from releases)
curl -sS https://raw.githubusercontent.com/extrawurst/gitui/master/install.sh | sh
```

### Key Features
- **Fast startup** — written in Rust, nearly instant
- **Full git operations** — add, commit, push, pull
- **Staging area** management
- **Diff viewer** with syntax highlighting
- **Branch/tag** management
- **Blame** view
- **Stash** support
- **File tree** view
- **Commit graph** visualization
- **External editor** support
- **Custom hooks** integration

### Keybindings
| Key | Action |
|-----|--------|
| `s` | Stage file/line |
| `S` | Stage all |
| `c` | Commit |
| `C` | Commit (with message editor) |
| `p` | Push |
| `P` | Pull |
| `b` | Blame |
| `B` | Branch manager |
| `t` | File tree toggle |
| `h` | Help |
| `q` | Quit |
| `r` | Refresh |
| `?` | Keybindings |

### Configuration (TOML)
```toml
# ~/.config/gitui/key_bindings.toml
[popup_select]
event_type = "popup"
key = "AltKey-Enter"

[popup_close]
event_type = "popup"
key = "Escape"

# ~/.config/gitui/theme.toml
[default]
background = "#1e1e2e"
foreground = "#cdd6f4"
```

### Key Differentiators
- **Fastest git TUI** — instant startup (Rust)
- **Low memory** footprint
- **Small binary** (~5MB)
- **Syntax highlighting** in diffs
- **External diff tools** support (delta)
- **Active development** — frequent releases

---

## 3. tig — ncurses Git Browser

**GitHub**: https://github.com/jonas/tig  
**Language**: C  
**Stars**: 13K+  
**Purpose**: ncurses-based text-mode interface for git

### Installation
```bash
# macOS
brew install tig
# Linux
apt install tig
```

### Key Modes
| Mode | Key | Description |
|------|-----|-------------|
| **Main** | `m` | View commit log |
| **Diff** | `d` | View changes |
| **Log** | `l` | Commit history |
| **Tree** | `t` | File tree |
| **Blame** | `b` | File annotation |
| **Refs** | `r` | References view |
| **Stash** | `y` | Stash list |
| **Status** | `s` | Working tree status |
| **Stage** | `u` | Stage area |

### Keybindings
| Key | Action |
|-----|--------|
| `k/j` | Up/down |
| `n/N` | Next/previous search match |
| `u` | Update view |
| `m` | Merge |
| `I` | Cherry-pick |
| `U` | Revert |
| `!` | Prompt |
| `:` | Execute command |
| `z` | Suspend |
| `q` | Quit |

### Configuration
```bash
# ~/.tigrc
set show-changes = true
set line-graphics = utf-8
set number = true
set show-author = true
set commit-order = default

bind diff 1 previous-file
bind diff 2 next-file
```

### Key Differentiators
- **Lightest git TUI** — C, ncurses, zero deps
- **Veteran tool** — battle-tested, 15+ years old
- **Widely available** — in most package managers
- **Scriptable** — can output to pipelines
- **Vi-like keybindings**
- **Extremely configurable** via `.tigrc`

---

## 4. GitHub CLI (gh)

**Website**: https://cli.github.com  
**Repository**: https://github.com/cli/cli  
**Language**: Go  
**Purpose**: Official GitHub CLI

### Installation
```bash
# macOS
brew install gh
# Linux
apt install gh
# Windows
winget install GitHub.cli
```

### Core Commands
| Command | Description |
|---------|-------------|
| `gh auth` | Authentication |
| `gh repo` | Repository management |
| `gh issue` | Issue management |
| `gh pr` | Pull request management |
| `gh run` | Actions workflow runs |
| `gh release` | Release management |
| `gh gist` | Gist management |
| `gh codespace` | Codespaces |
| `gh api` | Raw API access |
| `gh search` | Search (repos, issues, code) |
| `gh browse` | Open in browser |
| `gh config` | Configuration |
| `gh completion` | Shell completion |
| `gh extension` | Extension management |

### Common Workflows
```bash
# PR workflow
gh pr create --fill
gh pr checkout 123
gh pr review 123 --approve
gh pr merge 123 --squash

# Issue workflow
gh issue list --assignee @me
gh issue create --label bug --title "fix: ..."
gh issue close 456

# Repo operations
gh repo clone owner/repo
gh repo create my-repo --public
gh repo fork owner/repo

# Search
gh search repos "topic:cli"
gh search issues --label "good first issue"

# Actions
gh run list
gh run watch
gh run rerun

# Releases
gh release create v1.0.0 --notes "Release notes"
gh release download v1.0.0
```

### Authentication
```bash
# Login
gh auth login
# Token-based
gh auth login --with-token < token.txt
# Check status
gh auth status
```

### Extensions
```bash
# Install extensions
gh extension install dlvhdr/gh-dash
gh extension install gennaro-tedesco/gh-s
gh extension install gennaro-tedesco/gh-f

# List extensions
gh extension list
```

---

## 5. gh-dash — GitHub Dashboard

**GitHub**: https://github.com/dlvhdr/gh-dash  
**Language**: Go  
**Purpose**: Beautiful CLI dashboard for GitHub

### Installation
```bash
gh extension install dlvhdr/gh-dash
```

### Usage
```bash
# Open dashboard
gh dash
# Custom sections
gh dash --section issues
# With sections config
gh dash --config ~/.config/gh-dash/config.yml
```

### Configuration
```yaml
# ~/.config/gh-dash/config.yml
sections:
  - name: "My PRs"
    filters: "is:open author:@me"
    limit: 20
  - name: "Review Requests"
    filters: "is:open review-requested:@me"
    limit: 20
  - name: "My Issues"
    filters: "is:open issue author:@me"
    limit: 10

theme:
  background: "#1e1e2e"
  foreground: "#cdd6f4"
```

### Key Features
- **Customizable sections** — PRs, issues, repos
- **Multiple views** — list, grid, table
- **Keybindings** — vim-like navigation
- **Open in browser** — with `o`
- **Dynamic filtering**
- **Theme support**

---

## 6. git-cliff — Changelog Generator

**GitHub**: https://github.com/orhun/git-cliff  
**Language**: Rust  
**Purpose**: Highly customizable changelog generator

### Installation
```bash
brew install git-cliff
cargo install git-cliff
```

### Usage
```bash
# Generate changelog
git cliff --output CHANGELOG.md
# Unreleased changes
git cliff --unreleased
# From specific tag
git cliff --tag v1.0.0
# With custom config
git cliff --config cliff.toml
# Preview
git cliff --preview
# Next version
git cliff --bump --unreleased --tag v2.0.0
```

### Configuration (TOML)
```toml
# cliff.toml
[changelog]
header = "# Changelog\n"
body = """
{% for group, commits in commits | group_by(attribute="group") %}
### {{ group | upper_first }}
{% for commit in commits %}
- {{ commit.message | upper_first }}
{%- endfor %}
{% endfor %}
"""
trim = true

[git]
conventional_commits = true
filter_unconventional = true
```

### Key Differentiators
- **Conventional Commits** support
- **Template-based** — customizable output format
- **Semantic versioning** bump detection
- **Multiple output formats** — Markdown, JSON, TOML
- **Git integration** — reads tags, commits
- **CI/CD friendly**

---

## 7. onefetch — Git Repo Summary

**GitHub**: https://github.com/o2sh/onefetch  
**Language**: Rust  
**Purpose**: Git repository summary on your terminal

### Installation
```bash
brew install onefetch
cargo install onefetch
```

### Usage
```bash
# Show repo info
onefetch
# Specific directory
onefetch /path/to/repo
# ASCII art
onefetch --ascii
# JSON output
onefetch --json
# Disable colors
onefetch --no-color
```

### Key Features
- **Language breakdown** with bar chart
- **Code statistics** (lines, files, commits)
- **Contributor info**
- **Git history** stats
- **License detection**
- **ASCII art** logo display
- **JSON/plain output** for scripting

### Key Differentiators
- **Beautiful terminal output** with charts
- **Zero dependencies** beyond cargo
- **Fast** (Rust) — instant results
- **Customizable** — themes, colors
- **neofetch-style** for git repos

---

## 8. git-extras — Git Enhancement Suite

**GitHub**: https://github.com/tj/git-extras  
**Language**: Shell  
**Stars**: 17K+  
**Purpose**: Little git extras

### Installation
```bash
# macOS
brew install git-extras
# From source
git clone https://github.com/tj/git-extras && cd git-extras && make install
```

### Key Commands
| Command | Description |
|---------|-------------|
| `git ignore` | Generate .gitignore |
| `git changelog` | Generate changelog |
| `git release` | Tag and release |
| `git effort` | Show commit effort |
| `git summary` | Repo summary |
| `git undo` | Undo last commit |
| `git setup` | Initialize repo |
| `git fresh-branch` | Create empty branch |
| `git contrib` | Show contributors |
| `git count` | Count commits |
| `git delete-branch` | Delete branches |
| `git delete-submodule` | Remove submodule |
| `git delete-tag` | Delete tag |
| `git info` | Repo info |
| `git merge-into` | Merge into branch |
| `git pr` | Pull request |
| `git rename-file` | Rename with history |
| `git rename-tag` | Rename tag |
| `git repl` | Git REPL |
| `git reset-file` | Reset file |
| `git show-tree` | Show tree |
| `git squash` | Squash commits |
| `git touch` | Touch and add |

---

## Summary: When to Use Which Git TUI Tool

| Tool | Language | Key Strength | Best For |
|------|----------|--------------|----------|
| **Lazygit** | Go | Complete git TUI, user-friendly | Daily git work, beginners |
| **GitUI** | Rust | Blazing fast, low-memory | Performance-sensitive users |
| **tig** | C | Battle-tested, lightweight | Terminal purists, scripting |
| **gh** | Go | Full GitHub integration | GitHub workflows |
| **gh-dash** | Go | Beautiful dashboard | PR/issue tracking |
| **git-cliff** | Rust | Changelog generation | Release automation |
| **onefetch** | Rust | Repo summary display | Repo stats, READMEs |
| **git-extras** | Shell | Git power commands | Shell-savvy users |
