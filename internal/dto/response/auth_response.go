package response

// AuthResponse represents authentication response
type AuthResponse struct {
	User    *UserResponse `json:"user"`
	Message string        `json:"message"`
}

// NewAuthResponse creates a new AuthResponse
func NewAuthResponse(user *UserResponse, message string) *AuthResponse {
	return &AuthResponse{
		User:    user,
		Message: message,
	}
}
