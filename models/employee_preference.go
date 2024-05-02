package models

type EmployeePreference struct {
	Id          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Preference  int    `json:"preference"`
	ShiftTypeID int    `json:"shift_type_id"`
	Date        string `json:"date"`
}
