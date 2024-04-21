package models

type User struct {
	Id           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization string `json:"organization"`
	Position     string `json:"position"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}
