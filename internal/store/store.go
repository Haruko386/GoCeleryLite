package store

import (
	"GoCeleryLite/internal/task"
	"context"
)

type Store interface {
	Create(ctx context.Context, task *task.Task) error
	Get(ctx context.Context, id string) (*task.Task, error)
	List(ctx context.Context) ([]*task.Task, error)
	Update(ctx context.Context, task *task.Task) error
	Delete(ctx context.Context, id string) error
}
