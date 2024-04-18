package entity

import "time"

type Booking struct {
	BookingID   int       `json:"booking_id"`
	UserID      int       `json:"user_id"`
	EquipmentID int       `json:"equipment_id"`
	BookingDate time.Time `json:"booking_date"`
	ReturnDate  time.Time `json:"return_date"`
}
