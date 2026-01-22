package service

import (
	"github.com/ardianilyas/go-ticketing/internal/domain"
	"github.com/ardianilyas/go-ticketing/internal/dto/request"
	"github.com/ardianilyas/go-ticketing/internal/errors"
	"github.com/ardianilyas/go-ticketing/internal/repository/interfaces"
	"github.com/google/uuid"
)

type TicketService struct {
	repo interfaces.TicketRepository
}

func NewTicketService(repo interfaces.TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}

// CreateTicket creates a new ticket for the user
func (s *TicketService) CreateTicket(userID string, req *request.CreateTicketRequest) (*domain.Ticket, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.ValidationError("Invalid user ID")
	}

	ticket := &domain.Ticket{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		Status:      domain.StatusOpen,
		Priority:    domain.TicketPriority(req.Priority),
		UserID:      userUUID,
	}

	if err := s.repo.Create(ticket); err != nil {
		return nil, errors.InternalError("Failed to create ticket")
	}

	return ticket, nil
}

// FindAll returns all tickets for a user
func (s *TicketService) FindAll(userID string) ([]domain.Ticket, error) {
	tickets, err := s.repo.FindAll(userID)
	if err != nil {
		return nil, errors.InternalError("Failed to fetch tickets")
	}
	return tickets, nil
}

// FindByID returns a ticket by ID for a user
func (s *TicketService) FindByID(id, userID string) (*domain.Ticket, error) {
	ticket, err := s.repo.FindByID(id, userID)
	if err != nil {
		return nil, errors.NotFoundError("Ticket")
	}
	return ticket, nil
}

// UpdateTicket updates an existing ticket
func (s *TicketService) UpdateTicket(id, userID string, req *request.UpdateTicketRequest) (*domain.Ticket, error) {
	ticket, err := s.repo.FindByID(id, userID)
	if err != nil {
		return nil, errors.NotFoundError("Ticket")
	}

	// Update fields if provided
	if req.Title != "" {
		ticket.Title = req.Title
	}
	if req.Description != "" {
		ticket.Description = req.Description
	}
	if req.Status != "" {
		ticket.Status = domain.TicketStatus(req.Status)
	}
	if req.Priority != "" {
		ticket.Priority = domain.TicketPriority(req.Priority)
	}

	if err := s.repo.Update(ticket); err != nil {
		return nil, errors.InternalError("Failed to update ticket")
	}

	return ticket, nil
}

// DeleteTicket deletes a ticket
func (s *TicketService) DeleteTicket(id, userID string) error {
	// Check if ticket exists
	_, err := s.repo.FindByID(id, userID)
	if err != nil {
		return errors.NotFoundError("Ticket")
	}

	if err := s.repo.Delete(id, userID); err != nil {
		return errors.InternalError("Failed to delete ticket")
	}

	return nil
}
