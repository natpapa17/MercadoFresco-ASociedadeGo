package employee

type Employee struct {
	Id             int    `json:"id"`
	card_number_id string `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"Last_name"`
	warehouse_id   int    `json:"warehouse_id"`
}
