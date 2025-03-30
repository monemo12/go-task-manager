package validator

import (
	"task-manager/internal/domain"
)

// TaskValidator 實現任務的具體驗證邏輯
type TaskValidator struct {
	*BaseValidator
}

// NewTaskValidator 創建新的任務驗證器
func NewTaskValidator() *TaskValidator {
	v := &TaskValidator{
		BaseValidator: NewBaseValidator(),
	}
	v.registerRules()
	return v
}

// registerRules 註冊任務相關的驗證規則
func (v *TaskValidator) registerRules() {
	v.AddRule("Title", &NotEmptyRule{Field: "Title"})
	v.AddRule("Title", &MaxLengthRule{Field: "Title", MaxLength: 100})
	v.AddRule("DueDate", &FutureDateRule{Field: "DueDate"})
}

// ValidateTask 實現 domain.TaskValidator 接口
func (v *TaskValidator) ValidateTask(task *domain.Task) error {
	errors := v.Validate(task)
	if len(errors) > 0 {
		return errors[0] // 返回第一個錯誤
	}
	return nil
}

// 其他驗證規則的定義...
