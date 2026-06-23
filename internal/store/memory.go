package store

import (
	"GoCeleryLite/internal/task"
	"context"
	"errors"
	"sync"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrTaskExisted  = errors.New("task already exists")
	ErrTaskFailed   = errors.New("task failed")
	ErrTaskRunning  = errors.New("task is running")
)

type MemoryStore struct {
	mu    sync.RWMutex
	tasks map[string]*task.Task
}

// NewMemoryStore Create a new MemoryStorage
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: make(map[string]*task.Task),
	}
}

// Create a new task
func (s *MemoryStore) Create(ctx context.Context, task *task.Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// task is already exists
	if _, ok := s.tasks[task.ID]; ok {
		return ErrTaskExisted
	}
	// add the new task to dict
	s.tasks[task.ID] = task
	return nil
}

// Get the task
func (s *MemoryStore) Get(ctx context.Context, id string) (*task.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	// Get the task
	if _, ok := s.tasks[id]; !ok {
		return nil, ErrTaskNotFound
	}
	return s.tasks[id], nil
}

// List all tasks
func (s *MemoryStore) List(ctx context.Context) ([]*task.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	taskList := make([]*task.Task, 0, len(s.tasks))

	s.mu.RLock()
	defer s.mu.RUnlock()

	// iterating through the dict
	for _, t := range s.tasks {
		taskList = append(taskList, t)
	}
	return taskList, nil
}

// Update the task status
func (s *MemoryStore) Update(ctx context.Context, task *task.Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// check if task is existed
	if _, ok := s.tasks[task.ID]; !ok {
		return ErrTaskNotFound
	}

	s.tasks[task.ID] = task
	return nil
}

func (s *MemoryStore) Delete(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// check if task is existed
	if _, ok := s.tasks[id]; !ok {
		return ErrTaskNotFound
	}
	delete(s.tasks, id)
	return nil
}
