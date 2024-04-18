package entity

type EquipmentType struct {
	TypeID       int     `json:"typeID"`
	Name         string  `json:"name"`
	Availability bool    `json:"availability"`
	RentalCosts  float64 `json:"rentalCosts"`
	Category     string  `json:"category"`
}
