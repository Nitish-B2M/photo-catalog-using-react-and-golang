package errors

const (
	InvalidCredentials = iota
	InsufficientFunds
	FileTooLarge
	NoPermission
	InternalServerError
)

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

var customErrors = map[int]string{
	InvalidCredentials:  "Invalid credentials",
	InsufficientFunds:   "Insufficient funds",
	FileTooLarge:        "File too large",
	NoPermission:        "No permission to perform this action",
	InternalServerError: "Internal server error",
}

func GetCustomErrorMessage(code int) string {
	return customErrors[code]
}
