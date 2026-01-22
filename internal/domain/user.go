package domain

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey`
	Name string `gorm:"type:varchar(100) not null`
	Email string `gorm:"type:varchar(100) not null unique`
}

func (User) TableName() string {
	return "users"
}