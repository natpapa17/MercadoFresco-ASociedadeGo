package usecases

import "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/domain"

type Repository interface {
	Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error)
	GetAll() (domain.Warehouses, error)
	GetById(id int) (domain.Warehouse, error)
	UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error)
	DeleteById(id int) error
	GetByWarehouseCode(code string) (domain.Warehouse, error)
}
