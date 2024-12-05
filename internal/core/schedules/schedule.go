package schedules

type ScheduleEntity struct {
	Id        int    `json:"id"`
	PlanId    int    `json:"plan_id"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}
