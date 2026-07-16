package mcp

import "errors"

var (
	ErrNotFound      = errors.New("MCP server not found")
	ErrDuplicate     = errors.New("MCP server name already exists")
	ErrNotConnected  = errors.New("MCP server is not connected")
	ErrInvalidConfig = errors.New("invalid MCP server configuration")
)
