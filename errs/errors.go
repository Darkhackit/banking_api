package errs

import "net/http"

type AppErrors struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppErrors) AsMessage() *AppErrors {
	return &AppErrors{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppErrors {
	return &AppErrors{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
func NewUnexpectedError(message string) *AppErrors {
	return &AppErrors{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
