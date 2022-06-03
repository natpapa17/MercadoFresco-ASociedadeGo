package sections

type Section struct {
	ID                  int     `json:"id"`
	SectionNumber       int     `json:"section_number"`
	CurrentTemperature  float32 `json:"current_temperature"`
	MinimumTemprarature float32 `json:"minimum_temperature"`
	CurrentCapacity     int     `json:"current_capacity"`
	MinimumCapacity     int     `json:"minimum_capacity"`
	MaximumCapacity     int     `json:"maximum_capacity"`
	WarehouseID         int     `json:"warehouse_id"`
	ProductTypeID       int     `json:"product_type_id"`
}
