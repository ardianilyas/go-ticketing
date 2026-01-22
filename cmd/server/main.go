package main

import (
	"log"

	"github.com/ardianilyas/go-ticketing/internal/config"
	"github.com/ardianilyas/go-ticketing/internal/database"
	"github.com/ardianilyas/go-ticketing/internal/domain"
	"github.com/ardianilyas/go-ticketing/internal/handler"
	"github.com/ardianilyas/go-ticketing/internal/jwt"
	"github.com/ardianilyas/go-ticketing/internal/repository"
	"github.com/ardianilyas/go-ticketing/internal/routes"
	"github.com/ardianilyas/go-ticketing/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := database.Connect()
	db.AutoMigrate(&domain.User{}, &domain.Ticket{})

	// Initialize JWT service
	jwtService := jwt.NewJWTService()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	ticketRepo := repository.NewTicketRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, jwtService)
	ticketService := service.NewTicketService(ticketRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService, jwtService)
	ticketHandler := handler.NewTicketHandler(ticketService)

	// Setup routes
	r := gin.Default()
	routes.Register(r, userHandler, authHandler, ticketHandler, jwtService)

	log.Fatal(r.Run(":" + config.Get("APP_PORT")))
}
