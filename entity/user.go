package entity

type User struct {
	UserID        int     `json:"userid" gorm:"primaryKey:autoIncrement"`
	Email         string  `json:"email" gorm:"unique"`
	Password      string  `json:"password,omitempty"`
	DepositAmount float64 `json:"deposit"`
}
