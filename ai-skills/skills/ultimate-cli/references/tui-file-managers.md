# TUI File Managers — Complete A-to-Z Reference

> Essential terminal file managers from the Awesome CLI Apps ecosystem (29 tools).
> Source: toolleeo/awesome-cli-apps-in-a-csv

---

## 1. yazi — Blazing Fast Terminal File Manager

**GitHub**: https://github.com/sxyazi/yazi  
**Language**: Rust  
**Stars**: 20K+  
**Purpose**: Blazing fast terminal file manager, written in Rust

### Installation
```bash
# macOS
brew install yazi
# cargo
cargo install --locked yazi-fm
# Linux
apt install yazi  # or from releases
```

### Key Features
- **Previews**: Code highlighting, images, PDFs, archives, media
- **Multiple selection modes**: visual, mouse, vim-like
- **Async tasks**: copy, move, delete run in background
- **Tabs**: multiple directory tabs
- **Custom previewers**: 30+ built-in previewers
- **Plugin system**: Lua-based plugins
- **Built-in search**: fd, ripgrep, fzf integration
- **Image preview**: Kitty, iTerm2, Sixel, Überzug++
- **Package manager**: built-in plugin management
- **Cross-platform**: Linux, macOS, Windows
- **Bullet-point ASCII** file tree

### Keybindings
| Key | Action |
|-----|--------|
| `j/k` | Up/down |
| `h/l` | Parent/enter directory |
| `g` | Go to top |
| `G` | Go to bottom |
| `o` | Open file with opener |
| `y` | Yank (copy) |
| `d` | Cut |
| `p` | Paste |
| `dd` | Delete (trash) |
| `space` | Toggle selection |
| `v` | Visual mode |
| `V` | Enter visual mode |
| `Enter` | Open directory/file |
| `~` | Go home |
| `/` | Search |
| `n/N` | Next/previous search |
| `r` | Rename |
| `c` | Create directory |
| `:` | Command mode |
| `t` | New tab |
| `Tab` | Next tab |
| `Q` | Quit |
| `Z` | Quit without saving |
| `Esc` | Cancel / close preview |

### Configuration (TOML)
```toml
# ~/.config/yazi/yazi.toml
[manager]
show_hidden = false
show_symlink = false
sort_by = "natural"
sort_sensitive = false
sort_reverse = false
linemode = "size"

[preview]
max_width = 600
max_height = 900
cache_dir = "/tmp/yazi-cache"

[opener]
# File openers
rules = [
  { name = "*.{png,jpg,jpeg,gif}", opener = "image" },
  { name = "*.{mp4,mkv,avi}", opener = "video" },
  { name = "*.{pdf}", opener = "zathura" },
]

[opener.image]
run = "xdg-open $0"

# ~/.config/yazi/keymap.toml
[manager]
prepend_keymap = [
  { on = "F", run = "plugin fd" },
  { on = "R", run = "plugin rg" },
]
```

### Plugin System
```bash
# Install plugins
ya pack add yazi-rs/plugins:full-border
ya pack add dreamsofcode-io/ripgrep

# List plugins
ya pack list

# Update plugins
ya pack update
```

### Key Differentiators
- **Fastest file manager** — written in Rust, near-instant
- **Rich previews** — code, images, videos, PDFs
- **Async operations** — non-blocking I/O
- **Lua plugins** — extensible architecture
- **Modern UX** — vim-like keybindings, tabs
- **Active development** — 20K+ stars, fast-moving

---

## 2. lf — Go File Manager

**GitHub**: https://github.com/gokcehan/lf  
**Language**: Go  
**Stars**: 8K+  
**Purpose**: Terminal file manager (ranger-inspired, written in Go)

### Installation
```bash
# macOS
brew install lf
# Go
go install github.com/gokcehan/lf@latest
# Linux (from releases)
curl -sSL https://github.com/gokcehan/lf/releases/latest/download/lf-linux-amd64.tar.gz | tar xz
```

### Key Features
- **Single binary** — no dependencies
- **Ranger-like navigation** — h/j/k/l
- **Preview support** — via configurable previewer script
- **Tab completion** in command line
- **Bookmarks**
- **Tags**
- **Bullet-point ASCII** file tree
- **Client-server architecture**
- **Custom commands** via shell
- **File operations** — copy, move, delete, rename

### Keybindings
| Key | Action |
|-----|--------|
| `j/k` | Up/down |
| `h/l` | Parent/enter directory |
| `space` | Toggle selection |
| `v` | Invert selection |
| `y` | Yank (copy) |
| `d` | Cut |
| `p` | Paste |
| `dd` | Delete (trash) |
| `r` | Rename |
| `zh` | Toggle hidden |
| `:` | Command mode |
| `!` | Shell command |
| `$` | Run command in $SHELL |
| `s` | Run command in pager |
| `q` | Quit |

### Configuration
```bash
# ~/.config/lf/lfrc
set hidden true
set previewer ~/.config/lf/pv.sh
set cleaner ~/.config/lf/clean.sh
set ifs "\n"

# Preview script
cmd preview ${{
    if [ -f "$1" ]; then
        bat --style=numbers --color=always "$1"
    fi
}}

# Custom commands
map . set hidden!
map D delete
map X cut
```

### Key Differentiators
- **Minimal dependencies** — single Go binary
- **ranger-compatible** keybindings
- **Client-server** architecture for persistent state
- **Easily scriptable** via shell
- **Lightweight** — ~5MB binary

---

## 3. ranger — Python File Manager

**GitHub**: https://github.com/ranger/ranger  
**Language**: Python  
**Stars**: 16K+  
**Purpose**: Console file manager with VI keybindings

### Installation
```bash
# macOS
brew install ranger
# Linux
apt install ranger
# pip
pip install ranger-fm
```

### Key Features
- **VI keybindings** — h/j/k/l navigation
- **Column layout** — multi-column directory view
- **Preview pane** — file content preview
- **File type detection** — coloring by type
- **Custom scripts** — Python plugins
- **Image preview** — via w3mimgdisplay
- **Bookmarks**
- **Tags**
- **Rifle** — powerful file opener
- **Bulk rename**

### Keybindings
| Key | Action |
|-----|--------|
| `j/k` | Up/down |
| `h/l` | Parent/enter |
| `gg/G` | Top/bottom |
| `yy` | Copy |
| `dd` | Cut |
| `pp` | Paste |
| `dd` | Cut |
| `/` | Search |
| `i` | Display file |
| `r` | Open with rifle |
| `z` | Change settings |
| `zh` | Toggle hidden |
| `: ` | Command line |
| `S` | Shell in directory |
| `Q` | Quit |

### Configuration
```python
# ~/.config/ranger/rc.conf
set preview_images true
set preview_images_method kitty
set use_preview_script true
set draw_borders true

# Rifle config: ~/.config/ranger/rifle.conf
ext pdf = zathura "$1"
ext md = glow "$1"
ext {png,jpg,jpeg} = feh "$1"
```

### Key Differentiators
- **Most established** — 16K+ stars, 10+ years old
- **Python extensibility** — rich plugin ecosystem
- **Column view** — unique multi-column directory display
- **Rifle opener** — smart file type associations
- **Bulk rename** built-in
- **Vast documentation** and community

---

## 4. nnn — Tiny, Lightning-Fast

**GitHub**: https://github.com/jarun/nnn  
**Language**: C  
**Stars**: 20K+  
**Purpose**: Tiny, lightning-fast, feature-packed file manager

### Installation
```bash
# macOS
brew install nnn
# Linux
apt install nnn
# From source
git clone https://github.com/jarun/nnn && cd nnn && make install
```

### Key Features
- **Extremely small** — ~100KB binary
- **Zero dependencies** for basic use
- **Notebook-style** context switcher
- **Disk usage analyzer** mode
- **File launcher** — opener plugin
- **Batch rename**
- **Session management**
- **SSHFS mount**
- **Type-to-nav** — filter while typing
- **Trash/clobber** protection
- **FIFO previewer** integration
- **Plugins** — 30+ community plugins

### Keybindings
| Key | Action |
|-----|--------|
| `j/k` | Up/down |
| `h/l` | Parent/enter |
| `^J/^K` | Scroll preview |
| `/` | Filter |
| `Alt+<n>` | Context tab |
| `p` | Open with program |
| `Space` | Select |
| `y` | List selection |
| `r` | Rename |
| `x` | Delete |
| `s` | Manage session |
| `E` | Edit file |
| `P` | Open with pager |
| `d` | Detail view toggle |
| `o` | Sort options |
| `q` | Quit |

### Configuration
```bash
# Environment variables
export NNN_TRASH=1                # Use trash instead of rm
export NNN_SHOW_HIDDEN=1          # Show hidden files
export NNN_ORDER=1                # Directory first
export NNN_PLUG='p:preview-tui;f:fzopen;d:dragdrop'

# In ~/.config/nnn/plugins.ini
[plugins]
preview-tui = preview-tui
fzopen = fzopen
```

### Key Differentiators
- **Smallest footprint** — ~100KB, instant startup
- **C-based** — zero dependencies
- **Notebook contexts** — 4 independent workspace tabs
- **Disk usage analyzer** built-in
- **SSHFS support** — browse remote servers
- **Type-to-nav** — filter as you type
- **Massively configurable** via environment variables

---

## 5. broot — Tree Navigator

**GitHub**: https://github.com/Canop/broot  
**Language**: Rust  
**Stars**: 11K+  
**Purpose**: A new way to see and navigate directory trees

### Installation
```bash
# macOS
brew install broot
# cargo
cargo install broot
# Linux (Debian/Ubuntu)
apt install broot
```

### Key Features
- **Tree view** — starts with directory tree
- **Fuzzy search** — instant file finding
- **Permanent tree** — always visible, not just list
- **Customizable previews**
- **File operations** — copy, move, delete
- **Git integration** — status colors
- **Terminal panel integration**
- **Shell integration** — `br` for cd
- **Verb shortcuts** — custom commands

### Keybindings
| Key | Action |
|-----|--------|
| `↑/↓` | Up/down |
| `→` | Enter directory |
| `←` | Go up |
| `Typing` | Fuzzy filter |
| `Esc` | Clear filter / back |
| `Enter` | Open file |
| `Alt+Enter` | Change to directory and exit |
| `: ` | Command mode |
| `?` | Help |
| `q` | Quit |

### Key Commands
| Command | Description |
|---------|-------------|
| `:cd` | cd to selection and exit |
| `:cp` | Copy file |
| `:mv` | Move file |
| `:rm` | Delete file |
| `:mkdir` | Create directory |
| `:touch` | Create file |
| `:rename` | Rename |
| `:toggle_hidden` | Show/hide hidden |
| `:toggle_git_file_info` | Git status |
| `:open_leave` | Open and exit |
| `:quit` | Quit |

### Configuration (Hjson/TOML)
```toml
# ~/.config/broot/conf.toml
[verbs]
invocation = "edit"
apply_to = "file"
cmd = "nvim {file}"

invocation = "git_status"
apply_to = "file"
cmd = "git status {file}"

[skin]
tree = "rgb(100, 149, 237)"
file = "rgb(173, 216, 230)"
directory = "rgb(255, 255, 255)"
exe = "rgb(50, 205, 50)"
link = "rgb(255, 215, 0)"
pruning = "rgb(105, 105, 105)"
permissions = "rgb(100, 100, 100)"
size = "rgb(100, 100, 100)"
```

### Shell Integration
```bash
# Add to ~/.bashrc / ~/.zshrc
source ~/.config/broot/launcher/bash/br

# Now use br instead of cd
br
# Or use broot for selection
cd $(broot --cmd ":cd")
```

### Key Differentiators
- **Tree-first** — starts with a visual tree (not list)
- **Fuzzy search** — type to instantly filter
- **Git integration** — status colors on files
- **Shell integration** — `br` replaces `cd`
- **Custom verbs** — add your own commands
- **Skin system** — fully customizable colors
- **File operations** — all from within broot

---

## 6. joshuto — Rust File Manager

**GitHub**: https://github.com/kamiyaa/joshuto  
**Language**: Rust  
**Stars**: 4K+  
**Purpose**: Ranger-like terminal file manager in Rust

### Installation
```bash
cargo install joshuto
```

### Key Features
- **Ranger-inspired** UI and keybindings
- **Tab support**
- **Previews** with configurable previewer
- **Bulk rename**
- **File operations** — copy, move, delete, rename
- **Trash support**
- **Sort options**
- **Bookmarks**

### Keybeinding Similarities
- **Same keybindings as ranger** for most operations
- `j/k` for navigation
- `h/l` for parent/enter
- `yy/dd/pp` for copy/cut/paste

---

## 7. clifm — KISS File Manager

**GitHub**: https://github.com/leo-arch/clifm  
**Language**: C  
**Stars**: 3K+  
**Purpose**: KISS file manager with unique `;` command syntax

### Installation
```bash
# Linux
apt install clifm
# From source
git clone https://github.com/leo-arch/clifm && cd clifm && make install
```

### Key Features
- **Unique `;` prefix** — commands are just `;cmd`
- **No external dependencies**
- **Single binary** (~300KB)
- **Customizable prompts**
- **Tags**
- **Built-in file search**
- **Trash system**
- **Encryption** support
- **Plugin system** via shell scripts
- **Colors and icons**

### Key Differentiators
- **Unique command syntax** — `;cd`, `;rm`, `;cp`
- **Minimalist philosophy** — "keep it simple"
- **Very small** — ~300KB
- **No mouse required** — fully keyboard-driven
- **POSIX compliant**

---

## Summary: When to Use Which File Manager

| Tool | Language | Size | Key Strength | Best For |
|------|----------|------|--------------|----------|
| **yazi** | Rust | ~10MB | Blazing fast, rich previews | Modern users, best overall |
| **lf** | Go | ~5MB | Single binary, ranger-like | Minimalists |
| **ranger** | Python | ~2MB+ | Python extensibility | Plugin power users |
| **nnn** | C | ~100KB | Tiny, instant startup | Resource-constrained |
| **broot** | Rust | ~8MB | Tree-first, fuzzy search | Navigators, cd replacement |
| **joshuto** | Rust | ~5MB | Ranger in Rust | Rust enthusiasts |
| **clifm** | C | ~300KB | Unique syntax, minimal | CLI purists |
