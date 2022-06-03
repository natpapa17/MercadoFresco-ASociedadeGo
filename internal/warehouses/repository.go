package warehouses

import "errors"

var ws []Warehouse = []Warehouse{}

type Repository interface {
	Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error)
	GetAll() ([]Warehouse, error)
	GetById(id int) (Warehouse, error)
	UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error)
}

type repository struct{}

func CreateRepository() Repository {
	return &repository{}
}

func (r *repository) Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error) {
	w := Warehouse{
		Id:                 1,
		WarehouseCode:      warehouseCode,
		Address:            address,
		Telephone:          telephone,
		MinimumCapacity:    minimumCapacity,
		MinimumTemperature: minimumTemperature,
	}

	ws = append(ws, w)

	return w, nil
}

func (r *repository) GetAll() ([]Warehouse, error) {
	return ws, nil
}

func (r *repository) GetById(id int) (Warehouse, error) {
	result, found := Warehouse{}, false
	for _, w := range ws {
		if w.Id == id {
			result, found = w, true
			break
		}
	}

	if !found {
		return Warehouse{}, errors.New("can't find element with this id")
	}

	return result, nil
}

func (r *repository) UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error) {
	result, found := Warehouse{}, false
	for i, w := range ws {
		if w.Id == id {
			ws[i], found = Warehouse{
				Id:                 id,
				WarehouseCode:      warehouseCode,
				Address:            address,
				Telephone:          telephone,
				MinimumCapacity:    minimumCapacity,
				MinimumTemperature: minimumTemperature,
			}, true
			result = ws[i]
			break
		}
	}

	if !found {
		return Warehouse{}, errors.New("can't find element with this id")
	}

	return result, nil
}
