package utils

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	DevMessage string `json:"devMessage"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%d: %s, %s", e.Code, e.Message, e.DevMessage)
}

func NewAppError(code int, message, devMessage string) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		DevMessage: devMessage,
	}
}

func StatusCode(err error) int {
	var appError *AppError
	if errors.As(err, &appError) {
		return appError.Code
	} else {
		return http.StatusInternalServerError
	}

}
