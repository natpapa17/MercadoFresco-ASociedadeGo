package employee

type Employee struct {
	Id             int    `json:"id"`
	Card_number_id int    `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"Last_name"`
	Warehouse_id   int    `json:"warehouse_id"`
}

type Employees []Employee

type Warehouse struct {
	Id                 int     `json:"id"`
	WarehouseCode      string  `json:"warehouse_code"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimumTemperature float64 `json:"minimum_temperature"`
}
