package database

import (
	"fmt"
	"log"

	"github.com/ardianilyas/go-ticketing/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect()  *gorm.DB{
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Get("DB_HOST"),
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"),
		config.Get("DB_PORT"),
	)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	log.Println("connected to database")
	return db
}