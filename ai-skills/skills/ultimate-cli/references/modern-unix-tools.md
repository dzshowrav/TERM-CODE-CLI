# Modern Unix CLI Tools — Complete A-to-Z Reference

> Modern Rust/Go replacements for classic Unix commands.
> Source: Awesome CLI Apps (toolleeo), 2230 tools across 81 categories.

---

## 1. fd — Modern `find` Replacement

**GitHub**: https://github.com/sharkdp/fd  
**Language**: Rust  
**npm equivalent**: `fast-glob` (for JS ecosystem)  
**Purpose**: Fast, user-friendly alternative to `find`

### Installation
```bash
# macOS
brew install fd
# Linux (Debian/Ubuntu)
apt install fd-find
# cargo
cargo install fd-find
```

### Key Options
| Flag | Description |
|------|-------------|
| `-H` | Search hidden files |
| `-I` | Ignore .gitignore rules |
| `-s` | Case-sensitive search |
| `-i` | Case-insensitive search |
| `-g` | Glob-based search (default: regex) |
| `-e <ext>` | Filter by extension |
| `-t <type>` | Filter by type: `f`=file, `d`=dir, `l`=symlink |
| `-x <cmd>` | Execute command for each result |
| `-X <cmd>` | Execute command with all results at once |
| `--max-depth <n>` | Max search depth |
| `-p` | Show full path |
| `-S <size>` | Filter by file size (e.g. `+1M`) |
| `--changed-within <time>` | Files modified within time range |
| `--changed-before <time>` | Files modified before time range |
| `--format <fmt>` | Alternative output format |

### Usage Examples
```bash
# Find files matching pattern
fd pattern
# Find TypeScript files
fd -e ts
# Find files modified in last day
fd --changed-within 24h
# Find and execute
fd -e txt -x wc -l
# Find directories named "node_modules"
fd -td node_modules
# Find files larger than 100MB
fd -S +100M
```

### Key Differentiators
- **8-9x faster** than `find` due to parallel traversal
- **Sensible defaults**: ignores .gitignore, respects .git
- **Colorized output** by default
- **Intuitive regex** syntax (no `find -regex` complexity)
- **Icon support** when combined with `lsd` or `exa`

---

## 2. ripgrep (rg) — Modern `grep` Replacement

**GitHub**: https://github.com/BurntSushi/ripgrep  
**Language**: Rust (by Andrew Gallant / BurntSushi)  
**npm equivalent**: none (native binary)  
**Purpose**: Ultra-fast recursive regex search

### Installation
```bash
# macOS
brew install ripgrep
# Linux
apt install ripgrep
# cargo
cargo install ripgrep
```

### Key Options
| Flag | Description |
|------|-------------|
| `-i` | Case-insensitive |
| `-w` | Match whole words |
| `-l` | List matching files only |
| `-c` | Count matches per file |
| `-n` | Show line numbers |
| `-C <n>` | Show n lines of context |
| `-A <n>` | Show n lines after match |
| `-B <n>` | Show n lines before match |
| `-g <glob>` | Include only matching files |
| `-t <type>` | Search only filetype (e.g. `-t ts`) |
| `-T <type>` | Exclude filetype |
| `--no-ignore` | Search in ignored files |
| `--hidden` | Search hidden files |
| `--multiline` | Match across multiple lines |
| `-U` | Search binary files |
| `--json` | JSON output |
| `--no-filename` | Omit filename (single file) |
| `--color always/never/auto` | Color control |

### Usage Examples
```bash
# Basic search
rg "pattern"
# Search TypeScript files only
rg -t ts "import.*from"
# Find files containing pattern (just filenames)
rg -l "TODO"
# Show 3 lines of context
rg -C 3 "function"
# Case-insensitive whole word
rg -iw "error"
# Search with glob include
rg -g "*.test.ts" "describe"
# Multi-line search
rg --multiline "(?s)function \w+\(.*?\)\s*\{"
# JSON output for programmatic use
rg --json "pattern"
```

### Key Differentiators
- **Used by VS Code, IntelliJ, Sublime Text** as default search
- **10-100x faster** than `grep` with SIMD optimizations
- **Automatically respects .gitignore** — no need for `--exclude-dir`
- **Recursive by default** (unlike `grep -r`)
- **File type aware** — `rg -t py` only searches Python files
- **Supports PCRE2** for advanced regex (lookaheads, backreferences)

---

## 3. sd — Modern `sed` Replacement

**GitHub**: https://github.com/chmln/sd  
**Language**: Rust  
**Purpose**: Intuitive find & replace CLI  

### Installation
```bash
# macOS
brew install sd
# cargo
cargo install sd
```

### Key Options
| Flag | Description |
|------|-------------|
| (none) | `sd <find> <replace>` — basic usage |
| `-s` | String literal mode (no regex) |
| `-f` | File flags mode |
| `-p` | Preview changes |
| `-i` | In-place replacement |
| `--flags <flags>` | Regex flags (i, m, etc.) |

### Usage Examples
```bash
# Basic find and replace
sd "old" "new" file.txt
# Regex replace
sd "foo(\d+)" "bar$1" file.txt
# In-place edit
sd -i "foo" "bar" file.txt
# Preview changes first
sd -p "pattern" "replacement" file.txt
# String literal mode (no regex special chars)
sd -s "foo.bar" "foo_bar" file.txt
# Multi-file
sd "old" "new" *.ts
```

### Key Differentiators
- **No cryptic syntax** — `sd "foo" "bar"` instead of `sed 's/foo/bar/g'`
- **Automatic global** — replaces ALL occurrences by default (sed needs `/g`)
- **Preview mode** to verify before replacing
- **Understands regex** but doesn't require escaping for basic usage
- **Consistent across platforms** — same behavior on macOS, Linux, Windows

---

## 4. bat — Modern `cat` Replacement

**GitHub**: https://github.com/sharkdp/bat  
**Language**: Rust  
**Purpose**: Cat with syntax highlighting + Git integration

### Installation
```bash
brew install bat
apt install bat
cargo install bat
```

### Key Options
| Flag | Description |
|------|-------------|
| `-A` | Show all characters (non-printable) |
| `-n` | Show line numbers |
| `-p` | Plain output (no decorations) |
| `-l <lang>` | Explicit language highlighting |
| `--theme <theme>` | Color theme |
| `--wrap <mode>` | Wrap mode: `auto`, `never`, `character` |
| `--style <style>` | Control displayed elements |
| `--paging <mode>` | Pager: `auto`, `always`, `never` |
| `-H <n>` | Highlight line n |
| `--diff` | Show diff context |
| `--tabs <n>` | Tab width |

### Usage Examples
```bash
# Read file with highlighting
bat file.ts
# Concatenate files
bat file1.ts file2.ts
# Pipe (bat acts like cat when piped)
bat file.ts | grep pattern
# Show specific lines
bat --line-range 10:30 file.ts
# List available themes
bat --list-themes
# Use as pager
env PAGER=bat
```

### Key Differentiators
- **Auto-detects language** from file extension
- **Shows Git modifications** in gutter
- **Line numbers** by default
- **Theme support** (sublime-like)
- **Automatically uses a pager** for long files (like `less -R`)
- **Alias `cat` to `bat`**: `alias cat=bat`

---

## 5. delta — Modern `diff` Viewer

**GitHub**: https://github.com/dandavison/delta  
**Language**: Rust  
**Purpose**: Syntax-highlighting pager for git, diff, and grep output

### Installation
```bash
brew install git-delta
apt install git-delta
cargo install git-delta
```

### Git Configuration
```toml
[core]
    pager = delta
[interactive]
    diffFilter = delta --color-only
[delta]
    navigate = true
    line-numbers = true
    syntax-theme = Dracula
    side-by-side = true
[merge]
    conflictStyle = diff3
```

### Key Options
| Flag | Description |
|------|-------------|
| `--side-by-side` | Side-by-side view |
| `--line-numbers` | Show line numbers |
| `--diff-so-fancy` | Emulate diff-so-fancy |
| `--navigate` | Enable n/N navigation |
| `--dark` / `--light` | Theme mode |
| `--syntax-theme <theme>` | Code theme |
| `--file-style <style>` | File path styling |
| `--hunk-header-style <style>` | Hunk header style |
| `--file-decoration-style <style>` | File decoration style |
| `--color-only` | Strip diff but keep colors |
| `--features <features>` | Activate named features |

### Usage Examples
```bash
# Use with git
git diff
git log -p
git show
# Use with ripgrep
rg --json pattern | delta
# Side-by-side diffs
git diff --delta-side-by-side
```

### Key Differentiators
- **Syntax highlighting** in diffs (not just line-level, but code-level)
- **Side-by-side view** option
- **Line numbers** in diffs
- **Git blame annotations**
- **Zero-config** with sensible defaults
- **Emacs diff-mode like navigation** (n/N for next/previous hunk)

---

## 6. procs — Modern `ps` Replacement

**GitHub**: https://github.com/dalance/procs  
**Language**: Rust  
**Purpose**: Modern process list viewer

### Installation
```bash
brew install procs
cargo install procs
```

### Key Options
| Flag | Description |
|------|-------------|
| `-a` | Show all processes (includes other users) |
| `-l` | Show process tree |
| `-w` | Wide output (no truncation) |
| `--sort <col>` | Sort by column |
| `--watch` | Watch mode (auto-refresh) |
| `--theme <theme>` | Color theme |
| `--or` | OR conditions (default AND) |
| `--pager` | Use pager for output |
| `--config-dir` | Config directory |

### Usage Examples
```bash
# List processes
procs
# Show with tree
procs -l
# Filter by process name
procs node
# Filter by user
procs --user root
# Watch mode
procs --watch
# Show Docker containers
procs --docker
# Sort by memory
procs --sort mem
```

### Key Differentiators
- **Colorful, readable columns** by default
- **Docker-aware** — shows container names
- **Automatic filtering** by process name, user, etc.
- **Tree view** without extra flags
- **TCP/UDP port display**
- **Watch mode** for monitoring

---

## 7. btm (bottom) — Modern `top`/`htop` Replacement

**GitHub**: https://github.com/ClementTsang/bottom  
**Language**: Rust  
**Purpose**: Cross-platform graphical process/system monitor

### Installation
```bash
brew install bottom
cargo install bottom
```

### Key Options
| Flag | Description |
|------|-------------|
| `-l` | Layout override |
| `-W` | Show widgets only |
| `-w` | Set which widget is visible |
| `--theme <theme>` | Theme file |
| `--color <mode>` | Default / Gray / Group |
| `--tree` | Show process tree |
| `--rate <ms>` | Refresh rate |

### Key Widgets
- CPU (per-core usage)
- Memory (RAM + swap)
- Disk I/O
- Network I/O
- Temperature sensors
- Process list
- Battery

### Key Differentiators
- **Cross-platform** (Linux, macOS, Windows)
- **Graphical charts** in terminal (CPU, memory, network, disk)
- **Mouse support** for interactive navigation
- **Customizable layout** with TUI
- **Temperature and fan speed** monitoring
- **GPU monitoring** (NVIDIA)
- **Process filtering** by name

---

## 8. choose — Modern `cut` Replacement

**GitHub**: https://github.com/theryangeary/choose  
**Language**: Rust  
**Purpose**: Human-friendly alternative to `cut` and `awk`

### Installation
```bash
brew install choose-rust
cargo install choose
```

### Key Options
| Flag | Description |
|------|-------------|
| `-f <n>` | Select field(s) (1-indexed) |
| `-o <char>` | Output separator |
| `-i <char>` | Input separator |
| `-n` | Normalize whitespace |
| `-x` | Exclusive ranges (like Python) |
| `-t` | Tab-separated mode |
| `-b` | Bytes mode |
| `-c` | Characters mode |
| `--regex` | Regex-based field split |
| `--debug` | Debug split behavior |

### Usage Examples
```bash
# Select second field
choose -f 2
# Select fields 2-4
choose -f 2:4
# Select last field
choose -f -1
# Select all but first
choose -f 2:
# Custom separator
choose -f 1 -i ':'
# Regex split
choose -f 2 --regex '\s+'
# Range with negative indexing
choose -f :-2  # drop last 2 fields
```

### Key Differentiators
- **1-indexed** (not 0-indexed like awk)
- **Negative indexing** — `-f -1` for last field (like Python)
- **Range syntax** — `2:5`, `-3:-1`, `2:`
- **Whitespace normalization** option
- **Regex-aware** splitting

---

## 9. teip — Selective Command Application

**GitHub**: https://github.com/greymd/teip  
**Language**: Rust  
**Purpose**: Select parts of stdin and replace with command output

### Installation
```bash
brew install teip
cargo install teip
```

### Key Options
| Flag | Description |
|------|-------------|
| `-g <glob>` | Select by glob pattern |
| `-f <n>` | Select by field number |
| `-o <pattern>` | Use regex to select area |
| `-c <n>` | Select by character range |
| `-s` | Replace matched area only |
| `-z` | Entire input treated as one line |

### Usage Examples
```bash
# Hash only IP addresses in a file
cat data.txt | teip -o '\d+\.\d+\.\d+\.\d+' md5sum
# Apply command to specific field
cat data.csv | teip -f 2 -- 'tr a-z A-Z'
# In-place replace of matched areas
cat log.txt | teip -o 'ERROR.*' -- 'wc -c'
```

### Key Differentiators
- **Novel concept** — like `sed` but delegates to external commands
- **Composable** — combine with any Unix tool
- **Efficient** — only processes selected areas
- **Zero-copy** for unselected parts

---

## 10. zoxide — Smarter `cd`

**GitHub**: https://github.com/ajeetdsouza/zoxide  
**Language**: Rust  
**Purpose**: Smarter cd that learns your habits

### Installation
```bash
brew install zoxide
cargo install zoxide
```

### Shell Integration
```bash
# Bash
eval "$(zoxide init bash)"
# Zsh
eval "$(zoxide init zsh)"
# Fish
zoxide init fish | source
```

### Commands
| Command | Description |
|---------|-------------|
| `z <query>` | cd to best matching directory |
| `zi <query>` | Interactive selection with fzf |
| `zq <query>` | Echo path (no cd) |
| `za <path>` | Add directory to database |
| `zr <path>` | Remove directory from database |
| `zl` | List all directories in database |
| `zx <query>` | cd with subdir matching |

### Key Differentiators
- **Frecency-based** (frequency + recency scoring)
- **Matches partial strings** — `z proj/my-api` works
- **Learns automatically** from `cd` usage
- **fzf integration** for interactive selection
- **Portable database** across machines (XDG compliant)

---

## 11. eza — Modern `ls` Replacement

**GitHub**: https://github.com/eza-community/eza  
**Language**: Rust  
**Purpose**: Modern ls with icons, colors, and tree view

### Installation
```bash
brew install eza
cargo install eza
```

### Key Options
| Flag | Description |
|------|-------------|
| `-l` | Long format |
| `-a` | Show hidden files |
| `-T` | Tree view |
| `-1` | One file per line |
| `-R` | Recursive |
| `--icons` | Show file icons |
| `--git` | Show git status |
| `--group-directories-first` | Directories first |
| `--sort <field>` | Sort by: name, size, time, ext |
| `--header` | Show column headers |
| `--color-scale` | Color-code file sizes |

### Usage Examples
```bash
# Classic ls replacement
eza -la
# Tree view with git status
eza -la --git -T --icons
# Sort by size with color scale
eza -l --sort size --color-scale
# Show only directories
eza -D
```

### Key Differentiators
- **Git status** indicators in file listing
- **Icons** for file types
- **Color scale** for file sizes and dates
- **Tree view** built-in
- **Hyperlinks** when supported
- **Active fork** of exa (which is unmaintained)

---

## 12. dust — Modern `du` Replacement

**GitHub**: https://github.com/bootandy/dust  
**Language**: Rust  
**Purpose**: More intuitive `du`

### Installation
```bash
brew install dust
cargo install du-dust
```

### Key Options
| Flag | Description |
|------|-------------|
| `-d <n>` | Max depth |
| `-n <n>` | Show n entries |
| `-X <pattern>` | Exclude by glob |
| `-x` | Only dirs with 0% or more |
| `-p` | Show progress |
| `-R` | Reverse sort (smallest first) |
| `--invert-filter` | Invert filter |
| `--apparent-size` | Show apparent (not on-disk) size |

### Usage Examples
```bash
# Show largest directories
dust
# Show top 10 with depth 2
dust -d 2 -n 10
# Exclude node_modules
dust -X node_modules
# Show only apparent sizes
dust --apparent-size
```

### Key Differentiators
- **Intuitive bar chart** visualization
- **Faster** than `du` due to parallel processing
- **Automatic sorting** by size
- **Exclude patterns** for ignoring directories

---

## 13. doggo — Modern `dig` Replacement

**GitHub**: https://github.com/mr-karan/doggo  
**Language**: Go  
**Purpose**: Command-line DNS client with human-friendly output

### Installation
```bash
brew install doggo
cargo install doggo  # via cargo
```

### Key Options
| Flag | Description |
|------|-------------|
| `--nameserver <ns>` | DNS server |
| `--type <type>` | Record type: A, AAAA, MX, etc. |
| `--class <class>` | Query class |
| `--json` | JSON output |
| `--short` | Short output |
| `--color` | Enable/disable color |
| `--query` | Show query details |
| `--timeout <secs>` | Query timeout |
| `--tls-host <host>` | DNS-over-TLS |
| `@<server>` | Per-request DNS server |

### Usage Examples
```bash
# Basic query
doggo example.com
# MX record
doggo example.com MX
# Query specific server
doggo example.com @1.1.1.1
# DNS-over-TLS
doggo example.com --tls-host dns.cloudflare.com
# JSON output
doggo example.com --json
```

---

## 14. bandwhich — Bandwidth Monitor

**GitHub**: https://github.com/imsnif/bandwhich  
**Language**: Rust  
**Purpose**: Terminal bandwidth utilization tool

### Installation
```bash
brew install bandwhich
cargo install bandwhich
```

### Key Options
| Flag | Description |
|------|-------------|
| `-a` | Show all connections |
| `-r` | Raw number output |
| `-p <pid>` | Filter by process |
| `-n <n>` | Show n top entries |
| `-i <interface>` | Network interface |
| `-d <dns>` | DNS server |
| `-t` | Total bandwidth |
| `-v` | Verbose |

### Usage Examples
```bash
# Monitor bandwidth
sudo bandwhich
# Filter by interface
sudo bandwhich -i wlan0
# Show raw numbers
bandwhich -r
# Top 5 processes
bandwhich -n 5
```

### Key Differentiators
- **Process-level** bandwidth breakdown
- **Real-time** network usage
- **Connection-level** detail (remote IP, port, etc.)
- **DNS resolution** for connections

---

## 15. duf — Modern `df` Replacement

**GitHub**: https://github.com/muesli/duf  
**Language**: Go  
**Purpose**: Disk Usage/Free with better output

### Installation
```bash
brew install duf
cargo install duf  # via go install
```

### Key Options
| Flag | Description |
|------|-------------|
| `-all` | Show pseudo-devices |
| `-hide <pattern>` | Hide mount points |
| `-only <pattern>` | Show only mount points |
| `--json` | JSON output |
| `--theme <theme>` | Color theme |
| `--sort <field>` | Sort by: size, avail, usage, mount |
| `--width <n>` | Max output width |
| `--inodes` | Show inode info |

### Usage Examples
```bash
# Show disk usage
duf
# Only ext4/xfs mounts
duf -only ext4,xfs
# JSON for scripting
duf --json
# Sort by usage
duf --sort usage
```

---

## Summary: Modern Unix Toolkit

| Classic | Modern | Language | Command | Key Advantage |
|---------|--------|----------|---------|---------------|
| `find` | `fd` | Rust | `fd` | Faster, sensible defaults |
| `grep` | `ripgrep` | Rust | `rg` | 10-100x faster, .gitignore-aware |
| `sed` | `sd` | Rust | `sd` | Intuitive, no cryptic syntax |
| `cat` | `bat` | Rust | `bat` | Syntax highlighting |
| `diff` | `delta` | Rust | `delta` | Syntax-highlighted diffs |
| `ps` | `procs` | Rust | `procs` | Colorful, Docker-aware |
| `top/htop` | `bottom` | Rust | `btm` | Graphical charts |
| `cut` | `choose` | Rust | `choose` | Python-like indexing |
| `cd` | `zoxide` | Rust | `z` | Learns your habits |
| `ls` | `eza` | Rust | `eza` | Icons, git, tree |
| `du` | `dust` | Rust | `dust` | Bar chart visualization |
| `dig` | `doggo` | Go | `doggo` | Human-friendly DNS |
| `df` | `duf` | Go | `duf` | Colorful disk usage |
