package routes

import (
	"github.com/ardianilyas/go-ticketing/internal/handler"
	"github.com/ardianilyas/go-ticketing/internal/jwt"
	"github.com/ardianilyas/go-ticketing/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, userHandler *handler.UserHandler, authHandler *handler.AuthHandler, ticketHandler *handler.TicketHandler, jwtService *jwt.JWTService) {
	api := r.Group("/api")

	// Public auth routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		// Auth endpoints requiring authentication
		protected.GET("/auth/me", authHandler.Me)

		// User endpoints
		users := protected.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.FindAll)
		}

		// Ticket endpoints
		tickets := protected.Group("/tickets")
		{
			tickets.POST("", ticketHandler.Create)
			tickets.GET("", ticketHandler.FindAll)
			tickets.GET("/:id", ticketHandler.FindByID)
			tickets.PUT("/:id", ticketHandler.Update)
			tickets.DELETE("/:id", ticketHandler.Delete)
		}
	}
}
