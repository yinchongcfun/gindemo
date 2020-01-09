package models

import (
	"time"
)

type Category struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "category"
}
