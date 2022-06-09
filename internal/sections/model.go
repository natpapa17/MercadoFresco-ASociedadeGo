package sections

type Section struct {
	ID                  int     `json:"id"`
	SectionNumber       int     `json:"section_number" binding:"required"`
	CurrentTemperature  float32 `json:"current_temperature" binding:"required"`
	MinimumTemprarature float32 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity     int     `json:"current_capacity" binding:"required"`
	MinimumCapacity     int     `json:"minimum_capacity" binding:"required"`
	MaximumCapacity     int     `json:"maximum_capacity" binding:"required"`
	WarehouseID         int     `json:"warehouse_id" binding:"required"`
	ProductTypeID       int     `json:"product_type_id" binding:"required"`
}
