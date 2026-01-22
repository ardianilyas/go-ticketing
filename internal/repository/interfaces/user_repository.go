package interfaces

import "github.com/ardianilyas/go-ticketing/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}
