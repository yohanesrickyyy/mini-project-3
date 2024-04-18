package dto

type ActivityLog struct {
	ActivityLogId int    `json:"activity_log_id"`
	ActivityId    int    `json:"activity_id"`
	UserId        int    `json:"user_id"`
	ActivityName  string `json:"activity_name"`
}
