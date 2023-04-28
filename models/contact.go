package models

import (
	"time"
)

type Contact struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
