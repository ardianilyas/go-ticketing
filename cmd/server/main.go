package main

import (
	"log"

	"github.com/ardianilyas/go-ticketing/internal/config"
	"github.com/ardianilyas/go-ticketing/internal/database"
	"github.com/ardianilyas/go-ticketing/internal/domain"
	"github.com/ardianilyas/go-ticketing/internal/handler"
	"github.com/ardianilyas/go-ticketing/internal/repository"
	"github.com/ardianilyas/go-ticketing/internal/routes"
	"github.com/ardianilyas/go-ticketing/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := database.Connect()
	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	routes.Register(r, userHandler)

	log.Fatal(r.Run(":" + config.Get("APP_PORT")))
}