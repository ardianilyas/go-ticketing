package routes

import (
	"github.com/ardianilyas/go-ticketing/internal/handler"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, userHandler *handler.UserHandler) {
	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.FindAll)
		}
	}
}