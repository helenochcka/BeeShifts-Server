package dtos

type ShiftTypeFilterDTO struct {
	Ids     []int `form:"id"`
	PlanIds []int `form:"plan_id"`
}

type ShiftTypeDTO struct {
	Id                int    `json:"id"`
	NumberOfEmployees int    `json:"number_of_employees"`
	Plan              string `json:"plan"`
	StartTime         int    `json:"start_time"`
	EndTime           int    `json:"end_time"`
}

type CreateShiftTypeDTO struct {
	NumberOfEmployees int `json:"number_of_employees"`
	PlanId            int `json:"plan_id"`
	StartTime         int `json:"start_time"`
	EndTime           int `json:"end_time"`
}

type UpdateShiftTypeDTO struct {
	Id                int `json:"id"`
	NumberOfEmployees int `json:"number_of_employees"`
	PlanId            int `json:"plan_id"`
	StartTime         int `json:"start_time"`
	EndTime           int `json:"end_time"`
}
