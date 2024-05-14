package models

type Shift struct {
	Id          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Date        string `json:"last_name"`
	ScheduleID  int    `json:"schedule_id"`
	ShiftTypeID int    `json:"shift_type_id"`
}
