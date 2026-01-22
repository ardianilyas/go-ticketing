package response

import (
	"time"

	"github.com/ardianilyas/go-ticketing/internal/domain"
	"github.com/google/uuid"
)

// TicketResponse represents a ticket in API responses
type TicketResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewTicketResponse creates a TicketResponse from domain.Ticket
func NewTicketResponse(ticket *domain.Ticket) *TicketResponse {
	return &TicketResponse{
		ID:          ticket.ID,
		Title:       ticket.Title,
		Description: ticket.Description,
		Status:      string(ticket.Status),
		Priority:    string(ticket.Priority),
		UserID:      ticket.UserID,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}
}

// NewTicketListResponse creates a list of TicketResponse
func NewTicketListResponse(tickets []domain.Ticket) []*TicketResponse {
	responses := make([]*TicketResponse, len(tickets))
	for i, ticket := range tickets {
		responses[i] = NewTicketResponse(&ticket)
	}
	return responses
}
