package adapters

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

type warehouseFileRepositoryAdapter struct {
	file store.Store
}

func CreateWarehouseFileRepository(file store.Store) usecases.WarehouseRepository {
	return &warehouseFileRepositoryAdapter{
		file: file,
	}
}

func (r *warehouseFileRepositoryAdapter) Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error) {
	var ws domain.Warehouses
	if err := r.file.Read(&ws); err != nil {
		return domain.Warehouse{}, err
	}

	lastId, err := r.lastId()

	if err != nil {
		return domain.Warehouse{}, err
	}

	w := domain.Warehouse{
		Id:                 lastId + 1,
		WarehouseCode:      warehouseCode,
		Address:            address,
		Telephone:          telephone,
		MinimumCapacity:    minimumCapacity,
		MinimumTemperature: minimumTemperature,
	}
	ws = append(ws, w)

	if err := r.file.Write(ws); err != nil {
		return domain.Warehouse{}, err
	}
	return w, nil
}

func (r *warehouseFileRepositoryAdapter) GetAll() (domain.Warehouses, error) {
	var ws domain.Warehouses
	if err := r.file.Read(&ws); err != nil {
		return domain.Warehouses{}, err
	}
	return ws, nil
}

func (r *warehouseFileRepositoryAdapter) GetById(id int) (domain.Warehouse, error) {
	var ws domain.Warehouses

	if err := r.file.Read(&ws); err != nil {
		return domain.Warehouse{}, err
	}

	for _, w := range ws {
		if w.Id == id {
			return w, nil
		}
	}

	return domain.Warehouse{}, usecases.ErrNoElementFound
}

func (r *warehouseFileRepositoryAdapter) GetByWarehouseCode(code string) (domain.Warehouse, error) {
	var ws domain.Warehouses
	if err := r.file.Read(&ws); err != nil {
		return domain.Warehouse{}, err
	}

	for _, w := range ws {
		if w.WarehouseCode == code {
			return w, nil
		}
	}

	return domain.Warehouse{}, usecases.ErrNoElementFound
}

func (r *warehouseFileRepositoryAdapter) UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error) {
	var ws domain.Warehouses
	if err := r.file.Read(&ws); err != nil {
		return domain.Warehouse{}, err
	}

	result, updated := domain.Warehouse{}, false
	for i, w := range ws {
		if w.Id == id {
			ws[i], updated = domain.Warehouse{
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
		return domain.Warehouse{}, usecases.ErrNoElementFound
	}

	if err := r.file.Write(ws); err != nil {
		return domain.Warehouse{}, err
	}

	return result, nil
}

func (r *warehouseFileRepositoryAdapter) DeleteById(id int) error {
	var ws domain.Warehouses
	if err := r.file.Read(&ws); err != nil {
		return err
	}

	deleted := false
	for i, w := range ws {
		if w.Id == id {
			newWs := domain.Warehouses{}
			newWs = append(newWs, ws[:i]...)
			newWs = append(newWs, ws[i+1:]...)
			ws = newWs
			deleted = true
			break
		}
	}

	if !deleted {
		return usecases.ErrNoElementFound
	}

	if err := r.file.Write(ws); err != nil {
		return err
	}

	return nil
}

func (r *warehouseFileRepositoryAdapter) lastId() (int, error) {
	var ws domain.Warehouses
	if err := r.file.Read(&ws); err != nil {
		return 0, err
	}

	if len(ws) == 0 {
		return 0, nil
	}

	return ws[len(ws)-1].Id, nil
}
