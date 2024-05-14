package models

type User struct {
	Id             int
	OrganizationId *int
	PositionId     *int
	Role           string
	FirstName      string
	LastName       string
	Email          string
	Password       string
}
