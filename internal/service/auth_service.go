package service

import (
	"github.com/ardianilyas/go-ticketing/internal/domain"
	"github.com/ardianilyas/go-ticketing/internal/dto/request"
	"github.com/ardianilyas/go-ticketing/internal/errors"
	"github.com/ardianilyas/go-ticketing/internal/jwt"
	"github.com/ardianilyas/go-ticketing/internal/repository/interfaces"
	"github.com/ardianilyas/go-ticketing/internal/utils"
	"github.com/google/uuid"
)

type AuthService struct {
	userRepo   interfaces.UserRepository
	jwtService *jwt.JWTService
}

func NewAuthService(userRepo interfaces.UserRepository, jwtService *jwt.JWTService) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

// Register creates a new user account
func (s *AuthService) Register(req *request.RegisterRequest) (*domain.User, error) {
	// Check if email already exists
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.ConflictError("Email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.InternalError("Failed to hash password")
	}

	// Create user
	user := &domain.User{
		ID:       uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.InternalError("Failed to create user")
	}

	return user, nil
}

// Login authenticates a user and returns JWT token
func (s *AuthService) Login(req *request.LoginRequest) (*domain.User, string, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, "", errors.ValidationError("Invalid email or password")
	}

	// Check password
	if err := utils.CheckPassword(user.Password, req.Password); err != nil {
		return nil, "", errors.ValidationError("Invalid email or password")
	}

	// Generate JWT token
	token, err := s.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, "", errors.InternalError("Failed to generate token")
	}

	return user, token, nil
}

// ValidateToken validates a JWT token and returns the user
func (s *AuthService) ValidateToken(tokenString string) (*domain.User, error) {
	// Validate token
	claims, err := s.jwtService.ValidateToken(tokenString)
	if err != nil {
		return nil, errors.ValidationError("Invalid or expired token")
	}

	// Get user by ID
	user, err := s.userRepo.FindByID(claims.UserID.String())
	if err != nil {
		return nil, errors.NotFoundError("User")
	}

	return user, nil
}
