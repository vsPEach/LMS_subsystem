package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Email    string `json:"email" gorm:"primarykey"`
	Password string `json:"password"`
	Group    string `json:"group"`
	Role     string `json:"role"`
}
