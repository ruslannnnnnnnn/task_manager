package entity

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskDTO struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
