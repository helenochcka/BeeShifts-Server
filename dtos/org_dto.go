package dtos

type OrgsFilterDTO struct {
	Ids   []int    `form:"id"`
	Names []string `form:"name"`
}
