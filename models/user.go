package models

type User struct {
	Id           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization int    `json:"organization"`
	Position     int    `json:"position"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}
