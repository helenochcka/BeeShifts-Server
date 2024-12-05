package schedules

type PlanEntity struct {
	Id        int    `json:"id"`
	ManagerId int    `json:"manager_id"`
	Name      string `json:"name"`
	Position  int    `json:"positions"`
}
