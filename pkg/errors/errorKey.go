package errors

type ErrorKey string

const (
	// Naming convention: ERR_<NOUN>_<VERB/ACTION>
	// ----------------------------------

	ErrInputInvalid          ErrorKey = "ERR_INPUT_INVALID"
	ErrValidationFailed      ErrorKey = "ERR_VALIDATION_FAILED"
	ErrUserNotFound          ErrorKey = "ERR_USER_NOT_FOUND"
	ErrUserCreationFailed    ErrorKey = "ERR_USER_CREATION_FAILED"
	ErrPasswordInvalid       ErrorKey = "ERR_PASSWORD_INVALID"
	ErrPasswordHashFailed    ErrorKey = "ERR_PASSWORD_HASH_FAILED"
	ErrEmailAlreadyExists    ErrorKey = "ERR_EMAIL_ALREADY_EXISTS"
)

