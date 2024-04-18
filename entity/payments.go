package entity

import "time"

type Payment struct {
	PaymentID     int       `json:"payment_id"`
	UserID        int       `json:"user_id"`
	TransactionID int       `json:"transaction_id"`
	PaymentDate   time.Time `json:"payment_date"`
	Amount        float64   `json:"amount"`
}
