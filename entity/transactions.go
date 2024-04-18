package entity

import "time"

type Transaction struct {
	TransactionID   int       `json:"transaction_id"`
	RentalID        int       `json:"rental_id"`
	TransactionDate time.Time `json:"transaction_date"`
	Amount          float64   `json:"amount"`
	PaymentMethod   string    `json:"payment_method"`
}
