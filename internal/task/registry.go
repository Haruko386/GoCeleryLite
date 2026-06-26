package task

import (
	"errors"
	"sort"
	"sync"
)

var (
	ErrHandlerNotFound = errors.New("task handler not found")
	ErrHandlerExists   = errors.New("task handler already exists")
	ErrInvalidHandler  = errors.New("invalid task handler")
)

type Registry struct {
	mu       sync.RWMutex
	handlers map[string]Handler
}

func NewRegistry() *Registry {
	return &Registry{
		handlers: make(map[string]Handler),
	}
}

// Register a new handler
func (r *Registry) Register(name string, handler Handler) error {
	if name == "" || handler == nil {
		return ErrInvalidHandler
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.handlers[name]; ok {
		return ErrHandlerExists
	}

	r.handlers[name] = handler
	return nil
}

// Get handler
func (r *Registry) Get(name string) (Handler, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	handler, ok := r.handlers[name]
	if !ok {
		return nil, ErrHandlerNotFound
	}

	return handler, nil
}

// List all handler name
func (r *Registry) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]string, 0, len(r.handlers))
	for name := range r.handlers {
		list = append(list, name)
	}
	sort.Strings(list)

	return list
}

func (r *Registry) Name() string {
	return "registry"
}
