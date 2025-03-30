package service

import (
	"context"
	"time"

	"github.com/monemo12/task-manager/internal/domain"
)

// Service 實現 domain.TaskService 接口
type Service struct {
	repo      domain.TaskRepository
	validator domain.TaskValidator
	notifier  domain.TaskNotifier
}

// NewService 創建新的任務服務實例
func NewService(
	repo domain.TaskRepository,
	validator domain.TaskValidator,
	notifier domain.TaskNotifier,
) *Service {
	return &Service{
		repo:      repo,
		validator: validator,
		notifier:  notifier,
	}
}

// CreateTask 實現任務創建
func (s *Service) CreateTask(ctx context.Context, title, description string, priority domain.Priority, dueDate time.Time) (*domain.Task, error) {
	task := &domain.Task{
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      domain.Todo,
		DueDate:     dueDate,
	}

	if err := s.validator.ValidateTask(task); err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	s.notifier.NotifyTaskCreated(task)
	return task, nil
}

func (s *Service) UpdateTaskStatus(ctx context.Context, id string, status domain.Status) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	task.Status = status
	if err := s.repo.Update(ctx, task); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateTaskPriority(ctx context.Context, id string, priority domain.Priority) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	task.Priority = priority
	if err := s.repo.Update(ctx, task); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	return s.repo.List(ctx)
}

func (s *Service) DeleteTask(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
