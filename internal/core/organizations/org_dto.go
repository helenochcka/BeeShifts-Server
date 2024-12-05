package organizations

type FilterDTO struct {
	Ids   []int    `form:"id"`
	Names []string `form:"name"`
}
