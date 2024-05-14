package dtos

type GetOrganizationsDTO struct {
	Ids   []int    `form:"id"`
	Names []string `form:"name"`
}
