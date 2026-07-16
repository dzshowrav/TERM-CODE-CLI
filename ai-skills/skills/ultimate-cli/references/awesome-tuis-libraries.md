# Awesome TUIs — Library A-to-Z Reference

> Source: rothgar/awesome-tuis — 65+ TUI development libraries across 8 languages

A comprehensive reference for every TUI development library catalogued in the awesome-tuis list. Organized by language, with detailed entries and comparison tables for each category.

---

## Table of Contents

1. [Python (13 libraries)](#python)
2. [Go (7 libraries)](#go)
3. [C (4 libraries)](#c)
4. [C++ (12 libraries)](#c-1)
5. [Java (4 libraries)](#java)
6. [.NET (6 libraries)](#net)
7. [Rust (5 libraries)](#rust)
8. [Other Languages (8 libraries)](#other)

---

## Python

Python has the richest TUI library ecosystem with 13 libraries spanning from low-level curses wrappers to high-level web-inspired frameworks.

### 1. argent
**Language**: Python | **GitHub**: [software-mansion/argent](https://github.com/software-mansion/argent) | **Stars**: ~1K+ | **Status**: Active

**Description**: argent is an agentic toolkit for controlling, debugging, and profiling iOS and Android apps from the CLI. While primarily a mobile-dev tool, it includes a modular TUI application framework for building interactive terminal apps with command routing, autocompletion, and structured output formatting. Built on top of Rich and prompt-toolkit.

**Key Features**:
- Modular router/command architecture for building composable CLI apps
- Built on Rich for beautiful terminal output
- Autocompletion and interactive prompts
- System command orchestration
- Cross-platform mobile app interaction

**Use Case**: Building structured, modular CLI applications with autocomplete and command routing. Best for developer tooling and mobile app testing workflows.

---

### 2. blessed
**Language**: Python | **GitHub**: [jquast/blessed](https://github.com/jquast/blessed) | **Stars**: ~1K | **Status**: Active

**Description**: blessed is a thin, practical wrapper around Python's curses library that makes terminal programming easier. It provides a clean, Pythonic API for terminal I/O with proper Unicode support, keyboard input handling, colors, and terminal capabilities detection.

**Key Features**:
- Wraps curses with a Pythonic, easy-to-use API
- Full Unicode support with proper width calculation
- Location-aware text input (readline-like)
- Terminal color and style capabilities detection
- Cross-platform (Linux, macOS, Windows via windows-curses)
- Context managers for terminal state management

**Use Case**: Low-level terminal control when you need more than raw curses but less than a full framework. Great for apps that mix TUI and traditional CLI output.

---

### 3. blessings
**Language**: Python | **GitHub**: [erikrose/blessings](https://github.com/erikrose/blessings) | **Stars**: ~1.5K | **Status**: Maintained (low activity)

**Description**: blessings is a Python wrapper around ncurses that provides a cleaner, more Pythonic API. It was one of the first Python libraries to make terminal programming enjoyable, focusing on simplicity, ease of use, and beautiful output.

**Key Features**:
- Thin, clean wrapper around ncurses
- Simple color and style API
- Terminal width/height detection
- Context manager for terminal state
- Location-aware output (move cursor, clear, etc.)
- Minimal dependencies

**Use Case**: Simple terminal styling and cursor control when you want a lightweight alternative to the full curses API. Inspirational library for later tools like Rich.

---

### 4. notcurses
**Language**: C (with Python bindings) | **GitHub**: [dankamongmen/notcurses](https://github.com/dankamongmen/notcurses) | **Stars**: ~3.6K | **Status**: Active

**Description**: Notcurses is a blingful character graphics/TUI library for modern terminal emulators. It goes far beyond what ncurses can do, supporting vivid colors, multimedia, threads, Unicode, images, video, and complex visual effects. Python bindings are provided but the core is in C.

**Key Features**:
- True color, multimedia (images, video, GIFs) in terminal
- Thread-safe design with parallel rendering
- Unicode 14 support with full-width characters
- Visual effects: gradients, transparency, blending
- Planes-based compositing system (like layers in Photoshop)
- Direct-to-framebuffer rendering for maximum performance
- Mouse and keyboard input
- Multiple language bindings (Python, C++, Rust, Julia, etc.)

**Use Case**: When you need maximum visual bling in the terminal — video playback, image rendering, rich visualizations. Python bindings are best for non-performance-critical tasks; use the C API for performance.

---

### 5. py_cui
**Language**: Python | **GitHub**: [jwlodek/py_cui](https://github.com/jwlodek/py_cui) | **Stars**: ~791 | **Status**: Maintained

**Description**: py_cui is a widget-based TUI library inspired by gocui (Go). It uses a grid layout system (like Tkinter) where you define rows/columns and place widgets into grid cells. Comes with pre-built widgets including menus, textboxes, forms, scrollbars, and file explorers.

**Key Features**:
- Grid-based layout system (define rows/columns, place widgets)
- Pre-built widgets: buttons, labels, text boxes, scroll menus, checkboxes
- Popups: message boxes, menus, forms, file dialogs, prompts
- Custom keybindings for overview and focused modes
- Color rendering rules
- Popup stack system
- Cross-platform (Linux, macOS, Windows with windows-curses)

**Use Case**: Building grid-based TUI forms, menus, and file explorers quickly. Best when you want a familiar GUI-like layout system in the terminal.

---

### 6. pytermgui
**Language**: Python | **GitHub**: [bczsalba/pytermgui](https://github.com/bczsalba/pytermgui) | **Stars**: ~2.6K | **Status**: Active

**Description**: PyTermGUI is a powerful Python TUI framework with mouse support, a modular widget system, and a custom Terminal Markup Language (TIM) for styled text. It features a desktop-inspired window manager with modals and fully customizable windows.

**Key Features**:
- Desktop-inspired window manager with modals and stacking
- TIM (Terminal Markup Language) for inline text styling
- Mouse support out of the box (zero configuration)
- YAML/Python-based styling engines
- Full widget set: buttons, inputs, labels, containers, scrollable areas
- Cross-platform with NO_COLOR support that preserves usability
- JSON/string serialization of widgets
- Color wheel and automatic theme generation

**Use Case**: Building visually rich TUI applications with windows, modals, and complex layouts. Great for apps that need a desktop-like feel in the terminal.

---

### 7. python-prompt-toolkit
**Language**: Python | **GitHub**: [prompt-toolkit/python-prompt-toolkit](https://github.com/prompt-toolkit/python-prompt-toolkit) | **Stars**: ~9.4K | **Status**: Active

**Description**: python-prompt-toolkit is a library for building powerful interactive command-line applications. It's the engine behind popular tools like IPython, AWS CLI v2, and pgcli. It provides advanced input handling, syntax highlighting, auto-completion, and multi-line editing.

**Key Features**:
- Interactive input with syntax highlighting and auto-completion
- Multi-line text editing
- Vi and Emacs key binding modes
- Custom key bindings and input validation
- Layout system with split screens and floating panes
- Async support (asyncio, Trio)
- SSH server support for remote TUIs
- Formatted text output with HTML/ANSI support
- System clipboard integration

**Use Case**: Building complex interactive CLI prompts and REPL applications. The gold standard for Python CLI input handling.

---

### 8. pyTermTk
**Language**: Python | **GitHub**: [ceccopierangiolieugenio/pyTermTk](https://github.com/ceccopierangiolieugenio/pyTermTk) | **Stars**: ~897 | **Status**: Active

**Description**: pyTermTk (Python Terminal Toolkit) is a self-contained TUI library with a Qt-like API semantics. It uses a QT-like layout system (QLayout, QHBoxLayout, QVBoxLayout, QGridLayout) and the widget API closely mirrors Qt, GTK, and tkinter. It can even render to HTML5 for web deployment.

**Key Features**:
- Qt-like layout system (QLayout, HBox/VBox/Grid)
- Cross-compatible: Linux, macOS, Windows, HTML5
- True color support with full Unicode
- Built-in widget designer (ttkDesigner)
- Qt/GTK/tkinter-inspired API semantics
- Widget set: buttons, labels, inputs, tables, trees, scrollable containers
- HTML5 rendering for web deployment
- Full/half/zero-sized Unicode character support

**Use Case**: Python developers familiar with Qt who want to build TUIs with a similar mental model. Good for porting Qt apps to the terminal.

---

### 9. Rich
**Language**: Python | **GitHub**: [Textualize/rich](https://github.com/Textualize/rich) | **Stars**: ~50K | **Status**: Active

**Description**: Rich is the most popular Python terminal library for rich text and beautiful formatting. It adds colored text, tables, progress bars, markdown, syntax highlighting, tracebacks, and more to the terminal. Rich is both a standalone library and the rendering engine behind Textual.

**Key Features**:
- Rich text with colors, styles, and emoji
- Table rendering with auto-sizing and alignment
- Progress bars with multiple spinners and transients
- Markdown rendering
- Syntax highlighting (200+ languages)
- Beautiful tracebacks with rich diagnostics
- Tree views, panels, columns, and groups
- Layout system for terminal layout management
- Live display for dynamic content updates
- Minimal: works without a TUI event loop

**Use Case**: Enhancing CLI output with beautiful formatting — tables, progress bars, markdown, syntax-highlighted code. Use standalone for non-interactive CLI apps; use as the rendering layer for Textual TUIs.

---

### 10. textual
**Language**: Python | **GitHub**: [Textualize/textual](https://github.com/Textualize/textual) | **Stars**: ~27K | **Status**: Active

**Description**: Textual is a modern TUI framework inspired by web development. It uses CSS for styling, a reactive widget system, and an async event-driven architecture. It's built on top of Rich and provides a complete application framework for building sophisticated terminal user interfaces.

**Key Features**:
- CSS-based styling (familiar to web developers)
- Reactive widget system with message passing
- Mouse and keyboard input handling
- Built-in widgets: data table, tree, input, button, tabs, header/footer, listview
- Custom widget creation API
- Developer tools (inspector, CSS hot reload)
- Screen navigation and modal screens
- Async/await event loop
- Worker threads for background tasks
- Built-in testing framework
- Textual-web for running TUIs in the browser

**Use Case**: Building rich, interactive TUI applications with a web development workflow. The go-to framework for complex Python TUIs that need multiple screens, data tables, forms, and custom widgets.

---

### 11. UniCurses
**Language**: Python | **GitHub**: [bjshea/UniCurses](https://github.com/bjshea/UniCurses) | **Stars**: ~100 | **Status**: Maintained (low activity)

**Description**: UniCurses provides a unified curses API that works across all operating system platforms (Linux, macOS, Windows). It abstracts away platform-specific differences so that curses code written on one platform runs identically on others.

**Key Features**:
- Cross-platform curses compatibility layer
- Same API on Windows, macOS, and Linux
- No modification needed for cross-platform use
- Supports Python 2 and 3
- Wraps pdcurses on Windows, ncurses on Unix

**Use Case**: When you need curses-based code to work identically on Windows and Unix without platform-specific branches.

---

### 12. urwid
**Language**: Python | **GitHub**: [urwid/urwid](https://github.com/urwid/urwid) | **Stars**: ~2.8K | **Status**: Maintained

**Description**: urwid is a mature console user interface library for Python on Linux/Unix. It's known for its widget-based approach, flexible text layout, and support for both standard and advanced terminal features. One of the oldest Python TUI libraries still actively maintained.

**Key Features**:
- Widget-based architecture with composition
- Canvas-based rendering for flexible text layout
- Support for attributes, colors, and text alignment
- Event-driven main loop with async support
- Built-in widgets: edit, button, columns, piles, frames, overlays
- Advanced text layout (alignment, wrapping, clipping)
- Signal system for inter-widget communication
- Raw terminal mode for full control
- Extensive Unicode support

**Use Case**: Building traditional widget-based TUIs with a proven, stable library. Good for terminal email clients, chat clients, and text-heavy applications.

---

### 13. Vindauga
**Language**: Python | **GitHub**: [gabbpuy/vindauga](https://github.com/gabbpuy/vindauga) | **Stars**: ~100 | **Status**: Maintained

**Description**: Vindauga is a pure Python 3 implementation of the BSD-licensed C++ Turbo Vision library (the classic DOS-era application framework). It provides a cross-platform TUI with menus, dialogs, desktop windows, and Turbo Vision-style application structure.

**Key Features**:
- Pure Python implementation of Turbo Vision 2.0
- No dependencies beyond Python curses
- Menu bars, status lines, desktop windows
- Dialog boxes with buttons, inputs, listboxes, checkboxes, radio buttons
- Unicode by default (not CP437)
- Combo boxes and extra widgets beyond Turbo Vision
- Interactive shell windows
- Dynamic console window resizing
- Cross-platform (Mac, Windows, Linux, Cygwin, Putty, X-Terms)

**Use Case**: Building Turbo Vision-style desktop-like TUI applications. Best for developers nostalgic for the classic DOS Turbo Vision interface or porting Turbo Vision apps to Python.

---

### Python Libraries Comparison

| Library | Stars | Widgets | Layout | CSS/Styling | Mouse | Windows | Best For |
|---------|-------|---------|--------|-------------|-------|---------|----------|
| **Rich** | ~50K | Panels, tables, progress | Minimal | Programmatic | No | Yes | CLI output formatting |
| **textual** | ~27K | Full widget set | CSS + Dock | CSS-based | Yes | Yes | Complex TUI apps |
| **python-prompt-toolkit** | ~9.4K | Input-focused | Split/floating | HTML/ANSI | Yes | Yes | REPLs, prompts |
| **notcurses** | ~3.6K | Planes-based compositing | Manual/auto | Programmatic | Yes | Yes | Multimedia, visual bling |
| **urwid** | ~2.8K | Widget-based | Columns/rows | Canvas-based | Yes | No | Classic widget TUIs |
| **pytermgui** | ~2.6K | Full widget set | Window manager | YAML/Python | Yes | Yes | Desktop-like TUIs |
| **blessings** | ~1.5K | None | None | No | No | No | Simple terminal styling |
| **pyTermTk** | ~897 | Qt-like widgets | Qt layouts | Programmatic | Yes | Yes | Qt-style TUIs |
| **py_cui** | ~791 | Grid widgets | Grid layout | Color rules | No | Yes | Grid-based forms |
| **blessed** | ~1K | Input-focused | None | No | Yes | Yes | Terminal I/O wrapper |
| **argent** | ~1K | Command routers | Modular | Rich-based | No | Yes | Modular CLI apps |
| **UniCurses** | ~100 | None (raw curses) | None | No | No | Yes | Cross-platform curses |
| **Vindauga** | ~100 | Turbo Vision widgets | Desktop/windows | Programmatic | Yes | Yes | Turbo Vision apps |

---

## Go

Go has a strong TUI library ecosystem driven largely by Charmbracelet's Bubble Tea and the interactive infrastructure tooling community.

### 14. bubbletea
**Language**: Go | **GitHub**: [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) | **Stars**: ~28K+ | **Status**: Active

**Description**: Bubble Tea is the dominant Go TUI framework, based on The Elm Architecture. It provides a simple functional model where the UI is a function of state (Model -> View -> Update). Combined with the Charm ecosystem (Bubbles, Lip Gloss), it powers hundreds of terminal applications.

**Key Features**:
- Elm Architecture (model, update, view)
- Component-based composition with hierarchical models
- Keyboard and mouse input handling
- Window size change notifications
- Channel-based event system
- High-performance rendering (differential updates)
- Full integration with Lip Gloss (styling) and Bubbles (components)
- Huge ecosystem of pre-built components
- Extensive documentation and examples

**Use Case**: Building any interactive TUI in Go. The standard choice for new Go TUI projects — from simple prompts to complex multi-screen applications.

---

### 15. gocui
**Language**: Go | **GitHub**: [jroimartin/gocui](https://github.com/jroimartin/gocui) | **Stars**: ~10.5K | **Status**: Maintained

**Description**: gocui is a minimalist Go package for creating Console User Interfaces. It provides a simple API for managing terminal views, handling keyboard/mouse input, and building overlapping panel-based interfaces.

**Key Features**:
- View-based layout with overlapping panels
- Keyboard and mouse input handling
- Views can have borders, titles, and colors
- Custom keybinding management
- Scrollable views
- Minimal and easy-to-learn API
- Well-documented with good examples

**Use Case**: Simple overlapping-panel TUIs when you want minimal dependencies and a straightforward API. Good for building dashboards and multi-panel tools.

---

### 16. pterm
**Language**: Go | **GitHub**: [pterm/pterm](https://github.com/pterm/pterm) | **Stars**: ~5K | **Status**: Active

**Description**: pterm is a modern Go module for beautifying console output. It provides a wide range of pre-built components including charts, progress bars, tables, spinners, bullet lists, and styled text — all with a consistent, chainable API.

**Key Features**:
- Pre-built printers: text, header, table, tree, list, chart
- Progress bars with multiple styles and spinners
- Bar charts and area charts
- Tables with auto-sizing and alignment
- Multi-line live output
- Section and paragraph printers
- RGB and 256-color support
- Cross-platform with graceful fallbacks
- NO_COLOR support
- Chainable, fluent API

**Use Case**: Beautifying CLI output with styled text, tables, progress bars, and charts. Best for non-interactive CLI output enhancement (like Rich for Python, but in Go).

---

### 17. stickers
**Language**: Go | **GitHub**: [76creates/stickers](https://github.com/76creates/stickers) | **Stars**: ~382 | **Status**: Maintained

**Description**: stickers is a collection of reusable TUI building blocks for the Charmbracelet ecosystem. It provides FlexBox and Table components built for Bubble Tea using Lip Gloss for styling. Think of it as supplementary widgets for the Charm ecosystem.

**Key Features**:
- FlexBox layout component for Bubble Tea
- Table component with sorting and selection
- Built on Lip Gloss for consistent styling
- Designed to complement the Bubbles component library
- Lightweight and focused

**Use Case**: Adding FlexBox layouts and table widgets to Bubble Tea applications that need more layout control than Bubbles provides.

---

### 18. tui-go
**Language**: Go | **GitHub**: [marcusolsson/tui-go](https://github.com/marcusolsson/tui-go) | **Stars**: ~2.1K | **Status**: DEPRECATED (archived)

**Description**: tui-go was a Go UI library for terminal applications that provided a higher-level programming model for building layout-based user interfaces. It is now archived and unmaintained. The author recommends migrating to tview or Bubble Tea.

**Key Features**:
- Layout-based widget model
- Automatic resize handling
- Widget set: table, list, status bar, input
- Box layout with horizontal/vertical stacks
- Vim-like keybindings in some examples

**Use Case**: None — do not use for new projects. Migrate existing applications to tview or Bubble Tea.

---

### 19. tview
**Language**: Go | **GitHub**: [rivo/tview](https://github.com/rivo/tview) | **Stars**: ~10K+ | **Status**: Active

**Description**: tview is a Go TUI library with a rich set of interactive widgets including forms, tables, tree views, text views, list views, and input fields. It builds on top of tcell for terminal I/O and provides a complete application framework.

**Key Features**:
- Rich widget set: Form, Table, TreeView, TextView, List, InputField
- Flex layout system with rows and columns
- Keyboard shortcuts and navigation
- Mouse support
- Modal dialogs
- Primitive composition
- Color themes and customization
- Pagination for large data
- Focus management
- Extensible via custom primitives

**Use Case**: Building data-heavy TUI applications with tables, forms, tree navigation, and text views. Strong choice for database browsers, configuration tools, and monitoring dashboards.

---

### 20. tcell
**Language**: Go | **GitHub**: [gdamore/tcell](https://github.com/gdamore/tcell) | **Stars**: ~5.2K | **Status**: Active

**Description**: tcell is a Go terminal handling library that provides a low-level API for terminal I/O. It handles terminal characteristics, colors, mouse events, and screen rendering. tcell is the foundation for tview and many other Go TUI frameworks.

**Key Features**:
- Terminal screen creation and management
- True color (24-bit) support
- Mouse event handling (click, drag, release)
- Keyboard input with modifiers
- Unicode and wide character support
- Terminal resize events
- Style buffers for efficient rendering
- Character attributes: bold, underline, blink, reverse
- Concurrency-safe API
- Works on Windows, macOS, Linux

**Use Case**: Low-level terminal control in Go. Use directly when you need fine-grained control over terminal rendering. Most developers should use tview or Bubble Tea on top of tcell.

---

### Go Libraries Comparison

| Library | Stars | Architecture | Widgets | Layout | Styling | Status | Best For |
|---------|-------|-------------|---------|--------|---------|--------|----------|
| **bubbletea** | ~28K | Elm Architecture | Bubbles components | Flex/Layout | Lip Gloss | Active | General TUI apps |
| **gocui** | ~10.5K | View-based | Minimal views | Overlapping | Colors | Maintained | Simple panel TUIs |
| **tview** | ~10K | Primitive-based | Rich widget set | Flex | Themes | Active | Data-heavy TUIs |
| **tcell** | ~5.2K | Terminal I/O | None | None | Style buffers | Active | Low-level terminal control |
| **pterm** | ~5K | Printer-based | Charts, tables, bars | Auto | Chainable API | Active | CLI output beautification |
| **tui-go** | ~2.1K | Layout-based | Basic widgets | Auto | Colors | DEPRECATED | Do not use |
| **stickers** | ~382 | Component-based | FlexBox, Table | FlexBox | Lip Gloss | Maintained | Bubble Tea widgets |

---

## C

C libraries form the foundation of TUI development. ncurses is the grandparent of most TUI libraries, while newer libraries offer lighter alternatives.

### 21. AnbUI
**Language**: C | **GitHub**: [oerg866/anbui](https://github.com/oerg866/anbui) | **Stars**: ~50 | **Status**: Maintained

**Description**: AnbUI is a very minimal Text UI library in C designed with absolute ease of use in mind. It provides simple functions for creating text-based UI elements like boxes, buttons, and text areas with minimal ceremony.

**Key Features**:
- Minimal API (one or two function calls per element)
- Text boxes and frames
- Simple button and input support
- Tiny code footprint
- No external dependencies
- Single/small source file

**Use Case**: When you need a super minimal TUI in C with minimal code overhead. Good for embedded tools and simple terminal interfaces.

---

### 22. libuv
**Language**: C | **GitHub**: [libuv/libuv](https://github.com/libuv/libuv) | **Stars**: ~24K | **Status**: Active

**Description**: libuv is a cross-platform asynchronous I/O library, originally developed for Node.js. While not a TUI library itself, it provides the event loop foundation for many TUI frameworks and is commonly listed in TUI toolkits for its async I/O capabilities.

**Key Features**:
- Full-featured event loop (epoll, kqueue, IOCP, event ports)
- Asynchronous TCP and UDP sockets
- Asynchronous DNS resolution
- Async file and filesystem operations
- File system events
- ANSI TTY control
- IPC with socket sharing
- Thread pool
- Signal handling
- High-resolution clock

**Use Case**: Providing async I/O and event loop infrastructure for TUI applications that need network, file I/O, and timer management alongside their UI. Often used alongside or instead of ncurses.

---

### 23. ncurses
**Language**: C | **GitHub**: [mirror/ncurses](https://github.com/mirror/ncurses) (mirror) | **Stars**: N/A (foundational) | **Status**: Active (downstream via Thomas Dickey)

**Description**: ncurses (new curses) is the classic C library for building text-based user interfaces. It provides a terminal-independent API for drawing panels, menus, forms, and handling input. It underpins thousands of classic TUI applications (htop, mc, irssi, etc.) and has bindings for virtually every language.

**Key Features**:
- Terminal-independent rendering via terminfo database
- Window, panel, and subwindow management
- Menu and form libraries (panel.h, menu.h, form.h)
- Color pair management (up to 256 colors)
- Mouse event handling (GPM, xterm)
- Character attributes (bold, reverse, underline, blink)
- Soft label keys
- Pad windows for scrolling content
- Keypad mode for function keys
- Extensively documented and battle-tested

**Use Case**: The foundational C TUI library. Use when you need maximum portability and compatibility, or when building libraries that need bindings to many languages. For new projects, consider termbox2 or notcurses for a modern alternative.

---

### 24. tuibox
**Language**: C | **GitHub**: [Cubified/tuibox](https://github.com/Cubified/tuibox) | **Stars**: ~100 | **Status**: Maintained (low activity)

**Description**: tuibox ("toybox") is a single-header terminal UI library capable of creating mouse-driven, interactive applications on the command line. It has no external dependencies and uses pure ANSI escape sequences for rendering.

**Key Features**:
- Single-header library (one .h file)
- No external dependencies (no ncurses required)
- Pure ANSI escape sequence rendering
- Mouse-driven interactive applications
- Keyboard input handling
- Container and widget layout
- Simple API
- Cross-platform

**Use Case**: Building mouse-driven TUI apps in C with minimal dependencies. Good for simple interactive tools where ncurses is overkill.

---

### 24b. termbox2
**Language**: C | **GitHub**: [termbox/termbox2](https://github.com/termbox/termbox2) | **Stars**: ~720 | **Status**: Active

**Description**: termbox2 is a slim, modern alternative to ncurses for terminal I/O. It provides a tighter, cleaner API focused on minimalism and simplicity while supporting modern terminal features like true color and mouse input.

**Key Features**:
- Clean, minimal API (much smaller than ncurses)
- True color (24-bit) support
- Mouse event handling
- UTF-8 and Unicode support
- No external dependencies
- Simplified terminal initialization
- Cell-based buffer rendering
- Cross-platform

**Use Case**: When you need a lightweight, modern replacement for ncurses with a cleaner API. Good for new C TUI projects that don't need the full ncurses feature set.

---

### 24c. TermGL
**Language**: C | **GitHub**: [wojciech-graj/TermGL](https://github.com/wojciech-graj/TermGL) | **Stars**: ~500 | **Status**: Active

**Description**: TermGL is a terminal-based 2D/3D graphics library written in C. It provides vertex and fragment shaders, matrix transformations, camera control, and more — all rendered to the terminal using ANSI escape codes.

**Key Features**:
- 2D and 3D rendering in the terminal
- Custom vertex and fragment shaders
- Z-buffering and back-face culling
- Matrix transforms (translation, rotation, scaling)
- Camera with FOV control
- Triangle rasterization
- Color gradients
- No external dependencies
- C99 compliant

**Use Case**: Rendering 2D and 3D graphics directly in the terminal. Suitable for demos, visualizations, and terminal-based games with 3D elements.

---

### C Libraries Comparison

| Library | Stars | Dependencies | API Style | Rendering | Mouse | Best For |
|---------|-------|-------------|-----------|-----------|-------|----------|
| **libuv** | ~24K | None | Async I/O | N/A | N/A | Event loop for TUIs |
| **ncurses** | ~5K+ | terminfo | Traditional C | Windows/panels | Yes | Foundational TUI work |
| **termbox2** | ~720 | None | Minimal C | Cell buffer | Yes | Modern ncurses alternative |
| **TermGL** | ~500 | None | GL-like | Frame buffer | No | 2D/3D terminal graphics |
| **tuibox** | ~100 | None | Single-header | ANSI boxes | Yes | Mouse-driven TUIs |
| **AnbUI** | ~50 | None | Minimal | Text boxes | No | Ultra-minimal TUIs |

---

## C++

C++ offers a wide range of TUI libraries from game engines to functional UI frameworks, with FTXUI leading the modern C++ TUI renaissance.

### 25. ASCII_Board_Game_Engine
**Language**: C++ | **GitHub**: [tjunruh/ASCII_Board_Game_Engine](https://github.com/tjunruh/ASCII_Board_Game_Engine) | **Stars**: ~50 | **Status**: Maintained

**Description**: A C++ graphics engine designed for creating 2D board games using ASCII characters. Provides rendering, input, and game loop management for text-based board games.

**Key Features**:
- ASCII character-based 2D rendering
- Board/cell manipulation
- Game loop management
- Color support
- Input handling
- Tile and sprite management

**Use Case**: Building ASCII-based board games (chess, checkers, roguelikes) with a straightforward game engine in C++.

---

### 26. ConsoleCraftEngine
**Language**: C++ | **GitHub**: [ural89/ConsoleCraftEngine](https://github.com/ural89/ConsoleCraftEngine) | **Stars**: ~50 | **Status**: Maintained

**Description**: ConsoleCraftEngine is a terminal-based 2D game engine written in C++. It features Box2D physics integration and works on both Linux and Windows terminals.

**Key Features**:
- Box2D physics engine integration
- Terminal-based 2D rendering
- Cross-platform (Linux, Windows)
- Sprite and entity management
- Game loop and event handling
- Tile-based world rendering

**Use Case**: Building 2D games with physics simulation that run entirely in the terminal.

---

### 27. FINAL CUT
**Language**: C++ | **GitHub**: [ganssle/FINAL-CUT](https://github.com/ganssle/FINAL-CUT) | **Stars**: ~520 | **Status**: Active

**Description**: FINAL CUT is a C++ class library for creating a text-based user interface. It provides a full set of widget classes (dialog boxes, buttons, list boxes, etc.) and supports mouse and keyboard interaction on any terminal that supports ANSI escape sequences.

**Key Features**:
- Object-oriented C++ widget library
- Dialog boxes with buttons, inputs, list boxes, checkboxes
- Menu bars and popup menus
- Status bars
- Mouse and keyboard input
- True color support
- Unicode/UTF-8 support
- Application and event loop framework
- Focus management
- Modal and modeless dialogs

**Use Case**: Building traditional desktop-style TUI applications in C++ with a widget-based approach. Good for data entry forms and utility dialogs.

---

### 28. FTXUI
**Language**: C++ | **GitHub**: [ArthurSonzogni/FTXUI](https://github.com/ArthurSonzogni/FTXUI) | **Stars**: ~7K+ | **Status**: Active

**Description**: FTXUI is the leading C++ functional terminal user interface library. It follows a declarative, functional programming style where the UI is composed of elements that are combined to form complex layouts. It supports DOM-like element composition.

**Key Features**:
- Functional/declarative API (compose elements like React)
- Flexbox layout (similar to CSS)
- Screen and DOM-based rendering
- Rich widget set: button, checkbox, input, menu, tabs, table
- Canvas for custom drawing
- Mouse and keyboard input
- Event loop and timers
- Cross-platform (Linux, macOS, Windows, WebAssembly)
- True color and Unicode support
- Animation support

**Use Case**: The standard choice for modern C++ TUI development. Best for developers who want a functional, composable API similar to React/Elm.

---

### 29. GGUI
**Language**: C++ | **GitHub**: [Gabidal/GGUI](https://github.com/Gabidal/GGUI) | **Stars**: ~9 | **Status**: Maintained

**Description**: GGUI is a lightweight C++17 structured terminal user interface library with a single-header design. It provides dynamic containers, text fields, progress bars, buttons, canvas drawing, sprite animation, and event handling.

**Key Features**:
- Single header + library file
- Dynamic containers (horizontal and vertical lists)
- Canvas with built-in sprite animation
- Buttons and switches
- Progress bars with multi-line support
- Text fields with alignment
- Mouse support with hover and focus
- Transparency control
- Custom event handling
- Cross-platform (Windows, Linux, Android)

**Use Case**: Lightweight C++ TUIs needing sprite animation and game-like UIs. Good for real-time monitoring and simple games.

---

### 30. imtui
**Language**: C++ | **GitHub**: [ggerganov/imtui](https://github.com/ggerganov/imtui) | **Stars**: ~3K | **Status**: Maintained

**Description**: imtui (Immediate Mode Text-based UI) is an immediate mode GUI library for the terminal, based on Dear ImGui. It supports 256 ANSI colors, mouse/keyboard input, and provides the full ImGui API rendered to ncurses.

**Key Features**:
- Immediate mode GUI (like Dear ImGui)
- Full Dear ImGui API coverage in terminal
- 256 ANSI colors
- Mouse and keyboard input
- Windows, menus, widgets
- NCurses backend for rendering
- WebAssembly support via Emscripten
- Same programming model as Dear ImGui

**Use Case**: When you want Dear ImGui's immediate mode paradigm in the terminal. Good for debugging tools, profilers, and interactive visualizations.

---

### 31. rang
**Language**: C++ | **GitHub**: [agauniyal/rang](https://github.com/agauniyal/rang) | **Stars**: ~2.4K | **Status**: Maintained

**Description**: rang is a minimal, header-only modern C++ library for terminal output styling. It provides colors, styles (bold, dim, underline, blink, reverse), and OS-specific terminal handling through a simple iostream wrapper.

**Key Features**:
- Single header-only library
- No external dependencies
- Works with cout, clog, cerr
- Colors (foreground, background, 256-color, true color)
- Text styles: bold, dim, italic, underline, blink, reverse, concealed, crossed
- Automatic TTY detection (no ANSI codes when output is piped)
- Cross-platform (Windows, Unix)
- Modern C++11 API

**Use Case**: Adding colors and styles to C++ terminal output. Best for formatting CLI output without the overhead of a full TUI library.

---

### 32. termdb
**Language**: C++ | **GitHub**: [agauniyal/termdb](https://github.com/agauniyal/termdb) | **Stars**: ~50 | **Status**: Maintained

**Description**: termdb is a single-header C++ library for parsing terminfo databases. It allows querying terminal capabilities programmatically, supporting boolean, string, and numeric capabilities from the system terminfo database.

**Key Features**:
- Single-header library
- Parses system terminfo databases
- Boolean, string, and numeric capability queries
- Support for any terminal in the terminfo database
- Clean C++ API
- Error handling for missing capabilities

**Use Case**: When your C++ TUI library needs to query terminal capabilities (color support, key sequences, cursor movement) from the system terminfo database.

---

### 33. Tui Widgets
**Language**: C++ | **GitHub**: [tuiwidgets/tuiwidgets](https://github.com/tuiwidgets/tuiwidgets) | **Stars**: ~200 | **Status**: Active

**Description**: Tui Widgets is a high-level widget toolkit for terminal applications. It provides a comprehensive set of desktop-style widgets including buttons, text inputs, tables, trees, tabs, collapsible panels, and more.

**Key Features**:
- High-level widget set (similar to Qt/GTK)
- Buttons, labels, inputs, checkboxes, radio buttons
- Tables, tree views, list views
- Tab widgets and collapsible panels
- Progress bars and sliders
- Splitter panels (resizable)
- Composite widgets and custom widgets
- True color and Unicode
- Mouse and keyboard input

**Use Case**: Building desktop-style TUI applications in C++ with a comprehensive widget set. Good when you need many widget types out of the box.

---

### 34. tvision
**Language**: C++ | **GitHub**: [magiblot/tvision](https://github.com/magiblot/tvision) | **Stars**: ~1.1K | **Status**: Active

**Description**: tvision is a modern port of Borland's Turbo Vision 2.0 to C++. It preserves the classic Turbo Vision API while adding modern C++ features, true color support, Unicode, and cross-platform terminal support.

**Key Features**:
- Modern port of Turbo Vision 2.0
- Classic Turbo Vision API
- True color and Unicode support
- Event-driven application framework
- Desktop, windows, dialogs, menus
- Overlapping window management
- Controls: buttons, inputs, listboxes, checkboxes
- Modern C++ (C++17/20)
- Cross-platform (Linux, macOS, Windows)

**Use Case**: Building Turbo Vision-style applications with modern C++. Best for porting legacy Turbo Vision code or building apps with that classic desktop-in-terminal feel.

---

### 35. uvw
**Language**: C++ | **GitHub**: [skypjack/uvw](https://github.com/skypjack/uvw) | **Stars**: ~2.5K | **Status**: Active

**Description**: uvw is a header-only, event-based C++ wrapper around libuv. It provides a modern C++ API for libuv's async I/O capabilities, making it useful as an event loop foundation for C++ TUI applications.

**Key Features**:
- Header-only C++17 wrapper around libuv
- Event loop, timers, and async operations
- TCP, UDP, TTY, pipe, and filesystem operations
- Signals and child process management
- Modern C++ API with RAII
- Type-safe and non-intrusive
- Same performance as libuv

**Use Case**: Providing an async event loop for C++ TUI applications that need network I/O, timers, and filesystem operations alongside terminal rendering.

---

### 36. xtd
**Language**: C++ | **GitHub**: [gammasoft71/xtd](https://github.com/gammasoft71/xtd) | **Stars**: ~1.1K | **Status**: Active

**Description**: xtd is a modern C++ framework for creating console (TUI) and GUI applications, unit tests, and more. It provides a comprehensive API inspired by .NET (WinForms-like) but for C++, with native support for terminal rendering.

**Key Features**:
- Console forms (TUI forms like Windows Forms)
- GUI and TUI in a single framework
- Widgets: forms, buttons, labels, text boxes, list boxes
- Event-driven programming model
- .NET-inspired API
- Drawing and graphics support
- Cross-platform (Linux, macOS, Windows)
- Built-in unit testing framework
- String and collection utilities

**Use Case**: Building cross-platform C++ applications that can run as both TUI and GUI with a .NET-like programming model. Good for porting .NET desktop apps to C++.

---

### C++ Libraries Comparison

| Library | Stars | Paradigm | Widget Set | Header-Only | C++ Standard | Best For |
|---------|-------|----------|------------|-------------|-------------|----------|
| **FTXUI** | ~7K | Functional/DOM | Rich | No | C++17 | General C++ TUI (recommended) |
| **imtui** | ~3K | Immediate mode | Full (ImGui) | No | C++17 | Dear ImGui-style TUIs |
| **uvw** | ~2.5K | Event loop | None | Yes | C++17 | Async I/O event loop |
| **rang** | ~2.4K | Stream wrapper | None | Yes | C++11 | Terminal color/styling |
| **tvision** | ~1.1K | Turbo Vision | Rich | No | C++17 | Turbo Vision apps |
| **xtd** | ~1.1K | .NET Forms-like | Full widget set | No | C++17 | .NET-style apps |
| **FINAL CUT** | ~520 | OOP widgets | Full widget set | No | C++17 | Desktop-style TUI |
| **Tui Widgets** | ~200 | Widget toolkit | Comprehensive | No | C++17 | Rich desktop widgets |
| **termdb** | ~50 | Terminfo parser | None | Yes | C++11 | Terminfo querying |
| **ASCII_Board_Game** | ~50 | Game engine | Game-specific | No | C++ | ASCII board games |
| **ConsoleCraftEngine** | ~50 | Game engine | Physics/game | No | C++ | 2D physics games |
| **GGUI** | ~9 | Structured | Basic widgets | Yes | C++17 | Lightweight/game TUIs |

---

## Java

Java's TUI ecosystem includes both classic libraries (Lanterna) and modern frameworks with ports of popular Go/Rust libraries.

### 37. casciian
**Language**: Java | **GitHub**: [crramirez/casciian](https://github.com/crramirez/casciian) | **Stars**: ~39 | **Status**: Active

**Description**: casciian is a Text User Interface library for Java that deliberately avoids dependencies on the JDK's java.desktop module, enabling compilation with GraalVM native-image for AOT-compiled native executables with fast startup.

**Key Features**:
- No java.desktop dependency (GraalVM native-image compatible)
- AOT compilation support
- Text-based UI components
- Keyboard input handling
- Clean API
- Cross-platform

**Use Case**: Building Java TUIs that need GraalVM native-image compilation for fast startup and small footprint. Good for CLI tools and lightweight TUI apps.

---

### 38. Jexer
**Language**: Java | **GitHub**: [gitblit/Jexer](https://github.com/gitblit/Jexer) | **Stars**: ~400 | **Status**: Active

**Description**: Jexer is a Java text-based windowing system inspired by Turbo Vision. It provides a desktop-like environment in the terminal with overlapping windows, dialogs, menus, status bars, and a full set of UI controls.

**Key Features**:
- Turbo Vision-like windowing system
- Overlapping, resizable, movable windows
- Desktop, dialogs, menus, status bars
- Rich widget set: buttons, text, tables, trees
- True color support
- Mouse and keyboard input
- Backend rendering via ECMA-48 / ANSI X3.64
- Semicolon-separated widget layout definitions
- UTF-8 support

**Use Case**: Building desktop-like TUI applications in Java with a classic Turbo Vision-inspired interface. Good for terminal-based IDEs, file managers, and data browsers.

---

### 39. Lanterna
**Language**: Java | **GitHub**: [mabe02/lanterna](https://github.com/mabe02/lanterna) | **Stars**: ~2.2K | **Status**: Active

**Description**: Lanterna is a mature Java library for creating text-based terminal UIs. It provides a curses-like API with support for screens, windows, panels, and widgets. It's the most established Java TUI library.

**Key Features**:
- Terminal screen abstraction (Swing-like)
- Window and panel system
- Widgets: buttons, labels, text boxes, list boxes, tables
- Dialog boxes and message boxes
- True color support
- Mouse and keyboard input
- Unicode/UTF-8 support
- Swing terminal emulator for testing
- Screen buffering and double-buffering
- Cross-platform

**Use Case**: The go-to Java TUI library for building any kind of terminal application with widgets and windows. Battle-tested and feature-rich.

---

### 40. TUI4J
**Language**: Java | **GitHub**: [williamagh/tui4j](https://github.com/williamagh/tui4j) | **Stars**: ~70 | **Status**: Active (new)

**Description**: TUI4J (Terminal User Interface for Java) is a modern Java TUI framework inspired by Bubble Tea (Go) with elements from Textual (Python). It ports the Charmbracelet ecosystem (Bubble Tea, Bubbles, Lip Gloss) to Java with a 1:1 API mapping.

**Key Features**:
- Bubble Tea port to Java (Elm Architecture)
- Lip Gloss styling (colors, borders, layout)
- Bubbles components (viewport, textarea, table, progress)
- ANSI parsing
- Spring physics animation (Harmonica)
- Additive extensions beyond Charm API
- Compatible with GraalVM native-image
- Modern Java (17+)

**Use Case**: Building Java TUIs with the Elm Architecture pattern popularized by Bubble Tea. Best for Java developers who want a functional, component-based TUI framework.

---

### Java Libraries Comparison

| Library | Stars | Architecture | Widget Set | GraalVM | Best For |
|---------|-------|-------------|------------|---------|----------|
| **Lanterna** | ~2.2K | Traditional/Swing-like | Rich | Limited | General Java TUI |
| **Jexer** | ~400 | Turbo Vision | Rich | Yes | Desktop-like TUIs |
| **TUI4J** | ~70 | Elm (Bubble Tea) | Growing | Yes | Functional/component TUIs |
| **casciian** | ~39 | Minimal | Basic | Yes | Native-image TUIs |

---

## .NET

.NET has a robust TUI ecosystem with both mature toolkits (Terminal.Gui, Spectre.Console) and innovative newcomers (Hex1b, Consolonia).

### 41. Consolonia
**Language**: C# (.NET) | **GitHub**: [Consolonia/Consolonia](https://github.com/Consolonia/Consolonia) | **Stars**: ~808 | **Status**: Active

**Description**: Consolonia is a cross-platform terminal-based GUI framework for .NET built on top of Avalonia UI. It brings Avalonia's XAML-based UI development, data binding, templating, and styling to the terminal, allowing developers to reuse patterns from desktop GUI development.

**Key Features**:
- XAML-based declarative UI definition
- Avalonia UI framework integration
- Data binding and MVVM support
- UI templating and styling
- Rich theming system (Modern, Material)
- Responsive layouts
- Cross-platform (Windows, macOS, Linux)
- Async event handling
- Familiar to WPF/Avalonia developers

**Use Case**: Building .NET TUI applications with XAML, data binding, and MVVM patterns. Best for teams already using Avalonia/WPF who want to target the terminal.

---

### 42. Elaris.UI
**Language**: C# (.NET) | **GitHub**: [ambystechcom/Ambystech.Elaris.UI](https://github.com/ambystechcom/Ambystech.Elaris.UI) | **Stars**: ~20 | **Status**: Active

**Description**: Elaris.UI is a lightweight Terminal UI library for .NET with true 24-bit RGB color support. It provides a clean, extensible widget-based design with zero dependencies and double-buffered rendering.

**Key Features**:
- True 24-bit RGB ANSI color
- Zero external dependencies
- Double-buffered rendering
- Widget-based design
- Event loop management
- Keyboard and mouse input
- Multi-target (.NET 8, 9, 10)
- Cross-platform (Windows, macOS, Linux)
- Screen abstraction with buffering

**Use Case**: Lightweight .NET TUI applications needing true color support and no external dependencies. Good for simple tools and dashboards.

---

### 43. Hex1b
**Language**: C# (.NET) | **GitHub**: [mitchdenny/hex1b](https://github.com/mitchdenny/hex1b) | **Stars**: ~1K | **Status**: Active

**Description**: Hex1b is a .NET library for building rich terminal user interfaces with a React-inspired declarative API. It uses a virtual DOM approach where UI is composed using a fluent API with state management, widget trees, and event handling.

**Key Features**:
- React-inspired declarative API
- Virtual DOM diffing for efficient rendering
- Fluent widget composition (VStack, HStack, Border)
- State management with class-based state
- Built-in widgets: Text, Button, TextBox, List, InfoBar
- Focus management with Tab navigation
- Full-screen and flow (inline) modes
- Event handling (OnClick, OnTextChanged, etc.)
- Modular theming

**Use Case**: Building .NET TUIs with a React-like mental model. Good for interactive forms, wizards, todo apps, and any app that benefits from declarative UI composition.

---

### 44. SharpConsoleUI
**Language**: C# (.NET) | **GitHub**: [nickprotop/ConsoleEx](https://github.com/nickprotop/ConsoleEx) | **Stars**: ~100 | **Status**: Active

**Description**: SharpConsoleUI is a multi-window TUI framework for .NET with compositor effects. It features per-cell alpha blending, double-buffered compositing with occlusion culling, and a Measure-Arrange-Paint layout pipeline.

**Key Features**:
- Multi-window compositor with occlusion culling
- Per-cell Porter-Duff alpha blending
- Measure > Arrange > Paint layout pipeline
- Compositor effects: transitions, blur, fade, overlays
- CanvasControl for custom drawing (lines, circles, polygons, gradients)
- Post-processing hooks (PreBufferPaint, PostBufferPaint)
- Thread-safe buffer access
- Full Unicode support (Rune-based)
- Snapshot system for screenshots and recording

**Use Case**: Building visually sophisticated TUI applications with multi-window management and post-processing effects. Good for advanced dashboards and creative terminal apps.

---

### 45. Spectre.Console
**Language**: C# (.NET) | **GitHub**: [spectreconsole/spectre.console](https://github.com/spectreconsole/spectre.console) | **Stars**: ~10K+ | **Status**: Active

**Description**: Spectre.Console is a .NET library for creating beautiful console applications. It provides a comprehensive set of tools for rendering tables, charts, progress bars, calendars, trees, and more. It also includes a powerful widget system for building interactive prompts and form-like experiences.

**Key Features**:
- Table rendering with auto-sizing and alignment
- Progress bars with spinners and multi-line support
- Bar charts, breakdown charts, and calendars
- Tree views and file system rendering
- Figlet text and banner generation
- Canvases for pixel-level drawing
- Interactive prompts (selection, confirmation, text input)
- Multi-select, confirmation, and choice prompts
- AnsiConsole convenience API
- Markdown and JSON rendering
- Color system with color blending

**Use Case**: Beautifying .NET CLI output with tables, charts, progress bars, and interactive prompts. The standard choice for .NET CLI application output formatting (like Rich for Python).

---

### 46. Terminal.Gui
**Language**: C# (.NET) | **GitHub**: [gui-cs/Terminal.Gui](https://github.com/gui-cs/Terminal.Gui) | **Stars**: ~10K+ | **Status**: Active

**Description**: Terminal.Gui (gui-cs) is a cross-platform terminal UI toolkit for .NET that provides a full-featured application framework with over 50 built-in views (widgets). It supports responsive layouts, true color, Unicode, mouse, and keyboard-first interactions.

**Key Features**:
- 50+ built-in views/widgets
- Responsive layout system (web-like)
- Double-buffered rendering
- Rich text editor with syntax highlighting
- Data table with sorting, filtering, infinite scroll
- Tree views with virtualized nodes
- Charting and visualization widgets
- Color picker with true color
- Dialog and message box system
- Menu bars, status bars, tabs
- Theme and color scheme system
- Custom key binding configuration
- File dialog and file system views
- Focus-based keyboard navigation

**Use Case**: Building full-featured .NET TUI applications with the richest widget set available. The go-to framework for complex .NET terminal applications.

---

### .NET Libraries Comparison

| Library | Stars | Architecture | Widgets | Layout | XAML | Best For |
|---------|-------|-------------|---------|--------|------|----------|
| **Spectre.Console** | ~10K | Printer/widget | Tables, charts, progress | Auto | No | CLI output beautification |
| **Terminal.Gui** | ~10K | OOP framework | 50+ views | Responsive | No | Full TUI applications |
| **Consolonia** | ~808 | Avalonia/XAML | Avalonia widgets | XAML | Yes | XAML-based TUIs |
| **Hex1b** | ~1K | React-like/DOM | Basic widgets | VStack/HStack | No | Declarative TUIs |
| **SharpConsoleUI** | ~100 | Compositor | Window/Canvas | Measure-Arrange-Paint | No | Multi-window with effects |
| **Elaris.UI** | ~20 | Widget-based | Basic widgets | Manual | No | Minimal lightweight TUIs |

---

## Rust

Rust's TUI ecosystem is led by Ratatui (the community-maintained fork of tui-rs) and growing with innovative newcomers like iocraft and Zaz.

### 47. iocraft
**Language**: Rust | **GitHub**: [ccbrown/iocraft](https://github.com/ccbrown/iocraft) | **Stars**: ~1K | **Status**: Active

**Description**: iocraft is a Rust crate for crafting beautiful TUIs with a React-like declarative API, inspired by Ink (React for CLIs). It uses flexbox layouts powered by Taffy, supports hooks, event handling, and fullscreen applications.

**Key Features**:
- React-like declarative component API
- Flexbox layout via Taffy
- Element! macro for declarative UI
- Built-in components: View, Text, TextInput
- Custom components via #[component] macro
- Hooks for state and effects
- Event handling (keyboard, mouse)
- Fullscreen and inline modes
- Props and context by reference (no cloning)
- Unix and Windows support
- Output to terminal or ASCII strings

**Use Case**: Building TUIs in Rust with a React-like, declarative syntax. Great for Rust developers who want a component-based approach similar to Ink or React.

---

### 48. Ratatui
**Language**: Rust | **GitHub**: [ratatui/ratatui](https://github.com/ratatui/ratatui) | **Stars**: ~12K+ | **Status**: Active

**Description**: Ratatui is the community-maintained fork of the original tui-rs crate. It is the most popular Rust TUI library, providing a rich set of widgets, flexbox layout, and a proven instant-mode rendering engine. It powers major tools like bottom, kdash, and yazi.

**Key Features**:
- Rich widget set: Block, Table, List, Paragraph, Tabs, Chart, Gauge, Sparkline
- Flexbox-like layout with Constraints
- Instant-mode rendering (frame-based)
- Style system with colors, modifiers, and borders
- Canvas for custom drawing
- Scrollable regions
- Terminal backend abstraction (crossterm, termion, termwiz)
- Rect-based layout with chunks
- Bar chart and line chart widgets
- Comprehensive documentation and examples
- Active community with regular releases

**Use Case**: The standard choice for Rust TUI development. Use for any terminal application that needs rich widgets, layouts, and high performance.

---

### 49. tui-input
**Language**: Rust | **GitHub**: [sayantank/tui-input](https://github.com/sayantank/tui-input) | **Stars**: ~50 | **Status**: Active

**Description**: tui-input is a specialized Rust library for text input in TUI applications. It supports Ratatui backends and provides advanced input handling including undo/redo, text masking for passwords, input validation, and cursor management.

**Key Features**:
- Text input with undo/redo history
- Password masking support
- Input validation
- Cursor movement and selection
- Clipboard support
- Ratatui backend integration
- Vim-like keybindings
- Completion support

**Use Case**: Adding advanced text input functionality to Ratatui-based TUI applications.

---

### 50. tui-rs
**Language**: Rust | **GitHub**: [fdehau/tui-rs](https://github.com/fdehau/tui-rs) | **Stars**: ~14K | **Status**: NO LONGER MAINTAINED

**Description**: tui-rs was the original Rust TUI library that pioneered the instant-mode rendering approach with rich widgets. It is now archived and unmaintained. All users should migrate to Ratatui, the community-maintained fork.

**Key Features**:
- Instant-mode rendering
- Rich widget set
- Flexbox-like layout
- Cross-platform backends
- Comprehensive documentation

**Use Case**: Do NOT use for new projects. Migrate existing code to Ratatui.

---

### 51. Zaz
**Language**: Rust | **GitHub**: [raphamorim/zaz](https://github.com/raphamorim/zaz) | **Stars**: ~90 | **Status**: Active

**Description**: Zaz is a cross-platform Rust TUI library focused on efficient terminal rendering. It uses smart style caching, Paul Heckel's diff algorithm for minimal updates, cost-based cursor movement, and SIMD for large screens. Includes C/FFI bindings for other languages.

**Key Features**:
- Efficient rendering: smart style caching, Heckel diff, cost-based cursor movement
- SIMD acceleration for large screens (WIP)
- Kitty keyboard protocol support
- Graphics protocols: Kitty image, Sixel, iTerm2
- Unicode block mosaic rendering from images
- Panel and window management
- Z-ordering for overlapping elements
- RGB color with ANSI escape codes
- Terminal initialization and screen management
- C/FFI bindings for C++, Zig, and other languages

**Use Case**: High-performance Rust TUIs that need efficient rendering, image display, or graphics protocol support. Good for terminal multimedia apps.

---

### Rust Libraries Comparison

| Library | Stars | Architecture | Widgets | Rendering | Status | Best For |
|---------|-------|-------------|---------|-----------|--------|----------|
| **tui-rs** | ~14K | Instant-mode | Rich | Cell buffer | DEPRECATED | Do not use (migrate) |
| **Ratatui** | ~12K | Instant-mode | Rich | Cell buffer | Active | General Rust TUI (recommended) |
| **iocraft** | ~1K | React-like (declarative) | Basic | Flexbox/DOM | Active | Declarative TUIs |
| **Zaz** | ~90 | Efficient render | Panels/windows | Diff + SIMD | Active | Performance-critical TUIs |
| **tui-input** | ~50 | Input-focused | Text input only | Delegates to Ratatui | Active | Advanced text input |

---

## Other

Libraries in languages beyond the main categories, including Node.js, TypeScript, Swift, Nim, Dart, PHP, and shell.

### 52. Ashen
**Language**: Swift | **GitHub**: [colinta/Ashen](https://github.com/colinta/Ashen) | **Stars**: ~500 | **Status**: Maintained (low activity)

**Description**: Ashen is a framework for writing terminal applications in Swift, based on The Elm Architecture. It uses a declarative model-view-update pattern where the UI is a pure function of state.

**Key Features**:
- Elm Architecture (model, update, view)
- Declarative component composition
- Keyboard and mouse input
- Terminal rendering
- Swift-native API

**Use Case**: Building terminal applications in Swift with Elm-like architecture. Good for Swift developers who want functional TUI patterns.

---

### 53. blessed (Node.js)
**Language**: Node.js | **GitHub**: [chjj/blessed](https://github.com/chjj/blessed) | **Stars**: ~11K | **Status**: Maintained

**Description**: blessed is a high-level terminal interface library for Node.js. It provides a curses-like widget system for building rich terminal applications with screens, widgets, and mouse/keyboard input. It's the Node.js equivalent of ncurses with a modern API.

**Key Features**:
- Full widget system: boxes, buttons, text areas, lists, tables, forms
- Screen management with multiple terminals
- Mouse and keyboard input
- Unicode and wide character support
- Animated elements and progress bars
- Layout managers: grid, flex, absolute
- Evented API with Node.js streams
- Terminal multiplexing
- Virtual terminal support
- Scrollable content with auto-scroll

**Use Case**: Building rich terminal UIs in Node.js with a comprehensive widget set. The leading Node.js TUI library.

---

### 54. gum
**Language**: Go (shell tool) | **GitHub**: [charmbracelet/gum](https://github.com/charmbracelet/gum) | **Stars**: ~18K | **Status**: Active

**Description**: gum is a tool for creating glamorous shell scripts. It provides ready-made TUI utilities (input, choose, confirm, spin, write, file, filter, etc.) that can be composed in shell scripts without writing any Go code. It leverages the Charm Bubbles and Lip Gloss libraries.

**Key Features**:
- Pre-built TUI commands: choose, input, confirm, write, file, spin
- Filter for fuzzy searching
- Pager for scrolling content
- File picker dialog
- Write for multi-line text editing
- Spin for progress spinners
- Table and format utilities
- Style command for text formatting
- Join command for combining styled text
- Highly configurable via flags
- Customizable colors and styles
- Piped input/output for scripting

**Use Case**: Adding TUI interactions (selection, input, confirmation) to shell scripts without any programming. The easiest way to add a polished TUI to any bash/zsh script.

---

### 55. ink
**Language**: TypeScript/React (Node.js) | **GitHub**: [vadimdemedes/ink](https://github.com/vadimdemedes/ink) | **Stars**: ~37K | **Status**: Active

**Description**: ink is React for interactive command-line apps. It provides the same component-based UI building experience as React, but renders to the terminal using Yoga (Facebook's Flexbox layout engine). It powers major tools like Claude Code and opencode.

**Key Features**:
- Full React API for CLI output
- Flexbox layout via Yoga
- State management with React hooks
- Component composition and reusability
- Keyboard input handling (useInput hook)
- Focus management
- Color and text styling
- Transform and static rendering
- Custom reconciler for terminal output
- Thorough testing utilities
- Large ecosystem of third-party components

**Use Case**: Building interactive CLI applications with React. The premier choice for AI-powered CLI tools, interactive forms, and complex terminal UIs that benefit from React's component model.

---

### 56. ink-web
**Language**: TypeScript | **GitHub**: [cjroth/ink-web](https://github.com/cjroth/ink-web) | **Stars**: ~200 | **Status**: Active

**Description**: ink-web is a drop-in browser runtime for Ink that renders Ink components into an xterm.js terminal in the browser. It allows developers to build TUIs that work both in Node.js and in the browser via xterm.js.

**Key Features**:
- Drop-in replacement for Ink in the browser
- xterm.js-based terminal rendering
- Same Ink API in browser and Node.js
- Component registry (ink-ui) for browser-friendly widgets
- Webpack/alias-based setup
- Full TypeScript support
- Works with existing Ink components

**Use Case**: Running Ink-based TUIs in the browser (playgrounds, demos, web-based terminals). Good for providing a web UI alongside a CLI.

---

### 57. Melker
**Language**: TypeScript/Deno | **GitHub**: [wistrand/melker](https://github.com/wistrand/melker) | **Stars**: ~2K | **Status**: Active

**Description**: Melker is an HTML-like, document-first TUI framework. Melker apps are documents you can read before you run them — they declare permissions, have visible structure, and can be shared via URL. Supports Flexbox layout, CSS-like styling, tables, dialogs, and canvas graphics.

**Key Features**:
- Document-first: .melker files are readable markup
- Permission sandboxing (declared per-app permissions)
- Three abstraction levels: programmatic, declarative, literate
- HTML-like elements in .melker files
- CSS-like styling with 16M colors and auto theme detection
- Flexbox layout
- Tables with sorting and selection
- Dialogs and forms
- Canvas with sextant characters for graphics
- Mouse and keyboard throughout
- Dev Tools-style inspector
- Runs on Deno and Node.js

**Use Case**: Building terminal apps you want to share safely. Its document-first approach makes it ideal for distributable TUI apps where trust and inspectability matter.

---

### 58. moulti
**Language**: Python/shell | **GitHub**: [xavierog/moulti](https://github.com/xavierog/moulti) | **Stars**: ~2K | **Status**: Active

**Description**: moulti is a CLI-driven TUI tool that displays arbitrary command output inside visual, collapsible blocks called "steps". It's designed for shell scripts and Ansible playbooks, allowing scripts to organize their output into structured, navigable panels.

**Key Features**:
- Collapsible step blocks with titles and colors
- CLI-driven from shell scripts (pipe output to steps)
- Maximize single steps (like tmux zoom)
- Progress bars per step
- Quick search through steps
- Askpass support (moulti-askpass)
- Python/Textual-based backend
- Scrollback buffer with less-like navigation
- Unified diff display

**Use Case**: Organizing shell script output into structured, collapsible steps. Best for CI/CD scripts, deployment scripts, and Ansible playbooks that need better output organization.

---

### 59. nimwave
**Language**: Nim | **GitHub**: [ansiwave/nimwave](https://github.com/ansiwave/nimwave) | **Stars**: ~540 | **Status**: Active

**Description**: nimwave is a TUI framework for Nim that can render text interfaces to the terminal, desktop (via OpenGL/GLFW), and the web (via WebAssembly). It decouples TUIs from the terminal, allowing them to run anywhere.

**Key Features**:
- Multi-target: terminal, desktop (OpenGL/GLFW), web (WASM)
- Node-based UI hierarchy
- TerminalBuffer for low-level cell manipulation
- Color and style support
- Input handling
- Starter project template
- Cross-platform

**Use Case**: Building Nim TUIs that can run in the terminal, as a desktop app, or in the browser from a single codebase.

---

### 60. nocterm
**Language**: Dart | **GitHub**: [nocterm/nocterm](https://github.com/nocterm/nocterm) | **Stars**: ~200 | **Status**: Active

**Description**: nocterm is a Flutter-like TUI framework for Dart that provides the same APIs Flutter developers already know: StatefulComponent, setState(), Row, Column, Expanded, ListView. It also features hot reload for instant code updates during development.

**Key Features**:
- Flutter-like API (StatefulComponent, setState, Row, Column)
- Hot reload for instant updates
- State management
- Animations
- Component testing
- Layout widgets: Row, Column, Expanded, ListView
- Styling system
- Mouse and keyboard input

**Use Case**: Building Dart/Flutter-style TUIs in the terminal. Best for Flutter developers who want to leverage their existing knowledge for terminal applications.

---

### 61. OpenTUI
**Language**: TypeScript | **GitHub**: [sst/opentui](https://github.com/sst/opentui) | **Stars**: ~9K | **Status**: Active

**Description**: OpenTUI is a TypeScript library on a native Zig core for building terminal user interfaces. It provides a component-based architecture with flexible layout capabilities and React/SolidJS reconcilers. Powers OpenCode and terminal.shop in production.

**Key Features**:
- Native Zig core with TypeScript bindings
- Component-based architecture
- Flexbox layout via Yoga
- React reconciler (@opentui/react)
- SolidJS reconciler (@opentui/solid)
- Imperative API with all primitives
- C ABI for language interoperability
- High performance native rendering
- Focus on correctness and stability
- Extensible primitives

**Use Case**: Building high-performance TypeScript TUIs with React or SolidJS components. Best for production applications needing native rendering performance.

---

### 62. php-tui
**Language**: PHP | **GitHub**: [php-tui/php-tui](https://github.com/php-tui/php-tui) | **Stars**: ~1.5K | **Status**: Active

**Description**: php-tui is a comprehensive TUI library for PHP heavily inspired by Rust's Ratatui library. It provides the same widgets and shapes as Ratatui with advanced terminal control inspired by crossterm. Supports font and image rendering with Cassowary constraint-based layout.

**Key Features**:
- Ratatui-style API ported to PHP
- Rich widgets: table, list, paragraph, tabs, chart, gauge, sparkline, block
- Layout with Cassowary constraint solver
- Font and image rendering in terminal
- Advanced terminal control (colors, cursor, mouse)
- Crossterm-inspired terminal manipulation
- Full-screen TUI applications
- Canvas for custom drawing
- Border and style system

**Use Case**: Building full-screen TUI applications in PHP. Good for PHP developers who want Ratatui-style terminal apps without leaving PHP.

---

### 63. termbox2
*(See C section — listed above as entry 24b)*

**Language**: C | **GitHub**: [termbox/termbox2](https://github.com/termbox/termbox2) | **Stars**: ~720 | **Status**: Active

**Description**: termbox2 is a slim, modern alternative to ncurses for terminal I/O with a clean minimal API.

---

### 64. TermGL
*(See C section — listed above as entry 24c)*

**Language**: C | **GitHub**: [wojciech-graj/TermGL](https://github.com/wojciech-graj/TermGL) | **Stars**: ~500 | **Status**: Active

**Description**: TermGL is a terminal-based 2D/3D graphics library in C with shaders, transforms, and camera support.

---

### 65. Thermage
**Language**: PHP | **GitHub**: [thermage/thermage](https://github.com/thermage/thermage) | **Stars**: ~300 | **Status**: Active

**Description**: Thermage provides a fluent, powerful object-oriented interface for customizing CLI output text color, background, formatting, and theming in PHP. It uses a chainable, macroable API for composing styled terminal output.

**Key Features**:
- Fluent OOP interface for terminal styling
- Colors, backgrounds, and text formatting
- Border styles, padding, margin
- Theme and shortcode system
- Macroable for extending with custom elements
- Render to string or direct output
- Integration with Symfony Console, CodeIgniter
- Built-in elements: div, span, hr, br
- Responsive-like styling classes

**Use Case**: Beautifying PHP CLI output with styled text and formatting. Best for Laravel/Symfony artisan commands and CLI tools.

---

### Other Languages Comparison

| Library | Language | Stars | Architecture | Paradigm | Status | Best For |
|---------|----------|-------|-------------|----------|--------|----------|
| **ink** | TypeScript/React | ~37K | React reconciler | Component/React | Active | React CLI apps (recommended) |
| **gum** | Go/shell | ~18K | CLI utilities | Pipeline | Active | Shell script TUIs |
| **blessed** | Node.js | ~11K | Widget tree | Curses-like | Maintained | Node.js TUI apps |
| **OpenTUI** | TypeScript/Zig | ~9K | Component/React | Imperative/React | Active | Production TUIs |
| **php-tui** | PHP | ~1.5K | Ratatui port | Instant-mode | Active | PHP full-screen TUIs |
| **Melker** | TypeScript/Deno | ~2K | Document-first | HTML-like | Active | Shareable TUIs |
| **moulti** | Python/shell | ~2K | Step-based | CLI-driven | Active | Script output management |
| **nimwave** | Nim | ~540 | Node-based | Multi-target | Active | Cross-target TUIs |
| **Ashen** | Swift | ~500 | Elm Architecture | Functional | Low activity | Swift TUIs |
| **Thermage** | PHP | ~300 | Fluent OOP | Styling | Active | PHP CLI output styling |
| **nocterm** | Dart | ~200 | Flutter-like | Component | Active | Dart/Flutter TUIs |
| **ink-web** | TypeScript | ~200 | Runtime shim | React/browser | Active | Browser-based Ink |

---

## Quick Reference: Choosing a TUI Library

### By Language
| Language | Recommended Library(s) | Notes |
|----------|----------------------|-------|
| **Python** | textual (complex apps), Rich (CLI output), python-prompt-toolkit (REPLs) | textural for full apps, Rich for output |
| **Go** | bubbletea (general), tview (widget-heavy), pterm (output) | Bubble Tea is the ecosystem leader |
| **C** | termbox2 (modern), ncurses (compatibility), notcurses (bling) | termbox2 for new projects |
| **C++** | FTXUI (modern), imtui (ImGui-style) | FTXUI is the top recommendation |
| **Java** | Lanterna (mature), TUI4J (modern) | Lanterna for stability |
| **.NET** | Terminal.Gui (full), Spectre.Console (output), Hex1b (declarative) | Terminal.Gui for apps, Spectre for output |
| **Rust** | Ratatui (general), iocraft (declarative) | Ratatui is the standard |
| **TypeScript** | ink (React), OpenTUI (performance) | Ink for React devs, OpenTUI for speed |
| **Shell** | gum | The easiest way to add TUI to shell scripts |
| **PHP** | php-tui (full TUIs), Thermage (output styling) | php-tui for apps, Thermage for formatting |
| **Dart** | nocterm | Flutter-like API in terminal |

### By Use Case
| You want to... | Choose |
|----------------|--------|
| Build a full-screen TUI app | textual (Python), bubbletea (Go), Ratatui (Rust), Terminal.Gui (.NET), FTXUI (C++) |
| Add beautiful output to CLI | Rich (Python), pterm (Go), Spectre.Console (.NET), rang (C++) |
| Build a REPL/prompt | python-prompt-toolkit (Python) |
| Build a terminal game | ConsoleCraftEngine, ASCII_Board_Game_Engine (C++), TermGL (C) |
| Beautify shell scripts | gum |
| Build with React | ink (Node.js), OpenTUI (TypeScript) |
| Need max performance | Zaz (Rust), OpenTUI (Zig core), notcurses (C) |
| Need cross-platform | ncurses (C), Lanterna (Java), Terminal.Gui (.NET) |
| Need multimedia | notcurses (C), TermGL (C) |
| Want hot reload | nocterm (Dart) |
| Build a REPL/IDE | python-prompt-toolkit (Python), blessed (Node.js) |
| Port Turbo Vision apps | tvision (C++), Vindauga (Python) |
| Use Elm Architecture | bubbletea (Go), Ashen (Swift) |

---

## Version History

| Date | Author | Changes |
|------|--------|---------|
| 2026-07-13 | General | Initial comprehensive reference covering 65+ TUI libraries across 8 languages |
