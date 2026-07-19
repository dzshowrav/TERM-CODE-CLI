package apitypes

type ChatMessage struct {
	Role             string     `json:"role"`
	Content          string     `json:"content"`
	ReasoningContent string     `json:"reasoning_content,omitempty"`
	Reasoning        string     `json:"reasoning,omitempty"`
	ToolCalls        []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID       string     `json:"tool_call_id,omitempty"`
}

func (m ChatMessage) ReasoningText() string {
	if m.ReasoningContent != "" {
		return m.ReasoningContent
	}
	return m.Reasoning
}

type ToolCall struct {
	ID       string           `json:"id"`
	Type     string           `json:"type"`
	Function ToolCallFunction `json:"function"`
}

type ToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ChatRequest struct {
	Model       string           `json:"model"`
	Messages    []ChatMessage    `json:"messages"`
	Stream      bool             `json:"stream"`
	Temperature float64          `json:"temperature,omitempty"`
	MaxTokens   int              `json:"max_tokens,omitempty"`
	Tools       []map[string]any `json:"tools,omitempty"`
}

type ChatResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   *Usage   `json:"usage,omitempty"`
}

type Choice struct {
	Index            int         `json:"index"`
	Message          ChatMessage `json:"message"`
	FinishReason     string      `json:"finish_reason"`
	ReasoningContent string      `json:"reasoning_content,omitempty"`
	Reasoning        string      `json:"reasoning,omitempty"`
}

func (c Choice) ReasoningText() string {
	if c.ReasoningContent != "" {
		return c.ReasoningContent
	}
	return c.Reasoning
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type StreamChunk struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []StreamChoice `json:"choices"`
}

type StreamChoice struct {
	Index        int         `json:"index"`
	Delta        StreamDelta `json:"delta"`
	FinishReason string      `json:"finish_reason"`
}

type StreamDelta struct {
	Role             string     `json:"role,omitempty"`
	Content          string     `json:"content,omitempty"`
	ReasoningContent string     `json:"reasoning_content,omitempty"`
	Reasoning        string     `json:"reasoning,omitempty"`
	ToolCalls        []ToolCall `json:"tool_calls,omitempty"`
}

func (d StreamDelta) ReasoningText() string {
	if d.ReasoningContent != "" {
		return d.ReasoningContent
	}
	return d.Reasoning
}
