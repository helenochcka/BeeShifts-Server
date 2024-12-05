package positions

type Entity struct {
	Id        int    `json:"id"`
	ManagerId int    `json:"manager_id"`
	Name      string `json:"name"`
}
