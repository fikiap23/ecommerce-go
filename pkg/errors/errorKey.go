package errors

type ErrorKey string

const (
	// Naming convention: ERR_<NOUN>_<VERB/ACTION>
	// ----------------------------------

	ErrInputInvalid          ErrorKey = "ERR_INPUT_INVALID"
	ErrValidationFailed      ErrorKey = "ERR_VALIDATION_FAILED"
	ErrUserNotFound          ErrorKey = "ERR_USER_NOT_FOUND"
	ErrUserCreationFailed    ErrorKey = "ERR_USER_CREATION_FAILED"
	ErrUserUpdateFailed      ErrorKey = "ERR_USER_UPDATE_FAILED"
	ErrUserAlreadyVerified   ErrorKey = "ERR_USER_ALREADY_VERIFIED"
	ErrUserVerificationInvalid ErrorKey = "ERR_USER_VERIFICATION_FAILED"
	ErrUserVerificationExpired ErrorKey = "ERR_USER_VERIFICATION_EXPIRED"
	ErrPasswordInvalid       ErrorKey = "ERR_PASSWORD_INVALID"
	ErrPasswordHashFailed    ErrorKey = "ERR_PASSWORD_HASH_FAILED"
	ErrEmailAlreadyExists    ErrorKey = "ERR_EMAIL_ALREADY_EXISTS"
)

