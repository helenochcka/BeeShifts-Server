package dtos

type PositionsFilterDTO struct {
	Ids        []int    `form:"id"`
	ManagerIds []int    `form:"manager_id"`
	Names      []string `form:"name"`
}

type UpdatePositionDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreatePositionDTO struct {
	Name string `json:"name"`
}
