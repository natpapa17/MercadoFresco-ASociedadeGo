package employee

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

type WareHouseRepository interface {
	GetById(id int) (Warehouse, error)
}

type warehouseRepository struct {
	file store.Store
}

func CreateWarehouseRepository(file store.Store) WareHouseRepository {
	return &warehouseRepository{
		file: file,
	}
}

func (r *warehouseRepository) GetById(id int) (Warehouse, error) {
	var ws []Warehouse

	if err := r.file.Read(&ws); err != nil {
		return Warehouse{}, err
	}

	for _, w := range ws {
		if w.Id == id {
			return w, nil
		}
	}

	return Warehouse{}, &NoElementInFileError{errors.New("can't find element with this id")}
}
