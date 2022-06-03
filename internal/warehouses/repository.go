package warehouses

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

type Repository interface {
	Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error)
	GetAll() ([]Warehouse, error)
	GetById(id int) (Warehouse, error)
	UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error)
	DeleteById(id int) error
	GetByWarehouseCode(code string) (Warehouse, error)
}

type repository struct {
	file store.Store
}

func CreateRepository(file store.Store) Repository {
	return &repository{
		file: file,
	}
}

func (r *repository) Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error) {
	var ws []Warehouse
	if err := r.file.Read(&ws); err != nil {
		return Warehouse{}, err
	}
	w := Warehouse{
		Id:                 1,
		WarehouseCode:      warehouseCode,
		Address:            address,
		Telephone:          telephone,
		MinimumCapacity:    minimumCapacity,
		MinimumTemperature: minimumTemperature,
	}
	ws = append(ws, w)

	if err := r.file.Write(ws); err != nil {
		return Warehouse{}, err
	}
	return w, nil
}

func (r *repository) GetAll() ([]Warehouse, error) {
	var ws []Warehouse
	if err := r.file.Read(&ws); err != nil {
		return []Warehouse{}, nil
	}
	return ws, nil
}

func (r *repository) GetById(id int) (Warehouse, error) {
	var ws []Warehouse
	if err := r.file.Read(&ws); err != nil {
		return Warehouse{}, nil
	}

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

func (r *repository) GetByWarehouseCode(code string) (Warehouse, error) {
	var ws []Warehouse
	if err := r.file.Read(&ws); err != nil {
		return Warehouse{}, nil
	}

	result, found := Warehouse{}, false
	for _, w := range ws {
		if w.WarehouseCode == code {
			result, found = w, true
			break
		}
	}

	if !found {
		return Warehouse{}, errors.New("can't find element with this warehouse_code")
	}

	return result, nil
}

func (r *repository) UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error) {
	var ws []Warehouse
	if err := r.file.Read(&ws); err != nil {
		return Warehouse{}, nil
	}

	result, updated := Warehouse{}, false
	for i, w := range ws {
		if w.Id == id {
			ws[i], updated = Warehouse{
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

	if !updated {
		return Warehouse{}, errors.New("can't find element with this id")
	}

	if err := r.file.Write(ws); err != nil {
		return Warehouse{}, err
	}

	return result, nil
}

func (r *repository) DeleteById(id int) error {
	var ws []Warehouse
	if err := r.file.Read(&ws); err != nil {
		return nil
	}

	deleted := false
	for i, w := range ws {
		if w.Id == id {
			newWs := []Warehouse{}
			newWs = append(newWs, ws[:i]...)
			newWs = append(newWs, ws[i+1:]...)
			ws = newWs
			deleted = true
			break
		}
	}

	if !deleted {
		return errors.New("can't find element with this id")
	}

	if err := r.file.Write(ws); err != nil {
		return err
	}

	return nil
}
