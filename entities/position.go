package entities

type PositionEntity struct {
	Id        int    `json:"id"`
	ManagerId int    `json:"manager_id"`
	Name      string `json:"name"`
}
