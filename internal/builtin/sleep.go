package builtin

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

type SleepHandler struct{}

type SleepPayload struct {
	Seconds int `json:"seconds"`
}

func NewSleepHandler() *SleepHandler {
	return &SleepHandler{}
}

func (h *SleepHandler) Run(ctx context.Context, payload json.RawMessage) (any, error) {
	var p SleepPayload
	if err := json.Unmarshal(payload, &p); err != nil {
		return nil, err
	}

	if p.Seconds < 0 {
		return nil, errors.New("sleep time seconds must be greater than or equal to 0")
	}

	timer := time.NewTimer(time.Duration(p.Seconds) * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		return map[string]any{
			"message":       "sleep task finished",
			"slept_seconds": p.Seconds,
		}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
