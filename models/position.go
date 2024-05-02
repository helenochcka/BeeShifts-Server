package models

type Position struct {
	Id        int    `json:"id"`
	ManagerID int    `json:"manager_id"`
	Name      string `json:"name"`
}
