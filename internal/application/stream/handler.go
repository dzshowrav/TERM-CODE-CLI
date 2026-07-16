package stream

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"termcode/internal/infrastructure/llm"
	"termcode/pkg/apitypes"
)

type ChunkHandler func(content string, done bool)

type Handler struct {
	adapter *llm.OpenAIAdapter
	tracker *Tracker
	mu      sync.Mutex
	buffer  strings.Builder
}

func NewHandler(adapter *llm.OpenAIAdapter, tracker *Tracker) *Handler {
	return &Handler{
		adapter: adapter,
		tracker: tracker,
	}
}

func (h *Handler) Stream(ctx context.Context, req *apitypes.ChatRequest, onChunk ChunkHandler) (string, error) {
	h.mu.Lock()
	h.buffer.Reset()
	h.mu.Unlock()

	chunks, errs := h.adapter.ChatStream(ctx, req)

	fullContent := strings.Builder{}

	for {
		select {
		case chunk, ok := <-chunks:
			if !ok {
				chunks = nil
				continue
			}

			for _, choice := range chunk.Choices {
				content := choice.Delta.Content
				if content == "" {
					continue
				}

				fullContent.WriteString(content)

				tokenCount := estimateTokens(content)
				h.tracker.TrackOutput(tokenCount)

				if onChunk != nil {
					onChunk(content, choice.FinishReason != "")
				}
			}

			if len(chunk.Choices) > 0 && chunk.Choices[0].FinishReason != "" {
				return fullContent.String(), nil
			}

		case err, ok := <-errs:
			if !ok {
				errs = nil
				continue
			}
			if err != nil {
				return fullContent.String(), fmt.Errorf("stream error: %w", err)
			}

		case <-ctx.Done():
			return fullContent.String(), ctx.Err()
		}

		if chunks == nil && errs == nil {
			break
		}
	}

	return fullContent.String(), nil
}

func (h *Handler) Complete(ctx context.Context, req *apitypes.ChatRequest) (string, error) {
	resp, err := h.adapter.Chat(ctx, req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices")
	}

	content := resp.Choices[0].Message.Content

	if resp.Usage != nil {
		h.tracker.TrackInput(resp.Usage.PromptTokens)
		h.tracker.TrackOutput(resp.Usage.CompletionTokens)
	}

	return content, nil
}

func estimateTokens(text string) int {
	if len(text) == 0 {
		return 0
	}
	tokens := len(text) / 4
	if tokens < 1 {
		tokens = 1
	}
	return tokens
}
