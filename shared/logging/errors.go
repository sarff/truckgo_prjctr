package logging

const (
	ErrInvalidEmail      = "InvalidEmail"
	ErrUserAlreadyExists = "UserAlreadyExists"
	ErrUserNotFound      = "UserNotFound"

	ErrDBQueryFailed  = "DBQueryFailed"
	ErrDBCreateFailed = "DBCreateFailed"
	ErrDBUpdateFailed = "DBUpdateFailed"

	ErrInvalidPassword       = "InvalidPassword"
	ErrPasswordHashingFailed = "PasswordHashingFailed"

	ErrTokenGenerationFailed = "TokenGenerationFailed"
	ErrTokenValidationFailed = "TokenValidationFailed"
	ErrInvalidToken          = "InvalidToken"

	ErrInvalidPhone = "InvalidPhone"
)
