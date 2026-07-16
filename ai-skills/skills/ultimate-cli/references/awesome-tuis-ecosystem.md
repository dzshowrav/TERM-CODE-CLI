# Awesome TUIs — Complete Ecosystem Reference

> Source: rothgar/awesome-tuis (GitHub, 19.7K stars, 615 commits)
> 500+ TUI applications across 13 categories

---

## Table of Contents

1. [Overview](#overview)
2. [Dashboards (68 tools)](#1-dashboards)
3. [Development (85+ tools)](#2-development)
4. [Docker/LXC/K8s (30 tools)](#3-dockerlxck8s)
5. [Editors (24 tools)](#4-editors)
6. [File Managers (19 tools)](#5-file-managers)
7. [Games (54 tools)](#6-games)
8. [Libraries — see `awesome-tuis-libraries.md`](#7-libraries)
9. [Messaging (40 tools)](#8-messaging)
10. [Miscellaneous (70+ tools)](#9-miscellaneous)
11. [Multimedia (60+ tools)](#10-multimedia)
12. [Productivity (80+ tools)](#11-productivity)
13. [Screensavers (6 tools)](#12-screensavers)
14. [Web (30 tools)](#13-web)

---

## Overview

The awesome-tuis list catalogs **500+ TUI (Terminal User Interface) applications** across 13 categories. A TUI is an interactive terminal application — it's not just a CLI command that runs and exits, but an app you interact with continuously.

### Key Trends
1. **Rust dominates** modern TUI tooling (bottom, lazydocker, spotify-player, yazi, gitui, kdash, trippy, etc.)
2. **Go is strong** for infrastructure TUIs (k9s, lazygit, WTF, gh-dash, cointop, etc.)
3. **Python** powers data-heavy TUIs (Rich, textual, VisiData, Glances, etc.)
4. **Ink/React** powers AI CLI tools (ink has 33K stars, powers Claude Code, Cursor, etc.)
5. **TUI frameworks** are becoming more accessible (Ratatui, Bubble Tea, Textual, FTXUI)

### When to Recommend a TUI App vs Build One
| User needs... | Recommend an existing app | Build with TUI library |
|---------------|--------------------------|----------------------|
| Git management | lazygit, gitui, tig | — |
| Docker management | lazydocker, k9s, dockly | — |
| System monitoring | btop, bottom, htop | — |
| Code editing | helix, micro, kakoune | — |
| File management | yazi, lf, ranger | — |
| Custom workflow | — | Bubble Tea, Ratatui, Textual |
| AI agent in terminal | — | Ink + React |
| API client | ATAC, posting, resterm | — |

---

## 1. Dashboards (68 Tools)

### System Resource Monitors
| Tool | Language | Description |
|------|----------|-------------|
| **btop++** | C++ | Resource monitor with extras — GPU, disks, network, processes |
| **bottom** (btm) | Rust | Customizable graphical process/system monitor |
| **htop** | C | Classic interactive process viewer |
| **bashtop** | Bash | Resource manager written in bash |
| **bpytop** | Python | Python-based system monitor |
| **gotop** | Go | Terminal based graphical activity monitor |
| **zenith** | Rust | In-terminal graphical metrics |
| **glances** | Python | Comprehensive system monitoring |
| **atop** | C | Root-level system and process monitor |
| **ttop** | Python | System monitoring with historical data service |

### Network Monitors
| Tool | Language | Description |
|------|----------|-------------|
| **bandwhich** | Rust | Terminal bandwidth utilization tool |
| **trippy** | Rust | Network diagnostic tool (like mtr+) |
| **gping** | Rust | Ping with a graph |
| **bmon** | C | Network monitoring and debugging tool |
| **nethogs** | C++ | 'net top' tool |
| **Kyanos** | C | Linux network analysis based on eBPF |
| **netscanner** | Rust | Network scanner |
| **oryx** | Rust | TUI for sniffing traffic using eBPF |
| **rustnet** | Rust | Cross-platform network monitoring with DPI |
| **sockttop** | Go | Remote system monitor over WebSockets |
| **tcpterm** | Go | Packet visualizer in TUI |
| **termshark** | Go | Terminal UI for tshark/Wireshark |

### GPU/CPU Monitors
| Tool | Language | Description |
|------|----------|-------------|
| **nvtop** | C | GPU process monitoring (AMD, Intel, NVIDIA) |
| **s-tui** | Python | CPU stress and monitoring utility |
| **CoreFreq** | C | CPU monitoring for 64-bit processors |
| **macmon** | Rust | Performance monitoring for Apple Silicon |
| **tegratop** | Go | Monitoring for Nvidia Jetson boards |

### Cloud/Infrastructure Dashboards
| Tool | Language | Description |
|------|----------|-------------|
| **WTF** | Go | Personal information dashboard |
| **gh-dash** | Go | Rich TUI for GitHub PRs and Issues |
| **claws** | Go | AWS resource management with vim keybindings |
| **sacha** | Go | AWS TUI for 7 services |
| **Planor** | Go | Dashboard for AWS, Vultr, Heroku |
| **damon** | Go | HashiCorp Nomad TUI |
| **trek** | Go | ncurses explorer for Nomad clusters |
| **wander** | Go | HashiCorp Nomad terminal client |
| **cointop** | Go | Cryptocurrency tracking TUI |
| **kaskade** | Python | Kafka TUI |
| **Yozefu** | Rust | Kafka cluster data explorer |
| **dolphie** | Python | Real-time MySQL/MariaDB analytics |
| **gobang** | Rust | Cross-platform database management TUI |
| **chdig** | Rust | ClickHouse TUI interface |

### Log/Analysis Tools
| Tool | Language | Description |
|------|----------|-------------|
| **Goaccess** | C | Real-time web log analyzer |
| **gonzo** | Go | Real-time log analysis TUI (like k9s) |
| **nerdlog** | C++ | Remote-first, multi-host TUI log viewer |
| **hwatch** | Perl | Modern alternative to `watch` |
| **otel-tui** | Go | OpenTelemetry viewer |
| **Grafterm** | Go | Grafana-inspired terminal dashboards |

### Utility Dashboards
| Tool | Language | Description |
|------|----------|-------------|
| **AdGuardian-Term** | TypeScript | AdGuard Home traffic monitor |
| **binsider** | Rust | Linux binary analysis TUI |
| **cgdb** | C | Console front-end to GDB |
| **cheatshh** | Bash | fzf-based cheatsheet manager |
| **dashbrew** | TypeScript | TUI dashboard builder from scripts/APIs |
| **fubar** | Go | TUI for gtfobins power users |
| **process-compose** | Go | TUI for running apps/processes |
| **Puffin** | Scala | Terminal dashboard for hledger |
| **Raijin** | Rust | Weather TUI (no API key needed) |
| **ServerHub** | C# | Linux server monitoring dashboard |
| **sysz** | Bash | fzf TUI for systemctl |
| **tdash** | Go | Terminal dashboard (Google Analytics, GitHub, etc.) |
| **updo** | Go | Website uptime monitoring |
| **ticker** | Go | Stock/crypto price tracker |
| **tufw** | Python | Terminal UI for ufw |
| **tuicamp** | Go | Unofficial TimeCamp TUI |

---

## 2. Development (85+ Tools)

### Git & Version Control
| Tool | Language | Description |
|------|----------|-------------|
| **lazygit** | Go | Simple terminal UI for git commands |
| **gitui** | Rust | Blazing fast terminal UI for git |
| **tig** | C | Text-mode interface for git |
| **grv** | Go | Terminal interface for viewing git repos |
| **ddv** | Go | Terminal DynamoDB viewer |
| **serie** | Go | Rich git commit graph |
| **differ** | Rust | TUI git diff viewer |
| **Froggit** | Python | Minimalist Git TUI with GitHub CLI integration |
| **git-crecord** | Python | Interactive selective git commit |
| **git-scope** | Go | Dashboard for inspecting multiple git repos |
| **ggc** | Go | Terminal-based Git CLI tool |
| **gitv** | Go | Terminal client for GitHub issues |
| **serie** | Go | Rich git commit graph |

### API Clients & HTTP Tools
| Tool | Language | Description |
|------|----------|-------------|
| **ATAC** | Rust | Full-featured TUI API client (offline, account-less) |
| **posting** | Python | Powerful HTTP client in your terminal |
| **resterm** | Rust | HTTP/GraphQL/gRPC + WebSockets, SSE, OpenAPI |
| **CuTE** | Rust | Build, execute, save curl commands |
| **mitmproxy** | Python | Interactive HTTPS proxy |
| **Slumber** | Rust | Terminal-based HTTP/REST client |
| **chiko** | Rust | TUI gRPC client |
| **cnTUI** | TypeScript | Replay Chrome requests as curl |

### Database Clients
| Tool | Language | Description |
|------|----------|-------------|
| **harlequin** | Python | SQL IDE for your terminal |
| **dblab** | Go | Database client every CLI junkie deserves |
| **lazysql** | Go | Cross-platform TUI database management |
| **rainfrog** | Rust | Database management TUI (Postgres, MySQL, SQLite) |
| **sqlit** | Go | Lightweight SQL database TUI (inspired by lazygit) |
| **qrypad** | Go | Terminal SQL client (Postgres, MySQL, SQLite) |
| **termdbms** | Python | TUI for viewing and editing database files |
| **sq** | Go | Database client for vim users |
| **dbee** | Haskell | Fast & minimalistic database browser |
| **gobang** | Rust | Cross-platform TUI database management tool |

### AI Development Tools
| Tool | Language | Description |
|------|----------|-------------|
| **opencode** | TypeScript | AI coding agent for the terminal |
| **codex** | Go | Lightweight coding agent |
| **crush** | Go | AI coding agent by Charmbracelet |
| **amux** | Go | Easily run parallel coding agents |
| **Claude Code Bridge** | Python | Real-time multi-AI collaboration |
| **Claude Code Usage Monitor** | Rust | Monitor Claude token usage |
| **kagan** | Go | AI-powered Kanban for autonomous dev workflows |
| **models** | Go | TUI for browsing AI models |
| **Quorum** | Go | Multi-agent AI discussion system |
| **Toad** | TypeScript | Unified interface for AI |
| **turbostream** | Python | AI Agents for streaming data signals |
| **VT Code** | Go | Semantic coding agent |
| **fast-resume** | Go | Index and fuzzy search coding agent sessions |

### Code Analysis & Debugging
| Tool | Language | Description |
|------|----------|-------------|
| **delta** | Rust | Syntax-highlighting pager for git/diff/grep |
| **fx** | Go | Terminal JSON viewer & processor |
| **Twig** | Go | Interactive JSON/YAML file explorer |
| **jqp** | Go | TUI playground for jq queries |
| **qo** | Go | Interactive SQL filter for JSON, CSV, TSV |
| **pudb** | Python | Console-based visual Python debugger |
| **pproftui** | Go | TUI for Go's pprof profiling |
| **heretek** | Rust | GDB TUI dashboard |
| **austin-tui** | Python | Top-like TUI for Austin Python profiler |
| **blinkenlights** | C | TUI for debugging x86_64-linux/i8086 programs |

### Package & Dependency Management
| Tool | Language | Description |
|------|----------|-------------|
| **cargo-seek** | Rust | TUI for searching, adding, installing cargo crates |
| **deputui** | TypeScript | Review and install NPM package updates |
| **Feluda** | Rust | Detect restrictive licenses in dependencies |
| **toolui** | C# | Manage installed .NET nuget tools |
| **tokui** | Rust | Visualize code statistics from tokei |
| **nap** | Go | Code snippets in your terminal |
| **runme** | Go | Run code snippets from README.md |

### Infrastructure Tools
| Tool | Language | Description |
|------|----------|-------------|
| **amtui** | Go | Alertmanager TUI |
| **act3** | Go | Glance at last 3 runs of GitHub Actions |
| **catalyst** | TypeScript | Trigger GitHub Actions with matrix configs |
| **ddqa** | Go | Jira TUI for software releases |
| **prs** | Go | Stay updated on PRs |
| **brows** | Rust | CLI GitHub release browser |
| **nodebro** | Rust | View GitHub releases/tags from terminal |
| **burf** | Go | TUI for Google Cloud Storage |
| **stu** | Go | TUI for Amazon S3 |
| **csope** | C | C source code browser (cscope-based) |
| **euporie** | Python | Jupyter notebooks in terminal |
| **lean-tui** | Rust | Lean4 proof assistant visualization |
| **logradar** | Rust | Interactive log filtering and highlighting |
| **LogLens** | Rust | Structured log viewer and query engine |
| **logshark** | Go | Debugger CLI for JSON logs |
| **lazyjournal** | Go | TUI for journalctl, Docker, Podman logs |
| **proxymock** | Rust | Network recorder that auto-generates tests/mocks |
| **terraform-tui** | Go | View and interact with Terraform state |
| **vctui** | Go | Console interface for vCenter |
| **violet** | Go | TUI frontend for Vagrant commands |
| **Wikit** | Go | TUI for managing Wiki.js instances |
| **sls-dev-tools** | Node.js | Dev tools for serverless world |
| **soft-serve** | Go | Self-hostable Git server for the command line |
| **play** | Rust | TUI playground for grep, sed, awk, jq, yq |
| **regex-tui** | Python | Visualize and test regular expressions |
| **ec** | Rust | Native Git mergetool with 3 panes |
| **lazymake** | Go | Modern TUI for Makefiles |

---

## 3. Docker/LXC/K8s (30 Tools)

### Docker Management
| Tool | Language | Description |
|------|----------|-------------|
| **lazydocker** | Go | The lazier way to manage Docker |
| **dockly** | Node.js | Immersive Docker container management |
| **ctop** | Go | Top-like interface for container metrics |
| **d4s** | Go | Keyboard-driven Docker/Compose/Swarm management |
| **dry** | Go | Docker manager for terminal |
| **oxker** | Rust | Simple TUI to view & control containers |
| **Pocker** | Go | Docker task management TUI |
| **sen** | Python | Terminal UI for Docker engine |
| **docker-dash** | Go | Full TUI management for Docker |
| **dprs** | Rust | Docker containers with real-time monitoring |
| **dtop** | Go | Docker monitoring across multiple hosts |
| **DockMate** | Go | Lightweight TUI for Docker + Podman |
| **ducker** | Go | Docker TUI based on k9s |

### Kubernetes Management
| Tool | Language | Description |
|------|----------|-------------|
| **k9s** | Go | TUI for managing Kubernetes clusters |
| **kdash** | Rust | Simple and fast dashboard for Kubernetes |
| **kubetui** | Rust | Kubernetes resource monitoring |
| **ktop** | Go | Top-like tool for Kubernetes |
| **k8s-tui** | Rust | Kubernetes resource manager with multi-cluster |
| **eks-node-viewer** | Go | Visualize node usage in EKS clusters |

### K8s Tools & Utilities
| Tool | Language | Description |
|------|----------|-------------|
| **Argonaut** | Go | ArgoCD TUI |
| **kftui** | Go | Manage kubectl port-forward commands |
| **talos-pilot** | Go | TUI for Talos Linux |
| **lazytrivy** | Go | Laid-back image/k8s/filesystem scanning |
| **e1s** | Go | AWS ECS resource management |
| **etcd-walker** | Go | etcd key management |
| **cruise** | Rust | Container management TUI |
| **lazycontainer** | Swift | TUI for Apple containers |
| **Podman-tui** | Go | Podman container TUI |
| **SwarmCLI** | Go | Docker Swarm management (logs, shell, port-forwarding) |

---

## 4. Editors (24 Tools)

| Tool | Language | Description |
|------|----------|-------------|
| **helix** | Rust | Post-modern text editor (modal, built-in LSP) |
| **kakoune** | C++ | Modal text editor with interactivity focus |
| **micro** | Go | Modern, intuitive terminal-based text editor |
| **vis** | C | vi-like editor based on Plan 9 structural regex |
| **zee** | Rust | Modern text editor for terminal |
| **amp** | Rust | Complete text editor for terminal |
| **slap** | Node.js | Sublime-like terminal text editor |
| **frogmouth** | Python | Markdown browser for terminal |
| **treemd** | Go | Markdown navigator with tree-based navigation |
| **microNeo** | Go | Micro fork with in-place Markdown rendering |
| **markln** | Python | Terminal-based Markdown editor |
| **orbiton** | Go | VT100-limited text editor |
| **tilde** | C | Intuitive text editor |
| **kilo** | C | Minimal editor in ~1000 lines of C |
| **Fresh** | Go | Powerful and fast terminal text editor |
| **flow-control** | Zig | Lightning-fast text editor |
| **Durdraw** | Python | ASCII/Unicode/ANSI art editor |
| **hexed** | C | Hex editor |
| **Edit** | TypeScript | Simple editor paying homage to MS-DOS Editor |
| **C-Edit** | C | Text editor with MS-DOS Editor style menus |
| **nino** | C | Small terminal text editor |
| **maki** | C | Tabbed text editor focused on battery life |
| **turbo** | C++ | Experimental editor based on Scintilla + Turbo Vision |
| **thymus** | Python | Interactive browser & editor for network configs |

---

## 5. File Managers (19 Tools)

| Tool | Language | Description |
|------|----------|-------------|
| **yazi** | Rust | Blazing fast file manager with async I/O |
| **lf** | Go | Ranger-inspired file manager |
| **ranger** | Python | VIM-inspired file manager |
| **nnn** | C | Tiny, lightning-fast file manager |
| **broot** | Rust | Tree-based directory navigator |
| **Vifm** | C | File manager with vi keybindings |
| **mc** | C | Midnight Commander — classic orthodox file manager |
| **superfile** | Go | Fancy modern terminal file manager |
| **far2l** | C++ | Linux port of FAR v2 file manager |
| **goful** | Go | Powerful TUI file manager |
| **TUIFIManager** | Python | Cross-platform file manager (supports Termux) |
| **rovr** | Rust | Post-modern terminal file manager |
| **xplr** | Rust | Hackable, minimal TUI file explorer |
| **sfm** | Go | Simple file manager |
| **s3duck-tui** | Go | TUI S3 client |
| **ytreenova** | Go | XTree-style file manager |
| **adbtuifm** | Python | ADB-based file manager for Android |
| **fml** | Bash | Simple fast file manager in Bash |
| **deletor** | Rust | Interactive file deletion TUI |

---

## 6. Games (54 Tools)

> TUI games — chess, minesweeper, snake, tetris, sudoku, Wordle, typing tests, roguelikes, and more.

| Tool | Language | Description |
|------|----------|-------------|
| **chess-tui** | Rust | Play chess in terminal |
| **clidle** | Rust | Wordle in terminal (works over SSH!) |
| **NetHack** | C | Classic dungeon exploration |
| **BrogueCE** | C | Beautiful roguelike dungeon crawler |
| **pokete** | Python | Pokemon-like game |
| **DOOM-ASCII** | Python | Text-based DOOM |
| **balatrotui** | Rust | Balatro clone |
| **typeinc** | C | typing speed test |
| **thokr** | Rust | Sleek typing TUI |
| **nudoku** | C | ncurses sudoku |
| **games** | — | 54 total games in list |

---

## 7. Libraries

> **See dedicated reference**: `references/awesome-tuis-libraries.md`
>
> Covers 47+ TUI development libraries across 8 languages:
> - **Python**: 13 libs (Rich, textual, urwid, prompt-toolkit, blessed, etc.)
> - **Go**: 7 libs (Bubble Tea, tview, pterm, gocui, etc.)
> - **C**: 4 libs (ncurses, termbox2, libuv, etc.)
> - **C++**: 12 libs (FTXUI, imtui, tvision, FINAL CUT, etc.)
> - **Java**: 4 libs (Lanterna, Jexer, TUI4J, casciian)
> - **.NET**: 6 libs (Spectre.Console, Terminal.Gui, Consolonia, etc.)
> - **Rust**: 5 libs (Ratatui, iocraft, Zaz, tui-input, tui-rs)
> - **Other**: ink (Node.js), OpenTUI (TypeScript), gum (shell), etc.

---

## 8. Messaging (40 Tools)

### Email Clients
| Tool | Language | Description |
|------|----------|-------------|
| **aerc** | Go | Email client |
| **alpine** | C | Email client |
| **meli** | Rust | Email client |
| **Mutt** | C | Classic email client |
| **sup** | Ruby | Curses threads-with-tags email client |
| **matcha** | Rust | Email client |

### Chat/Messaging Clients
| Tool | Language | Description |
|------|----------|-------------|
| **irssi** | C | IRC terminal client |
| **Weechat** | C | Extensible chat client |
| **Profanity** | C | XMPP (Jabber) client |
| **mcabber** | C | XMPP client |
| **discordo** | Go | Discord terminal client |
| **endcord** | Rust | Discord TUI client |
| **concord** | Rust | Discord TUI client |
| **Slack-term** | Go | Slack client |
| **sclack** | Python | Slack terminal client |
| **matterhorn** | Haskell | Mattermost terminal client |
| **nchat** | C | Telegram/WhatsApp client |
| **tgt** | Rust | Telegram TUI |
| **gomuks** | Go | Matrix client |
| **iamb** | Rust | Matrix client for Vim addicts |
| **gurk-rs** | Rust | Signal Messenger client |
| **scli** | Python | Signal TUI |
| **siggo** | Go | Terminal UI for signal-cli |
| **toot** | Python | Mastodon CLI & TUI |
| **tut** | Rust | Mastodon TUI client |
| **mastui** | Rust | Mastodon TUI |
| **tuisky** | Rust | BlueSky TUI client |
| **nostui** | Rust | Nostr client |
| **nostratui** | Rust | Nostr post browser |
| **Devzat** | Go | Chat over SSH |
| **SuperChat** | Rust | Terminal-based threaded chat |
| **twitch-tui** | Rust | Twitch chat |
| **zulip-terminal** | Python | Official Zulip terminal client |
| **basalt** | Rust | Obsidian vault manager |
| **instagram-cli** | Node.js | Instagram from terminal |

---

## 9. Miscellaneous (70+ Tools)

### File & System Tools
| Tool | Language | Description |
|------|----------|-------------|
| **fzf** | Go | General-purpose fuzzy finder |
| **television** | Rust | Fast fuzzy finder TUI |
| **xplr** | Rust | Hackable TUI file explorer |
| **ncdu** | C | Disk usage analyzer (ncurses) |
| **gdu** | Go | Fast disk usage analyzer |
| **diskonaut** | Rust | Terminal disk space navigator |
| **cfdisk** | C | Partition editor |
| **cgdisk** | C++ | GPT partition editor |
| **npm** | — | recursive fuzzy finder |

### Bluetooth & WiFi
| Tool | Language | Description |
|------|----------|-------------|
| **bluetuith** | Go | Bluetooth connection manager |
| **bluetui** | Rust | Bluetooth device management |
| **impala** | Rust | WiFi management TUI |
| **nmtui** | C | ncurses network manager |
| **wavemon** | C | Wireless device monitoring |
| **wifitui** | Go | WiFi terminal UI |

### Security & Encryption
| Tool | Language | Description |
|------|----------|-------------|
| **gpg-tui** | Rust | TUI for GnuPG |
| **keydex** | Rust | KeePass password manager |
| **pass-cli** | Go | Password manager with rclone sync |
| **flawz** | Rust | Security vulnerability (CVE) browser |
| **tlock** | Rust | 2FA tokens manager |
| **vortix** | Rust | WireGuard/OpenVPN TUI |
| **WG Commander** | Bash | WireGuard VPN setup TUI |

### Visualization & Graphics
| Tool | Language | Description |
|------|----------|-------------|
| **mapscii** | Node.js | Braille & ASCII world map |
| **cava** | C | Cross-platform audio visualizer |
| **gif-for-cli** | Go | Convert GIF to ASCII |
| **asciiMol** | Python | ASCII molecule viewer |
| **physics-TUI** | Rust | Physics for undergraduate study |

### Logs & Monitoring
| Tool | Language | Description |
|------|----------|-------------|
| **lnav** | C++ | Advanced log file viewer |
| **termshark** | Go | Terminal UI for tshark |
| **oha** | Rust | HTTP load generator |

### Productivity Misc
| Tool | Language | Description |
|------|----------|-------------|
| **wego** | Go | Weather app |
| **wttr.in** | Python | Weather info |
| **jrnl** | Python | Journal/notes CLI |
| **arttime** | Python | Text-art clock and timer |
| **Caligula** | Rust | Disk imaging TUI |
| **csvlens** | Rust | CSV file viewer (like less for CSV) |
| **digisurf** | Rust | Signal waveform viewer |
| **golazo** | Go | Soccer match updates |
| **mqttui** | Rust | MQTT client |
| **nemu** | C | QEMU TUI |
| **neoss** | Rust | Socket statistics visualization |
| **tcpterm** | Go | Packet visualizer |

---

## 10. Multimedia (60+ Tools)

### Music Players
| Tool | Language | Description |
|------|----------|-------------|
| **cmus** | C | Small, fast console music player |
| **spotify-player** | Rust | Full Spotify player in terminal |
| **ncspot** | Rust | Cross-platform ncurses Spotify client |
| **spotui** | Python | Spotify client |
| **termusic** | Rust | Music player TUI |
| **kew** | Go | Terminal music player for Linux |
| **moc** | C | Console audio player |
| **mpvc** | Bash | mpc-like control for mpv |

### Video & Image
| Tool | Language | Description |
|------|----------|-------------|
| **chafa** | C | Image to ANSI/Unicode art converter |
| **timg** | C++ | Terminal image viewer |
| **viu** | Rust | Terminal image viewer |
| **vv** | C | Terminal image viewer |
| **tortuise** | Rust | 3D Gaussian splatting viewer |
| **Trophy** | Go | 3D model viewer (OBJ/GLB) |

### YouTube & Streaming
| Tool | Language | Description |
|------|----------|-------------|
| **mps-youtube** | Python | YouTube player/downloader |
| **pipe-viewer** | Perl | Lightweight YouTube client |
| **ytfzf** | Shell | YouTube finder/player |
| **terminal-yt** | Rust | Terminal YouTube manager |
| **GopherTube** | Go | YouTube client with mpv |
| **invidtui** | Go | Invidious client |
| **xytz** | Rust | YouTube video downloader |
| **ytui-music** | Rust | YouTube music player |

### Readers & Browsing
| Tool | Language | Description |
|------|----------|-------------|
| **bookokrat** | Rust | EPUB reader with Vim keybindings |
| **tdf** | Rust | PDF viewer |
| **fancy-cat** | Go | Terminal PDF reader |
| **Gorae** | Go | PDF/EPUB librarian |
| **manga-tui** | Rust | Manga reader |

### Editors & Drawing
| Tool | Language | Description |
|------|----------|-------------|
| **draw** | Go | Simple drawing tool |
| **textual-paint** | Python | MS Paint in terminal |
| **cmdpxl** | Rust | Command-line image editor |
| **favicon-editor** | Go | Grayscale favicon editor |

---

## 11. Productivity (80+ Tools)

### Terminal Multiplexers & Workspaces
| Tool | Language | Description |
|------|----------|-------------|
| **tmux** | C | Terminal multiplexer |
| **zellij** | Rust | Terminal workspace with batteries included |
| **dvtm** | C | dwm-like terminal multiplexer |
| **openmux** | TypeScript | Master-stack layout multiplexer |
| **TUIOS** | Python | TUI window manager |

### Note-Taking & Documentation
| Tool | Language | Description |
|------|----------|-------------|
| **Glow** | Go | Markdown reader |
| **patat** | Haskell | Terminal-based presentations |
| **presenterm** | Rust | Markdown terminal slideshow |
| **slides** | Go | Terminal-based presentations |
| **tui-slides** | Rust | Terminal presentations with images |
| **h-m-m** | Python | Hacker's mind map |
| **Toney** | Go | Lightweight note-taking |
| **tuidict** | Rust | Offline dictionary |

### Task & Project Management
| Tool | Language | Description |
|------|----------|-------------|
| **taskwarrior-tui** | Rust | TUI for Taskwarrior |
| **taskline** | Go | Tasks, boards & notes |
| **tododo** | Rust | TODO.md manager |
| **topydo** | Python | todo.txt-based task manager |
| **todoman** | Python | Simple standards-based todo |
| **kabmat** | Rust | Kanban board with vim keybindings |
| **kanban** | Go | Kanban with sprint tracking |
| **tiki** | Rust | Issue manager for git |
| **Judo** | Rust | Multi-database ToDo TUI |

### Calendars & Time Tracking
| Tool | Language | Description |
|------|----------|-------------|
| **calcurse** | C | Calendar and scheduling |
| **calcure** | Python | Modern TUI calendar |
| **khal** | Python | CalDAV-syncable calendar |
| **Chronos** | Rust | Vimlike calendar |
| **GeekCalendar** | Rust | Calendar with vim keys |
| **zeit** | Go | Time tracking |
| **Tock** | Rust | Time tracking TUI |
| **pomo** | Rust | Pomodoro timer |
| **helm** | Go | Minimalist pomodoro |

### SSH & Remote Access
| Tool | Language | Description |
|------|----------|-------------|
| **termscp** | Rust | File transfer (SCP/SFTP/FTP/S3) |
| **LazySSH** | Go | SSH manager from config files |
| **lssh** | Go | Remote access suite (SSH workflows) |
| **sshm** | Rust | SSH made easy |
| **ttm** | Go | SSH bookmark manager |

### Spreadsheet & Data
| Tool | Language | Description |
|------|----------|-------------|
| **VisiData** | Python | Terminal spreadsheet multitool |
| **sc-im** | C | ncurses spreadsheet program |
| **SheetsUI** | Rust | Console spreadsheet |
| **levite** | TypeScript | TUI spreadsheet with RPN |
| **Tabiew** | Rust | View/query CSV, TSV, parquet |
| **nless** | Rust | Tabular data pager |

### Jira & Project Management
| Tool | Language | Description |
|------|----------|-------------|
| **fjira** | Go | Jira TUI |
| **jiratui** | Rust | Jira TUI |
| **mynav** | Rust | Workspace and session management |
| **tuihub** | Rust | Utility hub/dashboard |
| **pagerduty-tui** | Go | PagerDuty incident management |

### Other Productivity
| Tool | Language | Description |
|------|----------|-------------|
| **mcfly** | Rust | Shell history search engine |
| **intelli-shell** | Python | Command templates with AI |
| **procmux** | Go | Run multiple commands in parallel |
| **agent-deck** | Go | Dashboard for AI coding agents |

---

## 12. Screensavers (6 Tools)

| Tool | Language | Description |
|------|----------|-------------|
| **astroterm** | Rust | Planetarium in your terminal |
| **gitlogue** | Rust | Visualize Git commit history |
| **neo** | C | Matrix digital rain |
| **rxpipes** | Rust | 2D pipes screensaver |
| **pond** | Rust | Soothing pond idle screen |
| **weathr** | Rust | Weather with ASCII animations |

---

## 13. Web (30 Tools)

### Browsers
| Tool | Language | Description |
|------|----------|-------------|
| **browsh** | Go | Modern text-based browser |
| **carbonyl** | Rust | Chromium running in terminal |
| **LYNX** | C | Classic text browser |
| **w3m** | C | Text-mode WWW browser |
| **elinks** | C | HTTP/FTP browser with JS |
| **Chawan** | Nim | TUI browser with CSS, images, JS |
| **bombadillo** | Go | Gopher/Gemini/Finger browser |
| **Lagrange** | C | Gemini browser |

### Feed Readers
| Tool | Language | Description |
|------|----------|-------------|
| **newsboat** | C++ | RSS/Atom feed reader |
| **bulletty** | Rust | Feed reader storing in Markdown |
| **eilmeldung** | Rust | RSS reader with bulk operations |
| **castero** | Python | Podcast app |
| **podliner** | Rust | Podcast client |

### Download Managers
| Tool | Language | Description |
|------|----------|-------------|
| **surge** | Go | TUI download manager |
| **nyaa** | Rust | Nyaa.si torrent browser |
| **rtorrent** | C++ | BitTorrent client |

### Social & News
| Tool | Language | Description |
|------|----------|-------------|
| **hackernews-TUI** | Rust | Hacker News browser |
| **haxor-news** | Python | Hacker News CLI |
| **rttt** | C++ | HN, RSS and Reddit reader |
| **tblogs** | Go | Development blog reader |
| **omaro** | Rust | Lobste.rs browser |
| **twterm** | Ruby | Twitter client |

### Specialized
| Tool | Language | Description |
|------|----------|-------------|
| **searxngr** | Rust | SearXNG search TUI |
| **rfc_reader** | Rust | RFC document browser |
| **CatenaVetus** | Rust | Church Fathers reader |
| **stegodon** | Go | ActivityPub microblog |
| **textual-web** | Python | Run TUIs in browser |
| **cloudflare-speed-cli** | Rust | Internet speed test |
| **Canard** | Go | Journalist RSS aggregator client |

---

## Key TUI Ecosystem Insights

### Dominant Languages
| Language | Count | Key Strengths |
|----------|-------|---------------|
| **Go** | 120+ | Infrastructure TUIs, Docker/K8s tools, simplicity |
| **Rust** | 100+ | Performance tools, file managers, modern replacements |
| **Python** | 60+ | Data analysis, prototyping, rapid development |
| **C/C++** | 60+ | Classic tools, editors, minimal footprint |
| **TypeScript/Node.js** | 20+ | AI tools, Ink-based apps |

### TUI Framework Popularity
| Framework | Language | Stars | Used By |
|-----------|----------|-------|---------|
| Ratatui | Rust | 12K+ | bottom, kdash, yazi |
| Bubble Tea | Go | 28K+ | lazydocker, charmbracelet tools |
| Textual | Python | 26K+ | harlequin, frogmouth |
| Ink | TypeScript | 33K+ | Claude Code, opencode |
| Blessed | Node.js | 11K+ | Various Node.js TUIs |
| FTXUI | C++ | 7K+ | C++ TUIs |
| tview | Go | 10K+ | Go TUIs |

---

## Cross-References

- **TUI Libraries (A-to-Z)**: `references/awesome-tuis-libraries.md`
- **Git TUIs**: `references/git-tui-tools.md` (lazygit, gitui, tig, gh, gh-dash)
- **TUI File Managers**: `references/tui-file-managers.md` (yazi, lf, ranger, nnn, broot)
- **Ink Ecosystem**: `references/ink-ecosystem.md`
- **Modern Unix Tools**: `references/modern-unix-tools.md`
