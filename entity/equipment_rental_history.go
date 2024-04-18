package entity

import "time"

type EquipmentRentalHistory struct {
	RentalID    uint      `gorm:"primaryKey"`
	UserID      string    `json:"userid"`
	EquipmentID uint      `json:"equipmentid"`
	RentalDate  time.Time `json:"rentaldate"`
	ReturnDate  time.Time `json:"returndate"`
	RentalCost  float64   `json:"rentalcost"`
	Status      string    `json:"status"`
}
