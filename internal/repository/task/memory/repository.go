package repository

import (
	"context"
	"fmt"
	"sync"
	"task-manager/internal/domain"
	"time"

	"github.com/google/uuid"
)

// Repository 實現 domain.TaskRepository 接口
type Repository struct {
	tasks map[string]*domain.Task
	mu    sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		tasks: make(map[string]*domain.Task),
		mu:    sync.RWMutex{},
	}
}

func (r *Repository) Create(ctx context.Context, task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	r.tasks[task.ID] = task
	return nil
}

// Delete removes a task by ID
func (r *Repository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return fmt.Errorf("task not found")
	}

	delete(r.tasks, id)
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return nil, fmt.Errorf("task not found")
	}

	return task, nil
}

func (r *Repository) Update(ctx context.Context, task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existingTask, ok := r.tasks[task.ID]
	if !ok {
		return fmt.Errorf("task not found")
	}

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.Priority = task.Priority
	existingTask.Status = task.Status
	existingTask.DueDate = task.DueDate
	existingTask.UpdatedAt = time.Now()

	return nil
}

func (r *Repository) List(ctx context.Context) ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
