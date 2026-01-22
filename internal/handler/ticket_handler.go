package handler

import (
	"net/http"

	"github.com/ardianilyas/go-ticketing/internal/dto/request"
	"github.com/ardianilyas/go-ticketing/internal/dto/response"
	"github.com/ardianilyas/go-ticketing/internal/errors"
	"github.com/ardianilyas/go-ticketing/internal/service"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service *service.TicketService
}

func NewTicketHandler(service *service.TicketService) *TicketHandler {
	return &TicketHandler{service: service}
}

// Create handles ticket creation
func (h *TicketHandler) Create(c *gin.Context) {
	var req request.CreateTicketRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errors.RespondWithValidationError(c, "Invalid request body", err.Error())
		return
	}

	userID, _ := c.Get("userId")

	ticket, err := h.service.CreateTicket(userID.(string), &req)
	if err != nil {
		errors.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse(
		response.NewTicketResponse(ticket),
		"Ticket created successfully",
	))
}

// FindAll handles listing all tickets for a user
func (h *TicketHandler) FindAll(c *gin.Context) {
	userID, _ := c.Get("userId")

	tickets, err := h.service.FindAll(userID.(string))
	if err != nil {
		errors.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(
		response.NewTicketListResponse(tickets),
		"",
	))
}

// FindByID handles getting a single ticket
func (h *TicketHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userId")

	ticket, err := h.service.FindByID(id, userID.(string))
	if err != nil {
		errors.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(
		response.NewTicketResponse(ticket),
		"",
	))
}

// Update handles ticket updates
func (h *TicketHandler) Update(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userId")

	var req request.UpdateTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.RespondWithValidationError(c, "Invalid request body", err.Error())
		return
	}

	ticket, err := h.service.UpdateTicket(id, userID.(string), &req)
	if err != nil {
		errors.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(
		response.NewTicketResponse(ticket),
		"Ticket updated successfully",
	))
}

// Delete handles ticket deletion
func (h *TicketHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userId")

	if err := h.service.DeleteTicket(id, userID.(string)); err != nil {
		errors.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse(
		nil,
		"Ticket deleted successfully",
	))
}
