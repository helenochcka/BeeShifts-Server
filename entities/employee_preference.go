package entities

type EmployeePreferenceEntity struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Preference  int    `json:"preference"`
	ShiftTypeId int    `json:"shift_type_id"`
	Date        string `json:"date"`
}
