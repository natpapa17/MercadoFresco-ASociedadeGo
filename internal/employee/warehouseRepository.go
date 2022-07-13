package employee

import (
	"database/sql"
	"errors"
)

type WareHouseRepository interface {
	GetById(id int) (Warehouse, error)
}

type warehouseRepository struct {
	db *sql.DB
}

func CreateWarehouseMySQLRepository(db *sql.DB) WareHouseRepository {
	return &warehouseRepository{
		db: db,
	}
}

func (r *warehouseRepository) GetById(id int) (Warehouse, error) {
	const query = `SELECT id, warehouse_code, address, telephone, minimum_capacity,	minimum_temperature FROM warehouse WHERE id=?`

	w := Warehouse{}
	err := r.db.QueryRow(query, id).Scan(&w.Id, &w.WarehouseCode, &w.Address, &w.Telephone, &w.MinimumCapacity, &w.MinimumTemperature)

	if errors.Is(err, sql.ErrNoRows) {
		return Warehouse{}, &NoElementInFileError{errors.New("can't find element with this id")}
	}

	if err != nil {
		return Warehouse{}, err
	}

	return w, nil
}
