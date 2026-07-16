# Awesome CLI Apps — Complete Guide (2230 Tools / 81 Categories)

> Source: https://github.com/toolleeo/awesome-cli-apps-in-a-csv
> The largest curated list of CLI/TUI programs, maintained as structured CSV data.
> 2230 apps across 81 categories — this file covers the most CLI-development-relevant categories.

---

## 1. Fuzzy Finders & Option Pickers (18 tools)

Interactive tools for fuzzy-searching and selecting from lists.

| Tool | Author | Description |
|------|--------|-------------|
| **fzf** | junegunn | The gold standard fuzzy finder. General-purpose, vim integration, previews. |
| **fzy** | jhawthorn | "Better fuzzy finder" — simpler, faster scoring algorithm. |
| **skim** | lotabout | Fuzzy finder in Rust (fzf-compatible). |
| **television** | alexpasmantier | Blazing fast general-purpose fuzzy finder TUI. |
| **choose** | jagprog5 | NCurses-based token selector. |
| **percol** | mooz | Python-based interactive filter — pipe stdin to interactive selection. |
| **pick** | mptre | Fuzzy search option picker. |
| **smenu** | p-gen | Terminal menu generator that evolved into a powerful selection tool. |
| **fnf** | leo-arch | Interactive fuzzy finder with instant sorted updates. |
| **luneta** | fbeline | Interactive filter composable within scripts. |
| **cmenu** | 10xJSChad | Minimal dmenu-like TUI menu utility. |
| **pmenu** | sgtpep | Dynamic terminal-based menu inspired by dmenu. |
| **shmenu** | duclos-cavalcanti | Menu TUI tool written solely in bash. |
| **lSel** | unsigned-enby | Simple no-fuss TUI selection menu for scripts. |

**Key Insight**: `fzf` is the de facto standard — it has the richest ecosystem (extensions, vim integration, git integration). `skim` is a Rust-native drop-in replacement. `television` is the newest contender with blazing speed.

---

## 2. Git Tools & Accessories (81 tools)

Tools that extend, enhance, or replace git workflows.

### Git TUI Clients
| Tool | Description |
|------|-------------|
| **Lazygit** | Simple terminal UI for git commands (jesseduffield) |
| **GitUI** | Keyboard-only git TUI with scalable UI (extrawurst) |
| **tig** | ncurses-based text-mode git browser (jonas) |
| **grv** | Git Repository Viewer with refs/commits/diffs (rgburke) |
| **pyautogit** | Python-based git TUI (jwlodek) |
| **Froggit** | Minimalist Git TUI with GitHub CLI integration |

### Git Enhancements
| Tool | Description |
|------|-------------|
| **git-extras** | Little git extras: git-ignore, git-changelog, git-release, git-effort (tj) |
| **forgit** | fzf-powered interactive git utility (wfxr) |
| **git absorb** | Automatic `git commit --fixup` (tummychow) |
| **git-secret** | Encrypted secrets in git repos using PGP |
| **git-bug** | Distributed offline-first bug tracker embedded in git |
| **git-cliff** | Highly customizable changelog generator |
| **onefetch** | Git repository summary on your terminal |
| **mergestat-lite** | Run SQL queries on git repositories |
| **gitleaks** | Detect hardcoded secrets in git repos |

### GitHub CLI Extensions
| Tool | Description |
|------|-------------|
| **gh** (GitHub CLI) | Official GitHub CLI — repos, issues, PRs, projects |
| **gh-dash** | Beautiful CLI dashboard for GitHub (dlvhdr) |
| **gh-s** | Interactive GitHub repo search (gennaro-tedesco) |
| **gh-f** | fzf extension for gh CLI |
| **gh-stars** | Show repository stargazers |

### Git Stats & Visualization
| Tool | Description |
|------|-------------|
| **git-stats** | Local git statistics with GitHub-like contribution calendars |
| **git-quick-stats** | Efficient git statistics |
| **git-recall** | Easy commit browsing |
| **Kusa** | Display GitHub contribution graphs |
| **stargazer** | GitHub stats from the command line |
| **Git Activity Visualization** | Git activity heat map |
| **GitUI** | Full git operations with keyboard control |

**Key Insight**: Lazygit is the most popular git TUI (25K+ stars). GitUI is a fast Rust alternative. The `gh` CLI ecosystem has exploded with extensions (gh-dash, gh-s, gh-f). `gitleaks` is essential for security.

---

## 3. Prompts & Shell Themes (14 tools)

Terminal prompt customization and welcome messages.

| Tool | Description |
|------|-------------|
| **Starship** | The cross-shell prompt (minimal, fast, any shell) |
| **Oh My Posh** | Most customizable cross-platform prompt renderer |
| **Powerlevel10k** | Fast Zsh theme with rich configuration |
| **Spaceship** | Minimalistic, powerful Zsh prompt |
| **Pure** | Pretty, minimal, fast ZSH prompt |
| **Liquid Prompt** | Carefully designed prompt with contextual changes |
| **powerline** | Statusline for vim, zsh, bash, tmux, and more |
| **geometry** | Minimalistic, fully customizable Zsh prompt |
| **blaze** | Customizable prompt for bash, zsh, fish |
| **Polyglot Prompt** | Dynamic prompt for 10+ shells |
| **Basta!** | Scroll-protected status line at bottom of terminal |
| **bashorg-motd** | 10K+ bash.org quotes on terminal open |

**Key Insight**: Starship is the modern default — works with any shell, has the widest ecosystem. Powerlevel10k is the Zsh power-user choice. Oh My Posh leads in customizability.

---

## 4. Programming & Development Tools (74 tools)

Tools for developers: debuggers, testing, linting, code generation.

| Tool | Description |
|------|-------------|
| **cloc** | Count lines of code (blank, comment, physical) |
| **ast-grep** | Code structural search, lint and rewriting |
| **ChatDBG** | AI-assisted debugging with LLM-powered answers |
| **airs** | Live reload for Go apps |
| **argbash** | Bash argument parsing code generator |
| **bashly** | Generate feature-rich bash CLIs (Ruby-based) |
| **bencher** | Continuous benchmarking platform |
| **Fastmod** | Large-scale codebase refactoring tool (Facebook) |
| **Cppcheck** | Static analysis for C/C++ |
| **gdb-dashboard** | Modular visual interface for GDB |
| **DEM** | Containerized development environment manager |
| **devbox** | Instant development environments (Jetify) |
| **Flox** | Developer environments you can take with you |
| **Euporie** | Run Jupyter notebooks from the terminal |
| **dotenvhub** | Central .env file management TUI |
| **dtool** | Collection of development tools |
| **fmake** | Make interface for any build system |
| **Crush** | Flexible AI coding agent (Charmbracelet) |
| **Blinkenlights** | TUI debugger for x86_64-linux / i8086 |
| **cgdb** | Console front-end to GDB |

**Key Insight**: `cloc` is the most universal tool. `ast-grep` is revolutionary for structural code search. `ChatDBG` represents the future of AI debugging. `bashly` is valuable for generating production-grade bash CLIs.

---

## 5. Text Processing (59 tools)

Text manipulation, formatting, extraction, and analysis.

### Essential Substitutes
| Tool | Description |
|------|-------------|
| **choose** | Human-friendly alternative to `cut` and `awk` |
| **hck** | Sharp cut clone (faster alternative to `cut`) |
| **xcut** | Extended version of Unix `cut` |
| **tuc** | Advanced cut with negative indexing, formatting |
| **sd** | Intuitive find & replace CLI, sed replacement |
| **teip** | Select partial stdin and replace with command result |

### Formatting & Color
| Tool | Description |
|------|-------------|
| **grc** | Generic colourizer for log files |
| **rich** (rich-cli) | Fancy terminal output toolbox (Textualize) |
| **lolcat** | Rainbow colorize output |
| **sprinkles** | Randomly colors input text |
| **detect-indent-cli** | Detect indentation style |

### Data Extraction
| Tool | Description |
|------|-------------|
| **pup** | Parse HTML at the command line |
| **squeeze** | Extract rich info from text (raw, JSON, HTML, YAML) |
| **yek** | Serialize repo files for LLM consumption (Rust) |
| **Ultimate Plumber** | Interactive pipeline building |
| **rare** | Real-time regex extraction + histograms |
| **trurl** | URL parsing and manipulation (curl project) |

### Utilities
| Tool | Description |
|------|-------------|
| **gzip-size-cli** | Get gzipped file size |
| **fullname-cli** | Get the fullname of the current user |
| **brok** | Find broken links in text documents |
| **deadlink** | Check HTTP URLs in files |
| **anew** | Add new lines to files, skipping duplicates |
| **kill-tabs** | Kill all Chrome tabs |
| **espanso** | Cross-platform text expander (Rust) |
| **toolong** | TUI to view, tail, merge, search logs (Textualize) |
| **grits** | Log parser/filter/formatter |

**Key Insight**: `sd` is the best modern `sed` replacement (Rust, intuitive). `choose` replaces `cut`. `pup` replaces regex for HTML parsing. `yek` is built for LLM context preparation. `toolong` is a game-changer for log viewing.

---

## 6. Text Search (18 tools) — Grep Alternatives

| Tool | Description |
|------|-------------|
| **ripgrep** (rg) | The gold standard — recursive regex search (Rust) |
| **ripgrep-all** | rg + PDF, eBooks, Office docs, zip |
| **ugrep** | Ultra-fast grep with interactive TUI, fuzzy search |
| **ag** (The Silver Searcher) | Source-code focused, skips .git |
| **ack** | Grep optimized for programmers (Perl) |
| **sift** | Fast, powerful grep alternative |
| **ast-grep** | Code structural search / lint / rewrite |
| **semantic-grep** | Grep for words with similar meaning |
| **hae** | Natural language grep |
| **vgrep** | User-friendly pager for grep |
| **krep** | SIMD-accelerated text search |
| **hypergrep** | Hyperscan-based recursive regex search |
| **zfind** | SQL-WHERE filter file search (also inside archives) |

**Key Insight**: `ripgrep` is the undisputed standard — used by VS Code and most editors. `ugrep` adds TUI interactivity. `ast-grep` is the future of code-aware search (AST-based, not regex). `ripgrep-all` extends search capabilities dramatically.

---

## 7. Data Management (103 tools across 3 sub-categories)

### General Data Management (18 tools)
| Tool | Description |
|------|-------------|
| **gnuplot** | 2D/3D data plotting |
| **sampler** | Shell command execution + visualization + alerting |
| **ttyplot** | Real-time plotting from stdin |
| **IRedis** | Interactive Redis CLI with autocomplete |
| **ROAPI** | Auto-spin read-only APIs for static datasets |
| **zq** | Zed language for pipeline search and analytics |
| **datadash** | Visualize/graph data in terminal |
| **lowcharts** | Simple data charts in terminal |
| **WOPR** | Terminal reports, presentations, infographics |

### JSON & YAML Management (48 tools)
| Tool | Description |
|------|-------------|
| **jq** | The gold standard JSON processor (sed for JSON) |
| **gojq** | Pure Go implementation of jq |
| **jaq** | jq clone in Rust |
| **fx** | Interactive JSON viewer |
| **jless** | JSON file reader/browser |
| **gron** | Convert JSON to discrete grep-able assignments |
| **jc** | Serialize CLI tool output to JSON |
| **jid** | Interactive JSON drill-down |
| **jnv** | Interactive jq filter |
| **dasel** | Query/modify JSON, YAML, TOML, XML with selectors |
| **yq** | jq for YAML (the most popular YAML processor) |
| **yaml-cli** | Simple YAML processing |
| **graphtage** | Semantic diff for JSON, XML, YAML, HTML, CSS |
| **jo** | Create JSON objects from CLI |
| **groq-cli** | GROQ query language for JSON |
| **jtc** | JSON manipulation and transformation |
| **jtbl** | Print JSON as terminal tables |
| **Frontmatter CLI** | Manage YAML frontmatter in text files |

### Tabular Data Management (37 tools)
| Tool | Description |
|------|-------------|
| **xsv** | Fast CSV command-line toolkit (Rust) |
| **csvkit** | CSV toolkit (Python) |
| **q** | Run SQL on CSV/TSV files |
| **visidata** | Interactive multitool for tabular data |
| **mlr** (Miller) | CSV/JSON processing like sed/awk/cut/join |
| **tabulous** | Pretty CSV viewer in terminal |
| **csvdiff** | Semantic CSV diff tool |
| **pup** | Parse HTML tables at CLI |
| **diffr** | CSV-aware diff |
| **tsv-utils** | eBay's CSV/TSV processing toolkit |

**Key Insight**: `jq` is indispensable for JSON. `yq` is its YAML counterpart. `dasel` is the most versatile (handles JSON, YAML, TOML, XML with one tool). `xsv` and `csvkit` are the CSV standards. `visidata` is a uniquely powerful interactive data exploration tool.

---

## 8. File & Filesystem Tools

### File Managers (29 tools)
| Tool | Description |
|------|-------------|
| **ranger** | File manager with VI keybindings |
| **lf** | Fast file manager (Go, ranger-inspired) |
| **nnn** | Tiny, lightning-fast file manager |
| **yazi** | Blazing fast terminal file manager (Rust) |
| **vifm** | VI-inspired file manager |
| **broot** | Tree view + fuzzy file navigation |
| **joshuto** | Rust file manager (ranger-like) |
| **clifm** | KISS file manager with unique `;` command syntax |
| **fm** | Terminal file manager with keyboard shortcuts |

### File Finding (10 tools)
| Tool | Description |
|------|-------------|
| **fd** | Fast and user-friendly find alternative (Rust) |
| **broot** | Tree view + fuzzy file finder |
| **fselect** | SQL-like file search |
| **zoxide** | Smarter cd command (remembers directories) |
| **autojump** | cd with learned directory weights |

### File Listing (12 tools)
| Tool | Description |
|------|-------------|
| **exa** / **eza** | Modern ls replacement (Rust) |
| **lsd** | Next-gen ls with icons, colors |
| **dua-cli** | Disk usage analyzer |
| **dust** | More intuitive du (Rust) |
| **ncdu** | NCurses disk usage analyzer |

### File Watching (8 tools)
| Tool | Description |
|------|-------------|
| **entr** | Run arbitrary command when files change |
| **watchexec** | Watch files and execute commands (Rust) |
| **fswatch** | Cross-platform file change monitor |
| **inotify-tools** | Linux inotify wrapper |

**Key Insight**: `fd` is the near-universal `find` replacement. `yazi` is the newest and fastest file manager (Rust, 15K+ stars). `zoxide` is the best `cd` replacement. `eza` (fork of `exa`) is the modern `ls`.

---

## 9. System & Process Monitoring (85 tools across 2 categories)

### Process Viewers & Monitoring (30 tools)
| Tool | Description |
|------|-------------|
| **htop** | Interactive process viewer |
| **btm** (bottom) | Cross-platform graphical process viewer (Rust) |
| **procs** | Modern ps replacement |
| **glances** | System monitoring with web UI |
| **ytop** | TUI system monitor |
| **zenith** | Top-like system monitor |
| **procs** | Modern ps with color, columns |
| **s-tui** | Terminal CPU stress and monitoring |

### System Monitoring (55 tools)
| Tool | Description |
|------|-------------|
| **neofetch** | System info displayed with ASCII logo |
| **fastfetch** | Fast neofetch replacement |
| **bandwhich** | Terminal bandwidth utilization tool |
| **ifstat** | Network interface statistics |
| **speedtest-cli** | Internet speed test |
| **ping** with enhancements | Network diagnostics |
| **nload** | Network traffic monitor |

**Key Insight**: `btm` (bottom) is the modern `htop` replacement. `procs` is the modern `ps`. `fastfetch` has largely replaced `neofetch` (faster, more maintained).

---

## 10. Security & Encryption (42 tools)

| Tool | Description |
|------|-------------|
| **gitleaks** | Detect secrets in git repos |
| **trivy** | Comprehensive vulnerability scanner |
| **snyk** | Find/fix vulnerabilities |
| **openssl** | Cryptographic toolkit |
| **age** | Simple modern file encryption |
| **gpg** | GNU Privacy Guard |
| **sslyze** | SSL/TLS configuration scanner |
| **nmap** | Network discovery and security scanning |
| **wireshark** (tshark) | Network protocol analyzer CLI |
| **lynis** | Security auditing tool |
| **hashcat** | Password recovery tool |
| **john** (John the Ripper) | Password security auditing |
| **pass** (password-store) | Standard Unix password manager |
| **gopass** | Team-based password manager |

---

## 11. Shells & Terminal Emulators (27+27 tools)

### Shells
| Tool | Description |
|------|-------------|
| **zsh** | Extended Bourne shell with rich plugin ecosystem |
| **fish** | Smart, user-friendly shell with autosuggestions |
| **nushell** | Shell with structured data (tables, lists) |
| **bash** | The ubiquitous GNU shell |
| **elvish** | Expressive scripting language + shell |
| **xonsh** | Python-powered shell |
| **oil** | A new Unix shell |
| **powershell** | Microsoft's cross-platform shell |

### Terminal Emulators
| Tool | Description |
|------|-------------|
| **alacritty** | GPU-accelerated terminal emulator (Rust) |
| **kitty** | Fast, feature-rich GPU terminal |
| **wezterm** | GPU-accelerated with multiplexing |
| **tmux** | Terminal multiplexer |
| **zellij** | Modern terminal workspace (Rust) |
| **screen** | Traditional terminal multiplexer |
| **warp** | Rust-based terminal with AI features |

**Key Insight**: `fish` leads in UX, `zsh` leads in ecosystem (oh-my-zsh). `nushell` is the most innovative — treats data as structured objects. `alacritty` is the fastest terminal. `zellij` is the modern `tmux`.

---

## 12. DevOps & Infrastructure (22 tools)

| Tool | Description |
|------|-------------|
| **terraform** | Infrastructure as code |
| **ansible** | Configuration management |
| **docker** / **docker-compose** | Container management |
| **kubectl** | Kubernetes CLI |
| **helm** | Kubernetes package manager |
| **k9s** | Kubernetes TUI dashboard |
| **stern** | Multi-pod log tailing for K8s |
| **k3s** | Lightweight Kubernetes |
| **minikube** | Local Kubernetes |
| **vagrant** | Development environment management |
| **packer** | Image builder |
| **pulumi** | Modern IaC with real programming languages |

---

## 13. Web Development (35 tools)

| Tool | Description |
|------|-------------|
| **httpie** | User-friendly HTTP client (curl alternative) |
| **curlie** | curl with httpie-like interface |
| **wrk** | HTTP benchmarking tool |
| **vegeta** | HTTP load testing |
| **siege** | HTTP regression testing and benchmarking |
| **nghttp2** | HTTP/2 tools |
| **websocat** | WebSocket client |
| **localtunnel** | Expose localhost to the web |
| **ngrok** | Secure introspectable tunnels |
| **serve** | Static file serving |
| **browser-sync** | Live browser reload |
| **postman-cli** | Postman from the terminal |

---

## 14. Editors (34 tools)

| Tool | Description |
|------|-------------|
| **vim** / **neovim** | Modal text editor |
| **emacs** | Extensible editor (terminal mode) |
| **helix** | Modal editor with LSP (Rust) |
| **zed** | High-performance editor (GPU, Rust) |
| **nano** | Simple, beginner-friendly editor |
| **micro** | Modern terminal editor |
| **vis** | Vim-like editor with Lua |
| **kakoune** | Modal editor with multi-cursor |
| **amp** | Vim-inspired terminal editor |

**Key Insight**: `neovim` is the modern vim. `helix` is the newest paradigm (selection-first, built-in LSP). `micro` is for users who want nano simplicity with modern defaults.

---

## 15. AI / ChatGPT CLI Tools (47+16+12 = 75 tools)

### AI Chat Clients (47 tools)
| Tool | Description |
|------|-------------|
| **ollama** | Run LLMs locally |
| **aichat** | ChatGPT/GPT-4 in terminal |
| **fabric** | Framework for augmenting humans using AI |
| **Mods!** (charmbracelet) | AI for the command line, built for pipelines |
| **Gemini CLI** (Google) | Official Gemini CLI |
| **mcp-manager** | Manage MCP servers across clients |
| **Blitzdenk** | Multi-provider coding agent TUI |
| **Elia** | Terminal ChatGPT with Textual |
| **tenere** | TUI for LLMs in Rust |
| **OrChat** | OpenRouter CLI |
| **termite** | Generative UI in terminal |
| **kb** (kwaak) | Autonomous AI agents on your code |
| **Chatblade** | Versatile ChatGPT CLI |
| **cai** | Fastest CLI for prompting multiple LLMs |
| **parllama** | Ollama management TUI |
| **Instrukt** | Integrated AI environment in terminal |
| **leettools** | AI search tools |

### AI Command Generators (16 tools)
| Tool | Description |
|------|-------------|
| **cmd-ai** | Natural language shell command generator |
| **OpenCode** | AI coding agent for terminal |
| **wut** | Terminal assistant that explains last command |
| **ht** | Answer shell command questions |
| **Spren** | Natural language to shell commands |
| **reTermAI** | Smart CLI assistant with LLM |

### Co-pilots (12 tools)
| Tool | Description |
|------|-------------|
| **aider** | AI pair programming in terminal |
| **Open Interpreter** | Code Interpreter in terminal |
| **codemancer** | GPT-4 code from CLI |
| **Yai** | Your AI terminal assistant |
| **CLI Co-Pilot** | GPT-4 → shell commands |

**Key Insight**: `ollama` is the standard for local LLMs. `fabric` for AI-powered patterns. `aider` for AI pair programming. `mcp-manager` for MCP server management.

---

## 16. Diff Tools (12 tools)

| Tool | Description |
|------|-------------|
| **delta** | Syntax-highlighting pager for git, diff, grep |
| **difftastic** | Structural diff that understands syntax |
| **icdiff** | Improved colored diff |
| **colordiff** | Color wrapper for diff |
| **diff-so-fancy** | Good-looking diffs |
| **graphtage** | Semantic diff for structured files |
| **dunk** | Prettier git diffs (darrenburns) |

**Key Insight**: `delta` is the most popular diff tool (used with git). `difftastic` is the most innovative — parses ASTs, not lines. `graphtage` handles structured data diffs.

---

## 17. Backup (20 tools)

| Tool | Description |
|------|-------------|
| **restic** | Fast, efficient, secure backup |
| **borg** | Encrypted backups with deduplication |
| **Kopia** | Cross-platform with GUI + CLI |
| **bupstash** | Secure, encrypted with deduplication |
| **rclone** | Cloud storage sync |
| **rsync** | Fast file copying/syncing |
| **duplicity** | GPG encrypted backups |
| **ZnapZend** | ZFS backup tool |

---

## 18. Markdown Tools (10 tools)

| Tool | Description |
|------|-------------|
| **glow** | Markdown renderer in terminal |
| **mdless** | Markdown pager like less |
| **mdv** | Terminal markdown viewer |
| **cmark** | CommonMark reference parser |
| **pandoc** | Universal document converter |

---

## Summary: CLI Development Ecosystem
The Awesome CLI Apps list provides the definitive landscape of 2230 CLI/TUI tools. Key takeaways for CLI builders:

1. **Rust is dominant**: Most new fast tools are written in Rust (fd, ripgrep, bat, delta, skim, teip, sd, bottom, yazi, etc.)
2. **Go is strong**: Many infrastructure tools in Go (gh, lazygit, k9s, etc.)
3. **AI is explosive**: 75+ AI CLI tools across 3 categories — the fastest-growing area
4. **Modern replacements**: Almost every Unix classic has a modern Rust/Go replacement
5. **Interactive TUIs are trending**: TUIs (not just CLIs) are becoming the norm for developer tooling
