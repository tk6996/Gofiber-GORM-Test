package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Employee_id uint32    `json:"emp_id"  validate:"required"`
	Name        string    `json:"name" validate:"required,min=3,max=50"`
	Lastname    string    `json:"lastname" validate:"required,min=3,max=50"`
	Birthday    time.Time `json:"birthday" validate:"required"`
	Age         uint8     `json:"age" validate:"required,gte=0"`
	Email       string    `json:"email" validate:"required,email"`
	Tel         string    `json:"tel" validate:"required"`
}
