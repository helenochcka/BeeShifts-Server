package models

type ShiftType struct {
	Id                int    `json:"id"`
	NumberOfEmployees int    `json:"number_of_employees"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	Plan              string `json:"plan"`
}
