package entity

import "time"

type Order struct {
	OrderID     uint      `json:"order_id"`
	CustomerID  uint      `json:"customer_id"`
	OrderDate   time.Time `json:"order_date"`
	TotalAmount float64   `json:"total_amount"`
}
