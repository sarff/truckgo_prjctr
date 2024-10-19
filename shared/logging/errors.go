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
	ErrDBDeleteFailed = "DBDeleteFailed"
	// Password
	ErrInvalidPassword       = "InvalidPassword"
	ErrPasswordHashingFailed = "PasswordHashingFailed"
	ErrPasswordUpdateFailed  = "PasswordUpdateFailed"
	//Token
	ErrTokenGenerationFailed = "TokenGenerationFailed"
	ErrTokenValidationFailed = "TokenValidationFailed"
	ErrInvalidToken          = "InvalidToken"
	//Phone
	ErrInvalidPhone = "InvalidPhone"
	// Info
	MsgInfo = "INFO"
)
