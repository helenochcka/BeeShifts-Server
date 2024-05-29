package entities

type Plan struct {
	Id        int    `json:"id"`
	ManagerId int    `json:"manager_id"`
	Name      string `json:"name"`
	Position  int    `json:"position"`
}
