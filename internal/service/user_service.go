package service

import (
	"github.com/ardianilyas/go-ticketing/internal/domain"
	"github.com/ardianilyas/go-ticketing/internal/dto/request"
	"github.com/ardianilyas/go-ticketing/internal/errors"
	"github.com/ardianilyas/go-ticketing/internal/repository/interfaces"
	"github.com/google/uuid"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo}
}

// CreateUser creates a new user from the request DTO
func (s *UserService) CreateUser(req *request.CreateUserRequest) (*domain.User, error) {
	// Check if email already exists
	existingUser, _ := s.repo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.ConflictError("Email already exists")
	}

	user := &domain.User{
		ID:    uuid.New(),
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, errors.InternalError("Failed to create user")
	}

	return user, nil
}

// FindAll returns all users
func (s *UserService) FindAll() ([]domain.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, errors.InternalError("Failed to fetch users")
	}
	return users, nil
}
