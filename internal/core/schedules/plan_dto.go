package schedules

type PlanFilterDTO struct {
	Ids         []int    `form:"id"`
	PositionIds []int    `form:"position_id"`
	Names       []string `form:"name"`
}

type PlanDTO struct {
	Id       int    `json:"id"`
	Position string `json:"positions"`
	Name     string `json:"name"`
}

type CreatePlanDTO struct {
	PositionId int    `json:"position_id"`
	Name       string `json:"name"`
}

type UpdatePlanDTO struct {
	Id         int    `json:"id"`
	PositionId int    `json:"position_id"`
	Name       string `json:"name"`
}
