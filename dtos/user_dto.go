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
	Id           int
	Organization *string
	Position     *string
	Role         string
	FirstName    string
	LastName     string
	Email        string
	Password     string
}

type CreateUserDTO struct {
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserDTO struct {
	Id           int
	Organization *int
	Position     *int
	Role         string
	FirstName    string
	LastName     string
	Email        string
	Password     string
}

type UpdateSelfUserDTO struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AttachEmployeeDTO struct {
	Id           int  `json:"id"`
	Organization *int `json:"organization_id"`
	Position     *int `json:"position_id"`
}
