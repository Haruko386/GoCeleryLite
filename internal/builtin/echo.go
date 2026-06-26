package builtin

import (
	"context"
	"encoding/json"
	"errors"
)

type EchoHandler struct{}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

type EchoPayload struct {
	Message string `json:"message"`
}

func (h *EchoHandler) Run(ctx context.Context, payload json.RawMessage) (any, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var echoPayload EchoPayload
	if err := json.Unmarshal(payload, &echoPayload); err != nil {
		return nil, err
	}

	if echoPayload.Message == "" {
		return nil, errors.New("message should not be empty")
	}

	return map[string]any{
		"message": echoPayload.Message,
	}, nil
}
