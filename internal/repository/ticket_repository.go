package repository

import (
	"github.com/ardianilyas/go-ticketing/internal/domain"
	"github.com/ardianilyas/go-ticketing/internal/repository/interfaces"
	"gorm.io/gorm"
)

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) interfaces.TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) Create(ticket *domain.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) FindAll(userID string) ([]domain.Ticket, error) {
	var tickets []domain.Ticket
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) FindByID(id, userID string) (*domain.Ticket, error) {
	var ticket domain.Ticket
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) Update(ticket *domain.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) Delete(id, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Ticket{}).Error
}
