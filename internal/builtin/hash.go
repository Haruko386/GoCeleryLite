package builtin

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type HashHandler struct{}

func NewHashHandler() *HashHandler {
	return &HashHandler{}
}

type HashPayload struct {
	Text string `json:"text"`
}

func (h HashHandler) Run(ctx context.Context, payload json.RawMessage) (any, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	var hashPayload HashPayload
	if err := json.Unmarshal(payload, &hashPayload); err != nil {
		return nil, err
	}

	sum := sha256.Sum256([]byte(hashPayload.Text))

	return map[string]any{
		"text":   hashPayload.Text,
		"sha256": hex.EncodeToString(sum[:]),
	}, nil
}
