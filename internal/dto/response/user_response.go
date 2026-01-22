package response

import (
	"github.com/ardianilyas/go-ticketing/internal/domain"
	"github.com/google/uuid"
)

// UserResponse represents a user in API responses
type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// NewUserResponse creates a UserResponse from a domain.User
func NewUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

// NewUserListResponse creates a list of UserResponse from domain.User slice
func NewUserListResponse(users []domain.User) []*UserResponse {
	responses := make([]*UserResponse, len(users))
	for i, user := range users {
		responses[i] = NewUserResponse(&user)
	}
	return responses
}

// SuccessResponse represents a generic success response
type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		Data:    data,
		Message: message,
	}
}
