package deps

import (
	// Core TUI
	_ "charm.land/bubbletea/v2"
	_ "github.com/charmbracelet/bubbles"
	_ "github.com/charmbracelet/glamour"
	_ "github.com/charmbracelet/harmonica"
	_ "github.com/charmbracelet/huh"
	_ "github.com/charmbracelet/lipgloss"
	_ "github.com/charmbracelet/log"

	// CLI Framework
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/pflag"
	_ "github.com/spf13/viper"

	// Terminal Rendering
	_ "github.com/mattn/go-runewidth"
	_ "github.com/muesli/ansi"
	_ "github.com/muesli/reflow/wordwrap"
	_ "github.com/muesli/termenv"
	_ "github.com/rivo/uniseg"

	// AI / MCP
	_ "github.com/modelcontextprotocol/go-sdk/mcp"

	// HTTP / Networking
	_ "github.com/go-resty/resty/v2"
	_ "github.com/gorilla/websocket"
	_ "golang.org/x/net/context"

	// File System
	_ "github.com/fsnotify/fsnotify"
	_ "github.com/go-git/go-git/v5"
	_ "github.com/shirou/gopsutil/v4"

	// Database
	_ "github.com/jackc/pgx/v5"
	_ "github.com/redis/go-redis/v9"
	_ "modernc.org/sqlite"

	// Configuration
	_ "github.com/joho/godotenv"
	_ "github.com/knadh/koanf"

	// JSON / YAML / TOML
	_ "github.com/pelletier/go-toml/v2"
	_ "github.com/tidwall/gjson"
	_ "github.com/tidwall/sjson"
	_ "gopkg.in/yaml.v3"

	// Markdown
	_ "github.com/yuin/goldmark"

	// Tree-sitter
	_ "github.com/tree-sitter/go-tree-sitter"

	// Search
	_ "github.com/lithammer/fuzzysearch/fuzzy"
	_ "github.com/sahilm/fuzzy"

	// Utilities
	_ "github.com/go-playground/validator/v10"
	_ "github.com/google/uuid"
	_ "github.com/patrickmn/go-cache"

	// Concurrency
	_ "github.com/avast/retry-go/v4"
	_ "github.com/panjf2000/ants/v2"

	// Progress
	_ "github.com/briandowns/spinner"
	_ "github.com/schollz/progressbar/v3"
	_ "github.com/vbauerster/mpb/v8"

	// Tables
	_ "github.com/jedib0t/go-pretty/v6/table"
	_ "github.com/olekukonko/tablewriter"

	// Diff
	_ "github.com/hexops/gotextdiff"
	_ "github.com/sergi/go-diff/diffmatchpatch"

	// Interactive Prompt
	_ "github.com/AlecAivazis/survey/v2"
	_ "github.com/manifoldco/promptui"

	// Terminal Utilities
	_ "github.com/fatih/color"
	_ "github.com/mattn/go-isatty"

	// Logging
	_ "go.uber.org/zap"

	// Security
	_ "golang.org/x/crypto/bcrypt"
	_ "golang.org/x/sys/unix"

	// Testing
	_ "github.com/stretchr/testify"
)
