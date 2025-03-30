package validator

// Validator 定義通用的驗證器接口
type Validator interface {
	// Validate 驗證給定的結構體
	// 返回錯誤切片，如果沒有錯誤則返回 nil
	Validate(interface{}) []error
}

// ValidationError 表示驗證錯誤
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// BaseValidator 提供基本的驗證功能
type BaseValidator struct {
	rules map[string][]ValidationRule
}

// ValidationRule 定義驗證規則接口
type ValidationRule interface {
	Validate(value interface{}) error
}

// NewBaseValidator 創建新的基本驗證器
func NewBaseValidator() *BaseValidator {
	return &BaseValidator{
		rules: make(map[string][]ValidationRule),
	}
}

// AddRule 添加驗證規則
func (v *BaseValidator) AddRule(field string, rule ValidationRule) {
	if v.rules[field] == nil {
		v.rules[field] = make([]ValidationRule, 0)
	}
	v.rules[field] = append(v.rules[field], rule)
}

// Validate 執行驗證
func (v *BaseValidator) Validate(value interface{}) []error {
	var errors []error
	// 具體的驗證邏輯將在這裡實現
	return errors
}
