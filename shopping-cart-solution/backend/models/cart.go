package models

import "time"

type Cart struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	Name      string      `json:"name"`
	Status    string      `json:"status"` // OPEN or ORDERED
	Items     []CartItem  `json:"items"`
	CreatedAt time.Time   `json:"created_at"`
}
