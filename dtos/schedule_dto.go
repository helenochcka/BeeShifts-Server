package dtos

type ScheduleFilterDTO struct {
	Ids       []int `form:"id"`
	PlanIds   []int `form:"plan_id"`
	StartedAt []int `form:"started_at"`
	EndedAt   []int `form:"ended_at"`
}

type ScheduleDTO struct {
	Id        int    `json:"id"`
	Plan      string `json:"plan"`
	StartedAt int    `json:"started_at"`
	EndedAt   int    `json:"ended_at"`
}

type CreateScheduleDTO struct {
	PlanId    int `json:"plan_id"`
	StartedAt int `json:"started_at"`
	EndedAt   int `json:"ended_at"`
}

type UpdateScheduleDTO struct {
	Id        int `json:"id"`
	PlanId    int `json:"plan_id"`
	StartedAt int `json:"started_at"`
	EndedAt   int `json:"ended_at"`
}
