package shifts

type ShiftTypeEntity struct {
	Id                int    `json:"id"`
	NumberOfEmployees int    `json:"number_of_employees"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	PlanId            int    `json:"plan_id"`
}
