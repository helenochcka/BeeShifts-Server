package positions

type FilterDTO struct {
	Ids        []int    `form:"id"`
	ManagerIds []int    `from:"manager_id"`
	Names      []string `form:"name"`
}

type GetDTO struct {
	Ids   []int    `form:"id"`
	Names []string `form:"name"`
}

type UpdateDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateDTO struct {
	Name string `json:"name"`
}
