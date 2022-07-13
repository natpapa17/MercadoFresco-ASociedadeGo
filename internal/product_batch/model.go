package product_batch

type ProductBatch struct {
	ID                 int    `json:"id"`
	BatchNumber        int    `json:"batch_number"`
	CurrentQuantity    int    `json:"current_quantity"`
	CurrentTemperature int    `json:"current_temperature"`
	DueDate            string `json:"due_date"`
	InitialQuantity    int    `json:"initial_quantity"`
	ManufacturingDate  string `json:"manufacturing_date"`
	ManufacturingHour  int    `json:"manufacturing_hour"`
	MinimumTemperature int    `json:"minimum_temperature"`
	ProductID          int    `json:"product_id"`
	SectionID          int    `json:"section_id"`
}

type ProductsReport struct {
	SectionID     int `json:"section_id""`
	SectionNumber int `json:"section_number"`
	ProductsCount int `json:"products_count"`
}
