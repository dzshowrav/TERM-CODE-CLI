package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"termcode/pkg/apitypes"
)

type OpenAIAdapter struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewOpenAIAdapter(baseURL, apiKey string) *OpenAIAdapter {
	return &OpenAIAdapter{
		baseURL: baseURL,
		apiKey:  apiKey,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (a *OpenAIAdapter) Chat(ctx context.Context, req *apitypes.ChatRequest) (*apitypes.ChatResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, a.baseURL+"/v1/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if a.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)
	}

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (HTTP %d): %s", resp.StatusCode, string(respBody))
	}

	var chatResp apitypes.ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &chatResp, nil
}

func (a *OpenAIAdapter) ChatStream(ctx context.Context, req *apitypes.ChatRequest) (<-chan apitypes.StreamChunk, <-chan error) {
	chunks := make(chan apitypes.StreamChunk, 64)
	errs := make(chan error, 1)

	req.Stream = true

	go func() {
		defer close(chunks)
		defer close(errs)

		body, err := json.Marshal(req)
		if err != nil {
			errs <- fmt.Errorf("marshal request: %w", err)
			return
		}

		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, a.baseURL+"/v1/chat/completions", bytes.NewReader(body))
		if err != nil {
			errs <- fmt.Errorf("create request: %w", err)
			return
		}

		httpReq.Header.Set("Content-Type", "application/json")
		httpReq.Header.Set("Accept", "text/event-stream")
		if a.apiKey != "" {
			httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)
		}

		resp, err := a.client.Do(httpReq)
		if err != nil {
			errs <- fmt.Errorf("http request: %w", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			respBody, _ := io.ReadAll(resp.Body)
			errs <- fmt.Errorf("API error (HTTP %d): %s", resp.StatusCode, string(respBody))
			return
		}

		decoder := NewSSEDecoder(resp.Body)
		for {
			event, err := decoder.Decode()
			if err == io.EOF {
				return
			}
			if err != nil {
				errs <- fmt.Errorf("SSE decode: %w", err)
				return
			}

			if event.Data == "[DONE]" {
				return
			}

			var chunk apitypes.StreamChunk
			if err := json.Unmarshal([]byte(event.Data), &chunk); err != nil {
				errs <- fmt.Errorf("parse chunk: %w", err)
				return
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
