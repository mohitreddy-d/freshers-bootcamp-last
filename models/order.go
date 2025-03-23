package models

import "time"

type Order struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID uint      `json:"customer_id"`
	ProductID  uint      `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Status     string    `json:"status"` // "order placed", "processed", "failed"
	CreatedAt  time.Time `json:"created_at"`
}
