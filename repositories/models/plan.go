package models

type Plan struct {
	Id        int    `json:"id"`
	ManagerID int    `json:"manager_id"`
	Name      string `json:"name"`
	Position  int    `json:"position"`
}
