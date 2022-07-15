package inbound_orders

type Inbound_orders struct {
	Id               int    `json:"id"`
	Order_date       string `json:"order_date"`
	Order_number     string `json:"order_number"`
	Employee_id      int    `json:"employee_id"`
	Product_batch_id int    `json:"product_batch_id"`
	Warehouse_id     int    `json:"warehouse_id"`
}

type Employee struct {
	Id             int    `json:"id"`
	Card_number_id int    `json:"card_number_id"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"Last_name"`
	Warehouse_id   int    `json:"warehouse_id"`
}

type Inbound_orders_report struct {
	Employee
	Inbound_orders_count int `json:"inbound_orders_count"`
}

type Orders []Orders

type Warehouse struct {
	Id int `json:"id"`
}

type ProductBatch struct {
	Id int `json:"id"`
}
