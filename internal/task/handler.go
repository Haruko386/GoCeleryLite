package task

import (
	"context"
	"encoding/json"
)

type Handler interface {
	Run(ctx context.Context, payload json.RawMessage) (any, error)
}
