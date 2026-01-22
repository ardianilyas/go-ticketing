package interfaces

import "github.com/ardianilyas/go-ticketing/internal/domain"

type TicketRepository interface {
	Create(ticket *domain.Ticket) error
	FindAll(userID string) ([]domain.Ticket, error)
	FindByID(id, userID string) (*domain.Ticket, error)
	Update(ticket *domain.Ticket) error
	Delete(id, userID string) error
}
