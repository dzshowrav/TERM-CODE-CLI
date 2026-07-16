# OpenChat CLI

# Model Context Protocol (MCP) Specification

Version: 1.0

---

# Overview

The Model Context Protocol (MCP) allows OpenChat CLI to connect to external tools, applications, databases, cloud services, APIs, and automation systems through a standardized protocol.

Instead of building every integration into OpenChat CLI, MCP enables external servers to expose capabilities dynamically.

OpenChat CLI acts as an MCP Client.

---

# Philosophy

AI

↓

OpenChat CLI

↓

MCP Client

↓

MCP Server

↓

Tools

↓

Resources

↓

Prompts

↓

External Systems

The AI should not know where a capability comes from.

It simply requests a tool.

---

# Goals

✓ Unlimited MCP Servers

✓ Local Servers

✓ Remote Servers

✓ Dynamic Tool Discovery

✓ Secure

✓ Permission Controlled

✓ Hot Reload

✓ Plugin Compatible

✓ Provider Independent

---

# Architecture

User

↓

Chat Engine

↓

Tool Manager

↓

MCP Manager

↓

Connected Servers

↓

External Services

---

# Supported Transport

stdio

HTTP

HTTPS

SSE (Server Sent Events)

WebSocket (Future)

Unix Socket (Future)

Named Pipe (Windows)

---

# MCP Manager

Responsible for

Connect Server

Disconnect Server

Reconnect

Health Check

Authentication

Permission

Tool Discovery

Prompt Discovery

Resource Discovery

Hot Reload

---

# Database Schema

mcp_servers

id

name

description

transport

command

arguments

url

environment

status

enabled

auto_connect

created_at

updated_at

---

# Add MCP Server

Command

/mcp add

Dialog

Server Name

[_______________]

Transport

▼ stdio

Command

[_______________]

Arguments

[_______________]

Working Directory

[_______________]

Environment Variables

[_______________]

Auto Connect

☑

Save

Cancel

---

# Remote MCP Server

Transport

HTTP

Fields

Name

URL

Authentication

Headers

Timeout

Auto Connect

---

# Authentication

Supported

None

Bearer Token

API Key

Basic Auth

Custom Headers

OAuth (Future)

---

# Server Status

Connected

Connecting

Disconnected

Offline

Authentication Failed

Timeout

Unknown

---

# Server Health

Every server stores

Last Ping

Latency

Version

Capabilities

Last Error

Reconnect Count

---

# Tool Discovery

When connected

↓

Server advertises tools

↓

OpenChat registers tools

↓

AI can use them immediately

No restart required.

---

# Resource Discovery

Servers may expose

Files

Databases

Logs

API Documentation

Knowledge Bases

Configuration

Images

Documents

Resources become available automatically.

---

# Prompt Discovery

Servers may expose reusable prompts.

Example

Code Review

API Documentation

Deployment Guide

Bug Report

Architecture Review

Users can invoke them directly.

---

# Tool Categories

Filesystem

Database

Cloud

Browser

Automation

Design

Office

Communication

AI

Custom

---

# Example MCP Servers

GitHub

GitLab

Docker

Kubernetes

Figma

Slack

Discord

Jira

Notion

Supabase

Firebase

PostgreSQL

MySQL

Redis

AWS

Azure

Google Cloud

Linux

Android

Local Scripts

---

# Android Optimizations

Designed for Termux.

Supports

Python MCP

Node MCP

Rust MCP

Go MCP

Bash MCP

No Docker required.

---

# Permissions

Every MCP Tool follows

Permission Manager

Allow Once

Always Allow

Ask

Always Deny

---

# Hot Reload

Command

/mcp reload

Reloads

Servers

Tools

Resources

Prompts

No application restart.

---

# Auto Connect

Servers marked

Auto Connect

Start automatically when OpenChat launches.

---

# Tool Invocation

AI

↓

Tool Request

↓

Permission Check

↓

MCP Server

↓

Execution

↓

Result

↓

AI Continues

---

# Error Recovery

Connection Lost

↓

Reconnect

↓

Retry

↓

Notify User

↓

Continue Session

---

# Search

Command

/mcp

Supports searching

Servers

Tools

Resources

Prompts

---

# Export

Export MCP Configuration

JSON

YAML (Future)

---

# Import

Import

↓

Validate

↓

Install

↓

Ready

---

# Performance

Server Connection

<500ms

Tool Discovery

Instant

Hot Reload

<1 second

---

# Security

Never execute unknown tools automatically.

Never expose secrets to AI unless required.

Mask authentication values in logs.

Require explicit permission for dangerous operations.

---

# Future Features

MCP Marketplace

Server Ratings

Cloud MCP Registry

One-Click Install

Server Templates

Shared Team Servers

Multi-User Authentication

---

# Design Principles

The MCP system transforms OpenChat CLI into an extensible development platform.

Instead of shipping hundreds of built-in integrations, OpenChat CLI discovers capabilities from external MCP servers at runtime.

This keeps the core lightweight while allowing users to connect any ecosystem they need, from GitHub and Docker to custom internal company tools, all through a consistent, secure, and provider-independent interface.