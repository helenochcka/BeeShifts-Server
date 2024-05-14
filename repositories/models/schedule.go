package models

type Schedule struct {
	Id        int    `json:"id"`
	PlanID    int    `json:"plan_id"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}
