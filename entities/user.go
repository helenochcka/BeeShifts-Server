package entities

type UserEntity struct {
	Id             int    `json:"id"`
	OrganizationId *int   `json:"organization_id"`
	PositionId     *int   `json:"position_id"`
	Role           string `json:"role"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}
