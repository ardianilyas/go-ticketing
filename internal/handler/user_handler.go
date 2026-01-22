package handler

import (
	"net/http"

	"github.com/ardianilyas/go-ticketing/internal/dto/request"
	"github.com/ardianilyas/go-ticketing/internal/dto/response"
	"github.com/ardianilyas/go-ticketing/internal/errors"
	"github.com/ardianilyas/go-ticketing/internal/service"
	"github.com/ardianilyas/go-ticketing/internal/validator"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}

// Create handles user creation requests
func (h *UserHandler) Create(c *gin.Context) {
	var req request.CreateUserRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.RespondWithValidationError(c, "Invalid request body", err.Error())
		return
	}

	// Validate request
	if err := validator.ValidateCreateUser(&req); err != nil {
		errors.RespondWithError(c, err)
		return
	}

	// Create user via service
	user, err := h.service.CreateUser(&req)
	if err != nil {
		errors.RespondWithError(c, err)
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, response.NewSuccessResponse(
		response.NewUserResponse(user),
		"User created successfully",
	))
}

// FindAll handles listing all users
func (h *UserHandler) FindAll(c *gin.Context) {
	users, err := h.service.FindAll()
	if err != nil {
		errors.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(
		response.NewUserListResponse(users),
		"",
	))
}
