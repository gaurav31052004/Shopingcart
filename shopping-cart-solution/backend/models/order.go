package models

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CartID    uint      `json:"cart_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
