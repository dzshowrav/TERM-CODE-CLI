# AI CLI Tools — Complete A-to-Z Reference

> Essential AI-powered CLI tools from the Awesome CLI Apps ecosystem (75+ AI tools across 3 categories).
> Source: toolleeo/awesome-cli-apps-in-a-csv (AI/ChatGPT: 47, AI Command Generator: 16, Co-pilot: 12)

---

## 1. ollama — Local LLM Runner

**Website**: https://ollama.com  
**GitHub**: https://github.com/ollama/ollama  
**Language**: Go  
**Purpose**: Get up and running with large language models locally

### Installation
```bash
# macOS
brew install ollama
# Linux
curl -fsSL https://ollama.com/install.sh | sh
# Docker
docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama
```

### Model Management
```bash
# Pull a model
ollama pull llama3.2
ollama pull mistral
ollama pull codellama

# List models
ollama list

# Remove a model
ollama rm llama3.2

# Copy a model
ollama cp llama3.2 my-model
```

### Running Models
```bash
# Interactive chat
ollama run llama3.2

# One-shot
ollama run llama3.2 "Explain the difference between TCP and UDP"

# Pipe input
cat file.txt | ollama run llama3.2 "Summarize this:"

# With custom system prompt
ollama run llama3.2 --system "You are a helpful assistant that speaks like a pirate"
```

### API Mode (REST API)
```bash
# Start server
ollama serve

# API calls
curl http://localhost:11434/api/generate -d '{
  "model": "llama3.2",
  "prompt": "Why is the sky blue?"
}'

curl http://localhost:11434/api/chat -d '{
  "model": "llama3.2",
  "messages": [{ "role": "user", "content": "Hello" }]
}'
```

### API Endpoints
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/generate` | POST | Generate completion |
| `/api/chat` | POST | Chat completion |
| `/api/embeddings` | POST | Generate embeddings |
| `/api/create` | POST | Create model from Modelfile |
| `/api/pull` | POST | Pull model from registry |
| `/api/push` | POST | Push model to registry |
| `/api/tags` | GET | List local models |
| `/api/ps` | GET | List running models |
| `/api/show` | POST | Show model details |
| `/api/version` | GET | Version information |

### Custom Modelfile
```dockerfile
FROM llama3.2
# Set parameters
PARAMETER temperature 0.7
PARAMETER top_p 0.9
PARAMETER stop "```"
# Set system prompt
SYSTEM You are an expert CLI developer. Answer concisely.
```

### Key Differentiators
- **Easiest local LLM setup** — single command install
- **Model library** — hundreds of models available
- **REST API** — programmatic access
- **GPU acceleration** (NVIDIA, AMD, Apple Silicon)
- **OpenAPI-compatible** endpoint
- **Lightweight** — runs on consumer hardware

---

## 2. fabric — AI Pattern Framework

**GitHub**: https://github.com/danielmiessler/fabric  
**Language**: Go  
**Purpose**: Open-source framework for augmenting humans using AI

### Installation
```bash
# macOS
brew install fabric
# pip
pip install fabric-ai
# Manual
curl -fsSL https://raw.githubusercontent.com/danielmiessler/fabric/main/install.sh | sh
```

### Core Commands
| Command | Description |
|---------|-------------|
| `fabric` | Run fabric with AI |
| `fabric --setup` | Setup fabric configuration |
| `fabric --update` | Update patterns |
| `fabric --list` | List available patterns |
| `fabric --pattern <name>` | Use specific pattern |
| `fabric --model <model>` | Choose AI model |
| `fabric-stock` | Run stock fabric demo |
| `fabric-api` | Run fabric as API server |

### Pattern Usage
```bash
# Summarize content
cat article.txt | fabric --pattern summarize
# Extract wisdom
cat transcript.txt | fabric --pattern extract_wisdom
# AI analyze
cat code.ts | fabric --pattern analyze_code
# Create essay
cat notes.txt | fabric --pattern write_essay

# List all patterns
fabric --list

# Custom pattern
cat input.md | fabric --pattern my_custom_pattern
```

### Pattern Categories
| Category | Patterns |
|----------|----------|
| **Summarize** | summarize, summarize_meeting, summarize_podcast |
| **Extract** | extract_wisdom, extract_ideas, extract_questions |
| **Analyze** | analyze_code, analyze_claims, analyze_malware |
| **Create** | write_essay, write_micro essay, create_recipe |
| **Filter** | filter_lists, filter_reminder |
| **Improve** | improve_writing, improve_academic_writing |
| **Labels** | label_reviews, label_emails |
| **Custom** | User-created patterns |

### Configuration
```yaml
# ~/.config/fabric/config.yaml
model: gpt-4o
temperature: 0.7
max_tokens: 2000
patterns_dir: ~/.config/fabric/patterns
```

### Key Differentiators
- **Crowdsourced patterns** — community-contributed prompts
- **Pipe-friendly** — works in Unix pipelines
- **Pattern-first** — reusable AI instructions
- **Multi-model** — supports OpenAI, Anthropic, Ollama
- **Extensible** — create custom patterns

---

## 3. aider — AI Pair Programming

**GitHub**: https://github.com/paul-gauthier/aider  
**Language**: Python  
**Purpose**: AI pair programming in your terminal

### Installation
```bash
pip install aider-chat
# Or via uv
uv tool install aider-chat
```

### Quick Start
```bash
# Set API key
export OPENAI_API_KEY=your-key-here
# Or use Anthropic
export ANTHROPIC_API_KEY=your-key-here

# Start aider
aider
# With files already added
aider file1.py file2.ts
# Read-only files
aider --read file3.py
```

### Key Commands
| Command | Description |
|---------|-------------|
| `/add <file>` | Add file to chat context |
| `/read-only <file>` | Add read-only file |
| `/drop <file>` | Remove file from context |
| `/commit` | Commit changes |
| `/diff` | Show diff of changes |
| `/undo` | Undo last change |
| `/test <cmd>` | Run tests |
| `/lint <cmd>` | Run linter |
| `/model <name>` | Switch model |
| `/copy` | Copy markdown to clipboard |
| `/clear` | Clear chat |
| `/exit` | Exit aider |

### Key Options
| Flag | Description |
|------|-------------|
| `--model <model>` | Specify AI model (gpt-4, claude-3, etc.) |
| `--4turbo` | Use GPT-4 Turbo |
| `--opus` | Use Claude 3 Opus |
| `--sonnet` | Use Claude 3.5 Sonnet |
| `--deepseek` | Use DeepSeek |
| `--ollama` | Use Ollama model |
| `--openai-api-key` | OpenAI API key |
| `--anthropic-api-key` | Anthropic API key |
| `--dark-mode` | Dark mode UI |
| `--auto-commits` | Auto-commit changes |
| `--git` | Enable/disable git integration |
| `--pretty` | Pretty output with color |
| `--show-model-warnings` | Show model capability warnings |
| `--yes` | Auto-confirm changes |
| `--chat-mode <mode>` | Chat mode: `chat`, `ask`, `code` |
| `--lint` | Enable linting |
| `--test` | Enable testing |
| `--read <file>` | Read-only files to add |

### Workflows
```bash
# Chat mode (ask questions)
aider --chat-mode ask
# Code mode (make changes)
aider --chat-mode code
# Fix a specific file
aider buggy.py --model claude-sonnet-4-20250514
# Refactor with tests
aider src/ --test "npm test"
```

### Key Differentiators
- **Git-aware** — commits automatically, understands repo context
- **Map of repo** — understands project structure
- **Multi-file editing** — modifies multiple files simultaneously
- **Best model support** — Claude, GPT-4, DeepSeek, Ollama
- **Architect mode** — planning + coding workflow
- **Context-aware** — understands your entire codebase

---

## 4. Mods! (charmbracelet) — AI Pipeline Tool

**GitHub**: https://github.com/charmbracelet/mods  
**Language**: Go  
**Purpose**: AI for the command line, built for pipelines

### Installation
```bash
brew install charmbracelet/tap/mods
go install github.com/charmbracelet/mods@latest
```

### Usage
```bash
# Pipe content to AI
cat code.go | mods "Explain this code"
# With system prompt
echo "Hello world" | mods --system "Translate to French"
# No pipe — just ask
mods "What is the capital of France?"
# Interactive mode (--fzf)
mods --fzf
# Using file input
mods -f prompt.txt
```

### Key Options
| Flag | Description |
|------|-------------|
| `-m, --model <model>` | AI model |
| `-a, --api <api>` | API provider |
| `-s, --system <prompt>` | System prompt |
| `-f, --file <file>` | Read prompt from file |
| `--fzf` | Interactive fzf selection |
| `--heading` | Show response heading |
| `--format <format>` | Response format |
| `--temperature <n>` | Temperature (0-1) |
| `-n, --no-limit` | No response length limit |
| `-v, --verbose` | Verbose output |

### Configuration
```yaml
# ~/.config/mods/mods.yaml
provider: openai
model: gpt-4o
temperature: 0.7
system-prompt: "You are a helpful assistant."
```

### Key Differentiators
- **Pipeline-first design** — made for `| mods`
- **Bubble Tea TUI** — beautiful Charmbracelet UI
- **Multi-provider** — OpenAI, Anthropic, Ollama, Gemini
- **Low friction** — no setup needed beyond API key
- **fzf integration** for interactive selection

---

## 5. Chatblade — Versatile ChatGPT CLI

**GitHub**: https://github.com/npiv/chatblade  
**Language**: Python  
**Purpose**: Versatile CLI for interacting with ChatGPT

### Installation
```bash
pip install chatblade
```

### Usage
```bash
# Interactive chat
chatblade
# One-shot
chatblade "What is the speed of light?"
# Pipe content
cat data.txt | chatblade "Analyze this data"
# With system prompt
chatblade --system "You are a code reviewer"
# Continue conversation
chatblade -c "Explain more"
# Export conversation
chatblade --export
# List conversations
chatblade --list
```

### Key Options
| Flag | Description |
|------|-------------|
| `--system <prompt>` | System prompt |
| `-c, --continue` | Continue last conversation |
| `--list` | List saved conversations |
| `--last` | Show last conversation |
| `--export` | Export conversation to file |
| `--tokens` | Show token count |
| `--model <model>` | GPT model to use |
| `--temperature <n>` | Temperature |
| `--raw` | Raw output (no markdown) |

### Key Differentiators
- **Conversation management** — save and continue chats
- **Session support** — multiple independent conversations
- **Token counting** — see token usage per message
- **Markdown rendering** — formatted output
- **Export/import** conversations

---

## 6. Gemini CLI — Google's Official CLI

**GitHub**: https://github.com/google-gemini/gemini-cli  
**Language**: Python  
**Purpose**: Official Google CLI for Gemini

### Installation
```bash
pip install google-genai-sdk
# Or via npm
npm install -g @google-gemini/cli
```

### Usage
```bash
# Basic query
gemini "Explain quantum computing"
# Interactive mode
gemini --interactive
# With context
cat code.py | gemini "Review this code"
# Vision (image input)
gemini "Describe this image" --image photo.jpg
# Multi-turn
gemini --interactive --model gemini-2.0-flash
```

### Key Differentiators
- **Official Google product** — first-party support
- **Gemini 2.0 Flash / Pro** — latest models
- **Vision capabilities** — image understanding
- **Google ecosystem** — Workspace integration
- **Multi-modal** — text, image, code

---

## 7. aichat — Universal AI CLI

**GitHub**: https://github.com/sigoden/aichat  
**Language**: Rust  
**Purpose**: Using ChatGPT/GPT-4/Gemini in the terminal

### Installation
```bash
brew install aichat
cargo install aichat
```

### Usage
```bash
# Chat
aichat "Hello"
# With roles
aichat --role code_reviewer "Review this"
# Session mode
aichat --session my-session
# Role list
aichat --list-roles
# Create role
aichat --create-role
# File context
aichat --file code.ts
```

### Key Differentiators
- **Fast** (Rust) — quick startup
- **Multi-provider** — OpenAI, Anthropic, Gemini, Ollama, Azure
- **Role system** — predefined roles (translator, coder, reviewer)
- **Session management** — persistent conversations
- **Syntax highlighting** in output
- **Streaming** responses

---

## 8. OpenCode — AI Coding Agent

**Website**: https://opencode.ai/download  
**Purpose**: AI coding agent built for the terminal

### Key Features
- **Terminal-native** AI coding assistant
- **Context-aware** code modifications
- **Multi-file editing**
- **Git integration**
- **Tool execution** — run commands, tests, etc.

---

## 9. fabric Patterns Reference

### Most Useful Patterns
| Pattern | Purpose |
|---------|---------|
| `summarize` | General text summarization |
| `extract_wisdom` | Extract key insights |
| `analyze_code` | Code review and analysis |
| `analyze_claims` | Fact-check statements |
| `write_essay` | Generate structured essays |
| `improve_writing` | Grammar/style improvement |
| `create_prompt` | Generate better AI prompts |
| `label_reviews` | Categorize product reviews |
| `extract_questions` | Pull out questions from text |
| `summarize_meeting` | Meeting notes summary |

### Creating Custom Patterns
```yaml
# ~/.config/fabric/patterns/my_pattern/system.md
You are an expert at analyzing CLI tool output.
Focus on error messages, warnings, and actionable insights.

# ~/.config/fabric/patterns/my_pattern/config.yaml
name: "cli_analyzer"
model: gpt-4o
temperature: 0.5
```

---

## Summary: When to Use Which AI CLI Tool

| Tool | Best For | Key Strength |
|------|----------|--------------|
| **ollama** | Local LLM inference | Privacy, offline, no API costs |
| **fabric** | Reusable AI patterns | Pattern library, pipeline-friendly |
| **aider** | AI pair programming | Git-aware code changes |
| **Mods!** | Quick AI in pipelines | Charmbracelet UX, multi-provider |
| **Chatblade** | Conversational AI | Chat sessions, token tracking |
| **Gemini CLI** | Google AI ecosystem | Official support, vision |
| **aichat** | Universal AI access | Speed (Rust), multi-provider |
| **OpenCode** | Autonomous coding | Terminal-native AI agent |
