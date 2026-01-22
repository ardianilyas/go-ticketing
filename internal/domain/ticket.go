package domain

import (
	"time"

	"github.com/google/uuid"
)

// TicketStatus represents the status of a ticket
type TicketStatus string

const (
	StatusOpen       TicketStatus = "open"
	StatusInProgress TicketStatus = "in_progress"
	StatusResolved   TicketStatus = "resolved"
	StatusClosed     TicketStatus = "closed"
)

// IsValid checks if the status is valid
func (s TicketStatus) IsValid() bool {
	switch s {
	case StatusOpen, StatusInProgress, StatusResolved, StatusClosed:
		return true
	}
	return false
}

// TicketPriority represents the priority of a ticket
type TicketPriority string

const (
	PriorityLow    TicketPriority = "low"
	PriorityMedium TicketPriority = "medium"
	PriorityHigh   TicketPriority = "high"
	PriorityUrgent TicketPriority = "urgent"
)

// IsValid checks if the priority is valid
func (p TicketPriority) IsValid() bool {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh, PriorityUrgent:
		return true
	}
	return false
}

// Ticket represents a support ticket
type Ticket struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Title       string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text"`
	Status      TicketStatus   `gorm:"type:varchar(20);not null;default:'open'"`
	Priority    TicketPriority `gorm:"type:varchar(20);not null;default:'medium'"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null"`
	User        User           `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
}

func (Ticket) TableName() string {
	return "tickets"
}
