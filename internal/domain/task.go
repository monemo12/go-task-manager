package domain

import (
	"context"
	"time"
)

// Priority 定義任務優先級
type Priority int

const (
	Low Priority = iota
	Medium
	High
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	default:
		return "unknown"
	}
}

func (p Priority) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.String() + `"`), nil
}

func (p *Priority) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"low"`:
		*p = Low
	case `"medium"`:
		*p = Medium
	case `"high"`:
		*p = High
	default:
		*p = Low
	}
	return nil
}

// Status 定義任務狀態
type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	switch s {
	case Todo:
		return "todo"
	case InProgress:
		return "in_progress"
	case Done:
		return "done"
	default:
		return "unknown"
	}
}

func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func (s *Status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"todo"`:
		*s = Todo
	case `"in_progress"`:
		*s = InProgress
	case `"done"`:
		*s = Done
	default:
		*s = Todo
	}
	return nil
}

// Task 代表一個任務的基本屬性
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DueDate     time.Time `json:"due_date"`
}

// TaskRepository 定義任務儲存的接口
// 遵循接口隔離原則(I)，每個接口只包含其必需的方法
type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	GetByID(ctx context.Context, id string) (*Task, error)
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Task, error)
}

// TaskService 定義任務管理的業務邏輯接口
type TaskService interface {
	CreateTask(ctx context.Context, title, description string, priority Priority, dueDate time.Time) (*Task, error)
	UpdateTaskStatus(ctx context.Context, id string, status Status) error
	UpdateTaskPriority(ctx context.Context, id string, priority Priority) error
	GetTask(ctx context.Context, id string) (*Task, error)
	ListTasks(ctx context.Context) ([]*Task, error)
	DeleteTask(ctx context.Context, id string) error
}

// TaskValidator 定義任務驗證的接口
type TaskValidator interface {
	ValidateTask(task *Task) error
}

// TaskNotifier 定義任務通知的接口
type TaskNotifier interface {
	NotifyTaskCreated(task *Task)
	NotifyTaskCompleted(task *Task)
	NotifyTaskDueSoon(task *Task)
}
