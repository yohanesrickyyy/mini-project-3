package entity

type UserActivity struct {
	UserId       int    `json:"user_id"`
	ActivityId   int    `json:"activity_id"`
	Description string `json:"description"`
}
