package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
