package errors

type ErrorKey string

const (
	ErrInvalidInput    ErrorKey = "ERR_INVALID_INPUT"
	ErrValidationFailed ErrorKey = "ERR_VALIDATION_FAILED"
	ErrUserNotFound    ErrorKey = "ERR_USER_NOT_FOUND"
	ErrInvalidPassword ErrorKey = "ERR_INVALID_PASSWORD"
	ErrEmailExists     ErrorKey = "ERR_EMAIL_EXISTS"
)
