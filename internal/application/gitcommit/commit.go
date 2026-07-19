package gitcommit

import (
	"context"
	"fmt"
	"strings"
)

type LLMClient interface {
	Generate(ctx context.Context, prompt string) (string, error)
}

type Generator struct {
	llm LLMClient
}

func NewGenerator(llm LLMClient) *Generator {
	return &Generator{llm: llm}
}

func (g *Generator) Generate(ctx context.Context, diff string) (string, error) {
	if diff == "" {
		return "", fmt.Errorf("diff is empty")
	}

	prompt := fmt.Sprintf(`Generate a concise git commit message for the following diff.
Follow conventional commits format: <type>(<scope>): <description>

Types: feat, fix, docs, style, refactor, test, chore
Scope is optional.

Diff:
%s

Commit message:`, diff)

	msg, err := g.llm.Generate(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("generate commit message: %w", err)
	}

	msg = strings.TrimSpace(msg)
	msg = strings.SplitN(msg, "\n", 2)[0]

	if len(msg) > 72 {
		msg = msg[:72]
	}

	return msg, nil
}
