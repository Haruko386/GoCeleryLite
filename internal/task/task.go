package task

import (
	"encoding/json"
	"time"
)

type Task struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Payload    json.RawMessage `json:"payload,omitempty"`
	Status     Status          `json:"status"`
	Result     any             `json:"result,omitempty"`
	Error      string          `json:"error,omitempty"`
	Retries    int             `json:"retries"`
	MaxRetries int             `json:"max_retries"`
	TimeoutSec int             `json:"timeout_sec"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

// NewTask Create a new task
func NewTask(id, name string, payload json.RawMessage) *Task {
	now := time.Now()

	return &Task{
		ID:         id,
		Name:       name,
		Payload:    payload,
		Status:     StatusQueued,
		Retries:    0,
		MaxRetries: 0,
		TimeoutSec: 30,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

// SetStatus Set task's status
func (t *Task) SetStatus(status Status) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

// SetResult Set task as done
func (t *Task) SetResult(result any) {
	t.Result = result
	t.UpdatedAt = time.Now()
	t.Status = StatusSuccess
	t.Error = ""
}

// SetError Set task's error when task was failed
func (t *Task) SetError(err error) {
	if err != nil {
		t.Error = err.Error()
	}
	t.Status = StatusFailed
	t.UpdatedAt = time.Now()
}
