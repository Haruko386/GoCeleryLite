package store

import (
	"context"
	"encoding/json"
	"testing"

	"GoCeleryLite/internal/task"
)

func TestMemoryStoreCreateAndGet(t *testing.T) {
	s := NewMemoryStore()

	payload := json.RawMessage(`{"seconds":5}`)
	task1 := task.NewTask("task-1", "sleep", payload)

	err := s.Create(context.Background(), task1)
	if err != nil {
		t.Fatalf("create task failed: %v", err)
	}

	got, err := s.Get(context.Background(), "task-1")
	if err != nil {
		t.Fatalf("get task failed: %v", err)
	}

	if got.ID != "task-1" {
		t.Fatalf("expected task id task-1, got %s", got.ID)
	}

	if got.Name != "sleep" {
		t.Fatalf("expected task name sleep, got %s", got.Name)
	}

	if got.Status != task.StatusQueued {
		t.Fatalf("expected status queued, got %s", got.Status)
	}
}

func TestMemoryStoreGetNotFound(t *testing.T) {
	s := NewMemoryStore()

	_, err := s.Get(context.Background(), "not-exist")
	if err != ErrTaskNotFound {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestMemoryStoreUpdate(t *testing.T) {
	s := NewMemoryStore()

	task1 := task.NewTask("task-1", "sleep", json.RawMessage(`{"seconds":5}`))

	if err := s.Create(context.Background(), task1); err != nil {
		t.Fatalf("create task failed: %v", err)
	}

	task1.SetStatus(task.StatusRunning)

	if err := s.Update(context.Background(), task1); err != nil {
		t.Fatalf("update task failed: %v", err)
	}

	got, err := s.Get(context.Background(), "task-1")
	if err != nil {
		t.Fatalf("get task failed: %v", err)
	}

	if got.Status != task.StatusRunning {
		t.Fatalf("expected status running, got %s", got.Status)
	}
}
