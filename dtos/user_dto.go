package dtos

type UsersFilterDTO struct {
	Ids             []int    `form:"id"`
	OrganizationIds []int    `form:"organization_id"`
	PositionIds     []int    `form:"position_id"`
	FirstNames      []string `form:"first_name"`
	LastNames       []string `form:"last_name"`
	Emails          []string `form:"email"`
}

type UserDTO struct {
	Id           int     `json:"id"`
	Organization *string `json:"organization"`
	Position     *string `json:"position"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
}

type CreateUserDTO struct {
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateSelfUserDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AttachUserDTO struct {
	Id             int  `json:"id"`
	OrganizationId *int `json:"organization_id"`
	PositionId     *int `json:"position_id"`
}
