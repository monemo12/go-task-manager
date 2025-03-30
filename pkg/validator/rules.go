package validator

import "time"

// NotEmptyRule 檢查字段不為空
type NotEmptyRule struct {
	Field string
}

func (r *NotEmptyRule) Validate(value interface{}) error {
	if str, ok := value.(string); ok && str == "" {
		return &ValidationError{
			Field:   r.Field,
			Message: "cannot be empty",
		}
	}
	return nil
}

// MaxLengthRule 檢查字段長度
type MaxLengthRule struct {
	Field     string
	MaxLength int
}

func (r *MaxLengthRule) Validate(value interface{}) error {
	if str, ok := value.(string); ok && len(str) > r.MaxLength {
		return &ValidationError{
			Field:   r.Field,
			Message: "exceeds maximum length",
		}
	}
	return nil
}

// FutureDateRule 檢查日期是否在未來
type FutureDateRule struct {
	Field string
}

func (r *FutureDateRule) Validate(value interface{}) error {
	if date, ok := value.(time.Time); ok && date.Before(time.Now()) {
		return &ValidationError{
			Field:   r.Field,
			Message: "must be a future date",
		}
	}
	return nil
}
