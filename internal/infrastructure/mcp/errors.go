package mcp

import "errors"

var (
	ErrNotConnected         = errors.New("MCP client not connected")
	ErrConnectionFailed     = errors.New("MCP connection failed")
	ErrToolNotFound         = errors.New("MCP tool not found")
	ErrToolExecution        = errors.New("MCP tool execution failed")
	ErrTimeout              = errors.New("MCP request timed out")
	ErrTransportClosed      = errors.New("MCP transport closed")
	ErrInvalidResponse      = errors.New("invalid MCP response")
	ErrProtocolVersion      = errors.New("unsupported MCP protocol version")
	ErrServerNotInitialized = errors.New("MCP server not initialized")
)
