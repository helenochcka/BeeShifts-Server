package dtos

type EmployeePreferenceFilterDTO struct {
	Ids          []int `form:"id"`
	ShiftTypeIds []int `form:"shift_type_id"`
	Dates        []int `form:"date"`
}

type EmployeePreferenceDTO struct {
	Id          int    `json:"id"`
	User        string `json:"user"`
	ShiftTypeId int    `json:"shift_type_id"`
	Preference  int    `json:"preference"`
	Date        int    `json:"date"`
}

type CreateEmployeePreferenceDTO struct {
	ShiftTypeId int `json:"shift_type_id"`
	Preference  int `json:"preference"`
	Date        int `json:"date"`
}
