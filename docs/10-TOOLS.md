# OpenChat CLI

# Tool System Specification

Version: 1.0

---

# Overview

The Tool System is the execution layer of OpenChat CLI.

Unlike a normal chatbot that only generates text, OpenChat CLI can safely interact with the user's development environment through a controlled set of tools.

Every tool is executed through the Tool Manager.

The AI never directly accesses the operating system.

All requests pass through

AI

↓

Permission Manager

↓

Tool Manager

↓

Operating System

↓

Result

↓

AI

---

# Philosophy

Tools extend the AI's capabilities.

The AI should never guess.

Instead it should

Read

Inspect

Search

Execute

Verify

Then answer.

---

# Tool Architecture

User Prompt

↓

AI

↓

Tool Decision

↓

Permission Check

↓

Tool Execution

↓

Result

↓

AI Response

---

# Design Goals

✓ Safe

✓ Transparent

✓ Fast

✓ Extensible

✓ Plugin Ready

✓ Mobile Friendly

✓ Provider Independent

✓ Permission Controlled

---

# Built-in Tools

Filesystem

Search

Editor

Terminal

Git

Workspace

Network

Clipboard

MCP

Task

Database

Diagnostics

Future Tools

---

# Filesystem Tools

read_file

Read file contents

write_file

Create file

edit_file

Modify file

append_file

Append content

rename_file

Rename file

delete_file

Delete file

create_directory

Create folder

delete_directory

Delete folder

list_directory

List files

file_info

Metadata

exists

Check existence

---

# Search Tools

search_text

Search project

grep

Regular expression search

glob

Pattern search

find_file

Locate file

find_symbol

Language symbols

find_reference

LSP references

---

# Code Editing Tools

replace_text

Replace selection

insert_text

Insert

remove_text

Delete selection

format_document

Code formatter

organize_imports

Language Server

---

# Terminal Tools

bash

Execute command

shell

Interactive shell

run_script

Execute script

stop_process

Kill process

view_process

Running processes

---

# Git Tools

git_status

Status

git_diff

Diff

git_log

History

git_branch

Branches

git_checkout

Checkout

git_commit

Commit

git_add

Stage

git_restore

Restore

git_pull

Pull

git_push

Push

---

# Workspace Tools

scan_workspace

Analyze project

detect_framework

Framework detection

detect_language

Programming language

dependency_tree

Package analysis

project_summary

Overview

---

# HTTP Tools

http_get

GET request

http_post

POST request

download

Download file

upload

Upload file

api_test

Test endpoint

---

# Database Tools

sqlite_query

SQLite

mysql_query

MySQL

postgres_query

PostgreSQL

Future

Redis

MongoDB

Supabase

Firebase

---

# Clipboard Tools

copy

Copy text

paste

Paste text

clear_clipboard

Clear

---

# Task Tools

todo_add

Add task

todo_remove

Delete task

todo_complete

Complete

todo_list

List tasks

---

# MCP Tools

connect_server

disconnect_server

list_servers

discover_tools

invoke_tool

reload_server

---

# Diagnostic Tools

system_info

Environment

disk_usage

Disk

memory_usage

Memory

cpu_usage

CPU

network_status

Network

---

# Tool Categories

Read Only

Safe

Write

Execution

Dangerous

Network

Plugin

---

# Permission Levels

Always Allow

Allow Once

Ask

Always Deny

Disabled

Every tool belongs to one permission group.

---

# Dangerous Tools

delete_file

delete_directory

bash

git_push

git_commit

database_write

network_upload

Always require confirmation unless explicitly trusted.

---

# Tool UI

When a tool runs

Display

✓ Reading composer.json

✓ Searching routes

✓ Executing bash

✓ Writing file

✓ Git Diff Generated

Users always know what is happening.

---

# Tool Output

Every tool returns

Status

Success

Failure

Execution Time

Output

Errors

Warnings

Metadata

Example

Status

Success

Time

12ms

Output

composer.json loaded

---

# Tool Result Format

{
    success: true,
    data: "...",
    duration: 18,
    warnings: [],
    errors: []
}

Every tool should return a consistent structure.

---

# Parallel Execution

Safe tools

May run simultaneously.

Example

Read File

↓

Read Package

↓

Search Project

↓

Read Git

Parallel execution improves speed.

---

# Sequential Execution

Required for

Write

Delete

Rename

Bash

Git

Database

Operations that modify state execute in order.

---

# Tool Timeouts

Read

5 seconds

Search

10 seconds

Bash

Configurable

Download

60 seconds

Long-running tools may be cancelled.

---

# Cancellation

Ctrl+C

↓

Cancel Tool

↓

Keep Session

↓

Continue Conversation

---

# Logging

Every tool stores

Start Time

Finish Time

Duration

Arguments

Status

Errors

Logs never include API keys or secrets.

---

# Tool Search

Command

/tools

Search

↓

read

↓

read_file

↓

Enter

↓

View Tool

---

# Tool Details

Selecting a tool displays

Name

Description

Category

Permissions

Arguments

Examples

Execution History

---

# Enable / Disable

Users may disable

Unused

Experimental

Plugin

Tools

Disabled tools are hidden from the AI.

---

# Tool Aliases

read

↓

read_file

write

↓

write_file

edit

↓

edit_file

delete

↓

delete_file

Simple aliases improve usability.

---

# Tool Versioning

Every tool stores

Version

Author

Source

Plugin

Compatibility

---

# Plugin Tools

Plugins may register

New Commands

New Tools

New Permissions

New Categories

The core system should automatically discover them.

---

# Tool Capabilities

Each tool advertises

Read

Write

Execute

Network

Filesystem

Database

Streaming

Async

The AI uses these capabilities when planning actions.

---

# Performance Goals

Tool Discovery

<20ms

Execution Start

<50ms

Read File

<10ms

Search

<100ms

Parallel Scheduling

Automatic

Minimal memory usage

---

# Future Tools

Browser Automation

Playwright

Docker

Kubernetes

SSH

FTP/SFTP

Cloud Storage

Figma

Notion

Slack

Jira

VS Code

Android Debug Bridge (ADB)

Local AI Models

Voice Input

OCR

Image Editing

---

# Design Principles

The Tool System is the bridge between the AI and the real development environment.

Every action must be visible, permission-controlled, and reversible whenever possible.

The AI should never silently modify the user's project.

Instead, OpenChat CLI should make every tool invocation transparent, safe, and predictable while remaining fast enough to feel like a natural extension of the conversation.

The Tool System is designed to be modular so that new capabilities can be added through plugins or MCP servers without requiring changes to the core application.