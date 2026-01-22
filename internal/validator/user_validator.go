package validator

import (
	"github.com/ardianilyas/go-ticketing/internal/dto/request"
	"github.com/ardianilyas/go-ticketing/internal/errors"
)

// ValidateCreateUser validates a CreateUserRequest
func ValidateCreateUser(req *request.CreateUserRequest) error {
	var validationErrors []string

	// Check required fields
	if IsEmpty(req.Name) {
		validationErrors = append(validationErrors, "Name is required")
	}

	if IsEmpty(req.Email) {
		validationErrors = append(validationErrors, "Email is required")
	} else if !IsValidEmail(req.Email) {
		validationErrors = append(validationErrors, "Email must be a valid email address")
	}

	// Additional business rules
	if len(req.Name) > 100 {
		validationErrors = append(validationErrors, "Name must be at most 100 characters")
	}

	if len(validationErrors) > 0 {
		return errors.ValidationError("Validation failed", validationErrors...)
	}

	return nil
}

// ValidateUpdateUser validates an UpdateUserRequest
func ValidateUpdateUser(req *request.UpdateUserRequest) error {
	var validationErrors []string

	// Email is optional, but if provided must be valid
	if !IsEmpty(req.Email) && !IsValidEmail(req.Email) {
		validationErrors = append(validationErrors, "Email must be a valid email address")
	}

	// Name length check if provided
	if !IsEmpty(req.Name) && len(req.Name) > 100 {
		validationErrors = append(validationErrors, "Name must be at most 100 characters")
	}

	if len(validationErrors) > 0 {
		return errors.ValidationError("Validation failed", validationErrors...)
	}

	return nil
}
