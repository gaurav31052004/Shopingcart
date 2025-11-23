package models

type CartItem struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	CartID uint `json:"cart_id"`
	ItemID uint `json:"item_id"`
}
