package errors

import (
	"fmt"
	"net/http"
)

// AppError represents a custom application error with HTTP status code
type AppError struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Details    []string `json:"details,omitempty"`
	StatusCode int      `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

// NewAppError creates a new AppError
func NewAppError(code, message string, statusCode int, details ...string) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		Details:    details,
		StatusCode: statusCode,
	}
}

// ValidationError creates a validation error (400)
func ValidationError(message string, details ...string) *AppError {
	return NewAppError("VALIDATION_ERROR", message, http.StatusBadRequest, details...)
}

// NotFoundError creates a not found error (404)
func NotFoundError(resource string) *AppError {
	return NewAppError(
		"NOT_FOUND",
		fmt.Sprintf("%s not found", resource),
		http.StatusNotFound,
	)
}

// ConflictError creates a conflict error (409)
func ConflictError(message string) *AppError {
	return NewAppError("CONFLICT", message, http.StatusConflict)
}

// InternalError creates an internal server error (500)
func InternalError(message string) *AppError {
	return NewAppError(
		"INTERNAL_ERROR",
		"An internal error occurred",
		http.StatusInternalServerError,
	)
}
