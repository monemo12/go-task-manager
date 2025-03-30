package errors

import "fmt"

// AppError 定義應用程序錯誤接口
type AppError interface {
	error
	Code() string
	Message() string
}

// TaskError 實現 AppError 接口
type TaskError struct {
	code    string
	message string
}

func (e *TaskError) Error() string {
	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

func (e *TaskError) Code() string {
	return e.code
}

func (e *TaskError) Message() string {
	return e.message
}

// NewTaskError 創建新的任務錯誤
func NewTaskError(code, message string) AppError {
	return &TaskError{
		code:    code,
		message: message,
	}
}

// 預定義的錯誤代碼
const (
	ErrTaskNotFound     = "TASK_NOT_FOUND"
	ErrTaskInvalidInput = "TASK_INVALID_INPUT"
	ErrTaskDuplicate    = "TASK_DUPLICATE"
	ErrInternalServer   = "INTERNAL_SERVER_ERROR"
)
