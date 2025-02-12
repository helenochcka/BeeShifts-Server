package users

type FilterDTO struct {
	Ids             []int    `form:"id"`
	OrganizationIds []int    `form:"organization_id"` //TODO remove filter by organization_id
	PositionIds     []int    `form:"position_id"`
	FirstNames      []string `form:"first_name"`
	LastNames       []string `form:"last_name"`
	Emails          []string `form:"email"`
}

type ViewDTO struct {
	Id           int     `json:"id"`
	Organization *string `json:"organizations"`
	Position     *string `json:"positions"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
}

type CreateDTO struct {
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateSelfDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AttachDTO struct {
	Id         int `json:"id"`
	PositionId int `json:"position_id"`
}

type DetachDTO struct {
	Id int `json:"id"`
}

type CredsDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	AccessToken string `json:"access_token"`
}

type TokenPayloadDTO struct {
	Id        int
	ExpiresAt int64
}
