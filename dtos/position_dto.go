package dtos

type GetPositionsDTO struct {
	Ids        []int    `form:"id"`
	ManagerIds []int    `form:"manager_id"`
	Names      []string `form:"name"`
}
