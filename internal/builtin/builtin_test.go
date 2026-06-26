package builtin

import (
	"context"
	"encoding/json"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	h := NewEchoHandler()

	result, err := h.Run(context.Background(), json.RawMessage(`{"message":"hello"}`))
	if err != nil {
		t.Fatalf("echo handler failed: %v", err)
	}

	m, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected map result, got %T", result)
	}

	if m["message"] != "hello" {
		t.Fatalf("expected hello, got %v", m["message"])
	}
}

func TestHashHandler(t *testing.T) {
	h := NewHashHandler()

	result, err := h.Run(context.Background(), json.RawMessage(`{"text":"hello"}`))
	if err != nil {
		t.Fatalf("hash handler failed: %v", err)
	}

	m, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected map result, got %T", result)
	}

	expected := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"

	if m["sha256"] != expected {
		t.Fatalf("expected sha256 %s, got %v", expected, m["sha256"])
	}
}

func TestSleepHandler(t *testing.T) {
	h := NewSleepHandler()

	result, err := h.Run(context.Background(), json.RawMessage(`{"seconds":0}`))
	if err != nil {
		t.Fatalf("sleep handler failed: %v", err)
	}

	m, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected map result, got %T", result)
	}

	if m["slept_seconds"] != 0 {
		t.Fatalf("expected slept_seconds 0, got %v", m["slept_seconds"])
	}
}
