package entity

type User struct {
	UserID        int     `json:"userID"`
	EmailUsername string  `json:"emailUsername"`
	Password      string  `json:"password,omitempty"`
	DepositAmount float64 `json:"depositAmount"`
}
