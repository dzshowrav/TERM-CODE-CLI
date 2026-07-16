package llm

import (
	"encoding/json"
	"fmt"

	"termcode/pkg/apitypes"
)

type ParsedResponse struct {
	Content      string
	ToolCalls    []ToolCallInfo
	FinishReason string
	Usage        *UsageInfo
}

type ToolCallInfo struct {
	ID        string
	Name      string
	Arguments map[string]any
}

type UsageInfo struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type ResponseParser struct{}

func NewResponseParser() *ResponseParser {
	return &ResponseParser{}
}

func (p *ResponseParser) Parse(resp *apitypes.ChatResponse) (*ParsedResponse, error) {
	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no response choices")
	}

	choice := resp.Choices[0]
	parsed := &ParsedResponse{
		Content:      choice.Message.Content,
		FinishReason: choice.FinishReason,
	}

	if resp.Usage != nil {
		parsed.Usage = &UsageInfo{
			InputTokens:  resp.Usage.PromptTokens,
			OutputTokens: resp.Usage.CompletionTokens,
		}
	}

	if len(choice.Message.ToolCalls) > 0 {
		for _, tc := range choice.Message.ToolCalls {
			var args map[string]any
			if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err != nil {
				args = map[string]any{}
			}
			parsed.ToolCalls = append(parsed.ToolCalls, ToolCallInfo{
				ID:        tc.ID,
				Name:      tc.Function.Name,
				Arguments: args,
			})
		}
	}

	return parsed, nil
}

func (p *ResponseParser) IsToolCall(finishReason string) bool {
	return finishReason == "tool_calls"
}
