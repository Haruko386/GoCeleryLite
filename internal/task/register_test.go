package task

import (
	"context"
	"encoding/json"
	"testing"
)

type fakeHandler struct{}

func (h *fakeHandler) Run(ctx context.Context, payload json.RawMessage) (any, error) {
	return "ok", nil
}

func TestRegistryRegisterAndGet(t *testing.T) {
	r := NewRegistry()

	err := r.Register("fake", &fakeHandler{})
	if err != nil {
		t.Fatalf("register handler failed: %v", err)
	}

	handler, err := r.Get("fake")
	if err != nil {
		t.Fatalf("get handler failed: %v", err)
	}

	result, err := handler.Run(context.Background(), nil)
	if err != nil {
		t.Fatalf("run handler failed: %v", err)
	}

	if result != "ok" {
		t.Fatalf("expected result ok, got %v", result)
	}
}

func TestRegistryDuplicateRegister(t *testing.T) {
	r := NewRegistry()

	if err := r.Register("fake", &fakeHandler{}); err != nil {
		t.Fatalf("register handler failed: %v", err)
	}

	err := r.Register("fake", &fakeHandler{})
	if err != ErrHandlerExists {
		t.Fatalf("expected ErrHandlerExists, got %v", err)
	}
}

func TestRegistryGetNotFound(t *testing.T) {
	r := NewRegistry()

	_, err := r.Get("not-exist")
	if err != ErrHandlerNotFound {
		t.Fatalf("expected ErrHandlerNotFound, got %v", err)
	}
}

func TestRegistryList(t *testing.T) {
	r := NewRegistry()

	_ = r.Register("sleep", &fakeHandler{})
	_ = r.Register("echo", &fakeHandler{})
	_ = r.Register("hash", &fakeHandler{})

	names := r.List()

	expected := []string{"echo", "hash", "sleep"}

	if len(names) != len(expected) {
		t.Fatalf("expected %d handlers, got %d", len(expected), len(names))
	}

	for i := range expected {
		if names[i] != expected[i] {
			t.Fatalf("expected %s at index %d, got %s", expected[i], i, names[i])
		}
	}
}
