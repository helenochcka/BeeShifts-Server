package entities

type Shift struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Date        string `json:"last_name"`
	ScheduleId  int    `json:"schedule_id"`
	ShiftTypeId int    `json:"shift_type_id"`
}
