package warehouses

type Warehouse struct {
	Id                 int
	WarehouseCode      string
	Address            string
	Telephone          string
	MinimumCapacity    int
	MinimumTemperature float64
}
