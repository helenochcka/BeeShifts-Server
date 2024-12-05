package shifts

type ShiftFilterDTO struct {
	Ids     []int `form:"id"`
	UserIds []int `form:"user_id"`
	Date    []int `from:"date"`
}

type ShiftDTO struct {
	Id          int    `json:"id"`
	User        string `json:"users"`
	Schedule    string `json:"schedules"`
	ShiftTypeId int    `json:"shift_type_id"`
	Date        int    `json:"date"`
}

type CreateShiftDTO struct {
	ScheduleId  int `json:"schedule_id"`
	ShiftTypeId int `json:"shift_type_id"`
	Date        int `json:"date"`
}

type UpdateShiftTDTO struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
}
