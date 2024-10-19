package logging

const (
	// Login
	ErrInvalidEmail      = "InvalidEmail"
	ErrUserAlreadyExists = "UserAlreadyExists"
	ErrUserNotFound      = "UserNotFound"
	// DB
	ErrDBQueryFailed  = "DBQueryFailed"
	ErrDBCreateFailed = "DBCreateFailed"
	ErrDBUpdateFailed = "DBUpdateFailed"
	// Password
	ErrInvalidPassword       = "InvalidPassword"
	ErrPasswordHashingFailed = "PasswordHashingFailed"
	//Token
	ErrTokenGenerationFailed = "TokenGenerationFailed"
	ErrTokenValidationFailed = "TokenValidationFailed"
	ErrInvalidToken          = "InvalidToken"
	//Phone
	ErrInvalidPhone = "InvalidPhone"
)
