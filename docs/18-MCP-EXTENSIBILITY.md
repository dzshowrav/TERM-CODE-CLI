# MCP & Extensibility

## Overview

The extensibility system allows the agent to connect to external services and load custom capabilities at runtime. The primary mechanism is the Model Context Protocol (MCP), but the system supports other extension patterns as well.

## MCP (Model Context Protocol)

MCP is a standard protocol for connecting AI agents to external tools and data sources.

### MCP Server
An MCP server is an external process that exposes tools and resources to the agent:
- Each server runs as a separate process
- Communication happens via stdin/stdout (stdio) or HTTP/SSE
- The server advertises its capabilities: tools, resources, prompts
- Multiple servers can run simultaneously

### MCP Connection Lifecycle
1. **Discovered** — server configuration found
2. **Connecting** — server process is started
3. **Handshaking** — capabilities are exchanged
4. **Connected** — server is active and ready
5. **Reconnecting** — connection lost, attempting to restore
6. **Disconnected** — server stopped or unreachable
7. **Error** — connection failed permanently

### MCP Tool Exposure
When an MCP server is connected:
- Its tools are registered in the agent's tool system
- Tools are namespaced (e.g., `mcp_server_name.tool_name`)
- Tool descriptions include the server name for context
- Permission system applies to MCP tools like any other tool
- The user can see which MCP server provides each tool

### MCP Resource Access
MCP servers can provide resources:
- File contents, database results, API responses
- Resources are accessible via a special tool or inline reference
- Resources are cached to avoid repeated requests

## Plugin System

Beyond MCP, the agent supports a native plugin architecture:

### Plugin Definition
- **Entry point** — how the plugin is loaded (file, URL, package)
- **Capabilities** — what the plugin provides (tools, hooks, UI elements)
- **Permissions** — what the plugin can access (filesystem, network, environment)
- **Dependencies** — other plugins or runtime requirements

### Plugin Capabilities
Plugins can provide:
- **Tools** — executed in the tool system
- **Hooks** — intercept and extend core behavior (e.g., onMessage, onToolCall)
- **Theme extensions** — add theme variables or color schemes
- **UI components** — custom display elements
- **Key bindings** — custom keyboard shortcuts
- **Commands** — extend the slash command system

### Plugin Isolation
- Plugins run in a restricted environment
- Plugins cannot access other plugins' data
- Plugins have defined permission boundaries
- A misbehaving plugin can be disabled without affecting the core

## Custom Tool Definition

Advanced users can define custom tools without writing a plugin:

### Inline Tool Definition
A tool defined in configuration:
```yaml
name: my_custom_tool
description: Does something useful
parameters:
  - name: input
    type: string
    description: Input parameter
handler: "shell command that receives arguments as env vars"
```

### Script Tool
A tool backed by a script:
- The script can be in any language
- Arguments are passed as environment variables or CLI arguments
- Output is captured and returned as the tool result
- Scripts can be version-controlled alongside the project

## Service Integration

The agent can integrate with external services:

### Version Control
- Direct integration with git hosting (GitHub, GitLab, Bitbucket)
- Create PRs, review code, manage issues
- Webhook support for CI/CD events

### Cloud Services
- Deploy code, manage infrastructure
- Query logs and metrics
- Manage secrets and configuration

### Communication
- Send notifications (Slack, Discord, email)
- Create and update tickets
- Post comments on PRs and issues

All service integrations respect the permission system — no service action happens without user approval.

## Extension Discovery

Extensions (MCP servers, plugins, services) can be:
- **Pre-installed** — shipped with the agent
- **Discovered** — found in a marketplace or registry
- **User-added** — manually configured
- **Workspace-local** — defined in the project's configuration

## Key Design Decisions

- MCP is the primary extension mechanism — standard, language-agnostic protocol
- Plugins provide deeper integration when MCP isn't sufficient
- Custom tool definitions are configuration-only — no code required
- All extensions are sandboxed — the core agent is protected from misbehaving extensions
- Service integrations always go through the permission system
- Extensions can be enabled/disabled per-session
- Discovery mechanisms make extensions findable without being intrusive
- Extensions can be project-specific (defined in workspace config)
