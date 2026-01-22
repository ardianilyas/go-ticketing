package request

// CreateTicketRequest represents the request to create a ticket
type CreateTicketRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Priority    string `json:"priority" binding:"required,oneof=low medium high urgent"`
}

// UpdateTicketRequest represents the request to update a ticket
type UpdateTicketRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty" binding:"omitempty,oneof=open in_progress resolved closed"`
	Priority    string `json:"priority,omitempty" binding:"omitempty,oneof=low medium high urgent"`
}
