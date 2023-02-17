package errs

import "net/http"

type AppError struct {
	Status  int
	Message string
}

func (a AppError) Error() string {
	return a.Message
}

func NewError(code int, errMsg string) error {
	return AppError{
		Status:  code,
		Message: errMsg,
	}
}

func ErrorBadRequest(errorMessage string) error {
	return AppError{
		Status:  http.StatusBadRequest,
		Message: errorMessage,
	}
}
func ErrorUnprocessableEntity(errorMessage string) error {
	return AppError{
		Status:  http.StatusUnprocessableEntity,
		Message: errorMessage,
	}
}

func ErrorInternalServerError(errorMessage string) error {
	return AppError{
		Status:  http.StatusInternalServerError,
		Message: errorMessage,
	}
}
