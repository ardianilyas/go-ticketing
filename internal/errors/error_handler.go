package errors

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ErrorResponse represents the structure of error responses
type ErrorResponse struct {
	Error *AppError `json:"error"`
}

// RespondWithError sends a properly formatted error response
func RespondWithError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.StatusCode, ErrorResponse{Error: appErr})
		return
	}

	// If it's not an AppError, treat it as internal error
	internalErr := InternalError(err.Error())
	c.JSON(internalErr.StatusCode, ErrorResponse{Error: internalErr})
}

// RespondWithValidationError sends a validation error response
func RespondWithValidationError(c *gin.Context, message string, details ...string) {
	// Parse validator errors for better messages
	var formattedDetails []string
	for _, detail := range details {
		formattedDetails = append(formattedDetails, formatValidationDetail(detail))
	}

	err := ValidationError(message, formattedDetails...)
	c.JSON(err.StatusCode, ErrorResponse{Error: err})
}

// formatValidationDetail formats validation error details to be more readable
func formatValidationDetail(detail string) string {
	// Parse gin validator error format
	if strings.Contains(detail, "Error:Field validation for") {
		parts := strings.Split(detail, "'")
		if len(parts) >= 4 {
			field := parts[1]
			tag := parts[3]

			switch tag {
			case "required":
				return fmt.Sprintf("%s is required", field)
			case "email":
				return fmt.Sprintf("%s must be a valid email address", field)
			case "min":
				return fmt.Sprintf("%s is too short", field)
			case "max":
				return fmt.Sprintf("%s is too long", field)
			default:
				return fmt.Sprintf("%s is invalid", field)
			}
		}
	}

	return detail
}

// ParseValidatorError converts validator.ValidationErrors to readable messages
func ParseValidatorError(err error) []string {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		var details []string
		for _, e := range validationErrs {
			details = append(details, formatFieldError(e))
		}
		return details
	}
	return []string{err.Error()}
}

func formatFieldError(err validator.FieldError) string {
	field := err.Field()

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, err.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
