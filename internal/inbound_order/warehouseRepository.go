package inbound_order

import (
	"database/sql"
	"errors"
)

type WareHouseRepositoryInterface interface {
	CheckIfWareHouseExistById(id int) (bool, error)
}

type WarehouseRepositoryStruct struct {
	db *sql.DB
}

func CreateWarehouseMySQLRepository(db *sql.DB) WareHouseRepositoryInterface {
	return &WarehouseRepositoryStruct{
		db: db,
	}
}

func (r *WarehouseRepositoryStruct) CheckIfWareHouseExistById(id int) (bool, error) {
	const query = `SELECT id FROM warehouse WHERE id=?`

	w := Warehouse{}
	err := r.db.QueryRow(query, id).Scan(&w.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
