package employee

type Employee struct {
	Id             int    `json:"id"`
	Card_number_id int    `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"Last_name"`
	Warehouse_id   int    `json:"warehouse_id"`
}
