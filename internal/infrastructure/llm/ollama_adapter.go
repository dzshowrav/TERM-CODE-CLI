package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"termcode/pkg/apitypes"
)

type OllamaAdapter struct {
	baseURL string
	client  *http.Client
}

func NewOllamaAdapter(baseURL string) *OllamaAdapter {
	u := strings.TrimRight(baseURL, "/")
	u = strings.TrimSuffix(u, "/v1")
	return &OllamaAdapter{
		baseURL: u,
		client: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

func (a *OllamaAdapter) Chat(ctx context.Context, req *apitypes.ChatRequest) (*apitypes.ChatResponse, error) {
	ollamaReq := a.toOllamaRequest(req)
	body, err := json.Marshal(ollamaReq)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, a.baseURL+"/api/chat", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		msg := string(respBody)
		if isHTML(msg) {
			msg = "server returned non-JSON response — check your Ollama URL"
		}
		return nil, fmt.Errorf("Ollama error (HTTP %d): %s", resp.StatusCode, msg)
	}

	var ollamaResp ollamaChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return a.fromOllamaResponse(&ollamaResp, req.Model), nil
}

func (a *OllamaAdapter) ChatStream(ctx context.Context, req *apitypes.ChatRequest) (<-chan apitypes.StreamChunk, <-chan error) {
	chunks := make(chan apitypes.StreamChunk, 64)
	errs := make(chan error, 1)

	ollamaReq := a.toOllamaRequest(req)
	ollamaReq.Stream = true

	go func() {
		defer close(chunks)
		defer close(errs)

		body, err := json.Marshal(ollamaReq)
		if err != nil {
			errs <- fmt.Errorf("marshal request: %w", err)
			return
		}

		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, a.baseURL+"/api/chat", bytes.NewReader(body))
		if err != nil {
			errs <- fmt.Errorf("create request: %w", err)
			return
		}
		httpReq.Header.Set("Content-Type", "application/json")

		resp, err := a.client.Do(httpReq)
		if err != nil {
			errs <- fmt.Errorf("http request: %w", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			respBody, _ := io.ReadAll(resp.Body)
			msg := string(respBody)
			if isHTML(msg) {
				msg = "server returned non-JSON response — check your Ollama URL"
			}
			errs <- fmt.Errorf("Ollama error (HTTP %d): %s", resp.StatusCode, msg)
			return
		}

		decoder := json.NewDecoder(resp.Body)
		for {
			var ollamaChunk ollamaStreamChunk
			if err := decoder.Decode(&ollamaChunk); err == io.EOF {
				return
			} else if err != nil {
				errs <- fmt.Errorf("decode chunk: %w", err)
				return
			}

			if ollamaChunk.Done {
				return
			}

			chunk := apitypes.StreamChunk{
				Model:   ollamaChunk.Model,
				Created: time.Now().Unix(),
				Choices: []apitypes.StreamChoice{
					{
						Delta: apitypes.StreamDelta{
							Content: ollamaChunk.Message.Content,
						},
					},
				},
			}

			select {
			case chunks <- chunk:
			case <-ctx.Done():
				errs <- ctx.Err()
				return
			}
		}
	}()

	return chunks, errs
}

type ollamaMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ollamaChatRequest struct {
	Model    string          `json:"model"`
	Messages []ollamaMessage `json:"messages"`
	Stream   bool            `json:"stream"`
	Options  map[string]any  `json:"options,omitempty"`
}

type ollamaChatResponse struct {
	Model     string        `json:"model"`
	CreatedAt string        `json:"created_at"`
	Message   ollamaMessage `json:"message"`
	Done      bool          `json:"done"`
}

type ollamaStreamChunk struct {
	Model     string        `json:"model"`
	CreatedAt string        `json:"created_at"`
	Message   ollamaMessage `json:"message"`
	Done      bool          `json:"done"`
}

func (a *OllamaAdapter) toOllamaRequest(req *apitypes.ChatRequest) *ollamaChatRequest {
	messages := make([]ollamaMessage, len(req.Messages))
	for i, m := range req.Messages {
		messages[i] = ollamaMessage{Role: m.Role, Content: m.Content}
	}
	return &ollamaChatRequest{
		Model:    req.Model,
		Messages: messages,
		Options: map[string]any{
			"temperature": req.Temperature,
		},
	}
}

func (a *OllamaAdapter) fromOllamaResponse(resp *ollamaChatResponse, model string) *apitypes.ChatResponse {
	return &apitypes.ChatResponse{
		Model: model,
		Choices: []apitypes.Choice{
			{
				Message: apitypes.ChatMessage{
					Role:    resp.Message.Role,
					Content: resp.Message.Content,
				},
			},
		},
	}
}
