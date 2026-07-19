package main

import (
	"fmt"
	"termcode/internal/adapters/tui/components"
)

func main() {
	card := components.NewToolCard("read")
	card.SetOutput("line 1\nline 2")
	card.SetArgs(`{"path": "/tmp/test.txt"}`)
	card.SetStatus(components.ToolCompleted)
	card.SetExpanded(false)
	fmt.Println(card.View())
}
