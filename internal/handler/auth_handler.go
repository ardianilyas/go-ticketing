package handler

import (
	"net/http"

	"github.com/ardianilyas/go-ticketing/internal/config"
	"github.com/ardianilyas/go-ticketing/internal/dto/request"
	"github.com/ardianilyas/go-ticketing/internal/dto/response"
	"github.com/ardianilyas/go-ticketing/internal/errors"
	"github.com/ardianilyas/go-ticketing/internal/jwt"
	"github.com/ardianilyas/go-ticketing/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
	jwtService  *jwt.JWTService
}

func NewAuthHandler(authService *service.AuthService, jwtService *jwt.JWTService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		jwtService:  jwtService,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errors.RespondWithValidationError(c, "Invalid request body", err.Error())
		return
	}

	// Validate password length
	if len(req.Password) < 8 {
		errors.RespondWithValidationError(c, "Validation failed", "Password must be at least 8 characters")
		return
	}

	// Register user
	user, err := h.authService.Register(&req)
	if err != nil {
		errors.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse(
		response.NewUserResponse(user),
		"User registered successfully",
	))
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errors.RespondWithValidationError(c, "Invalid request body", err.Error())
		return
	}

	// Login user
	user, token, err := h.authService.Login(&req)
	if err != nil {
		errors.RespondWithError(c, err)
		return
	}

	// Set HTTP-only cookie
	maxAge := int(h.jwtService.GetExpiresIn().Seconds())
	domain := config.Get("COOKIE_DOMAIN")

	c.SetCookie(
		"access_token", // name
		token,          // value
		maxAge,         // max age in seconds
		"/",            // path
		domain,         // domain
		false,          // secure (set to true in production with HTTPS)
		true,           // httpOnly
	)

	c.JSON(http.StatusOK, response.NewAuthResponse(
		response.NewUserResponse(user),
		"Login successful",
	))
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *gin.Context) {
	domain := config.Get("COOKIE_DOMAIN")

	// Clear cookie by setting maxAge to -1
	c.SetCookie(
		"access_token",
		"",
		-1,
		"/",
		domain,
		false,
		true,
	)

	c.JSON(http.StatusOK, response.NewSuccessResponse(
		nil,
		"Logout successful",
	))
}

// Me returns the current authenticated user
func (h *AuthHandler) Me(c *gin.Context) {
	// Get user from context (set by auth middleware)
	userVal, exists := c.Get("user")
	if !exists {
		errors.RespondWithError(c, errors.ValidationError("Unauthorized"))
		return
	}

	user := userVal.(map[string]interface{})
	c.JSON(http.StatusOK, response.NewSuccessResponse(user, ""))
}
