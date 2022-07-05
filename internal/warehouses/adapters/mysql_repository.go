package adapters

import (
	"database/sql"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
)

type mySQLRepositoryAdapter struct {
	db *sql.DB
}

func CreateMySQLRepository(db *sql.DB) usecases.Repository {
	return &mySQLRepositoryAdapter{
		db: db,
	}
}

func (r *mySQLRepositoryAdapter) Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return domain.Warehouse{}, err
	}

	const query = `INSERT INTO warehouse (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)`

	res, err := tx.Exec(query, warehouseCode, address, telephone, minimumCapacity, minimumTemperature)

	if err != nil {
		_ = tx.Rollback()
		return domain.Warehouse{}, err
	}

	if err = tx.Commit(); err != nil {
		return domain.Warehouse{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Warehouse{}, err
	}

	return domain.Warehouse{
		Id:                 int(id),
		WarehouseCode:      warehouseCode,
		Address:            address,
		Telephone:          telephone,
		MinimumCapacity:    minimumCapacity,
		MinimumTemperature: minimumTemperature,
	}, nil
}

func (r *mySQLRepositoryAdapter) GetAll() (domain.Warehouses, error) {
	const query = `SELECT * FROM warehouse`

	rows, err := r.db.Query(query)

	if err != nil {
		return domain.Warehouses{}, err
	}

	defer rows.Close()

	ws := domain.Warehouses{}

	for rows.Next() {
		w := domain.Warehouse{}
		rows.Scan(&w.Id, &w.WarehouseCode, &w.Address, &w.Telephone, &w.MinimumCapacity, &w.MinimumTemperature)
		ws = append(ws, w)
	}

	if err = rows.Err(); err != nil {
		return domain.Warehouses{}, err
	}

	return ws, nil
}

func (r *mySQLRepositoryAdapter) GetById(id int) (domain.Warehouse, error) {
	const query = `SELECT * FROM warehouse WHERE id=?`

	row := r.db.QueryRow(query, id)

	w := domain.Warehouse{}
	row.Scan(&w.Id, &w.WarehouseCode, &w.Address, &w.Telephone, &w.MinimumCapacity, &w.MinimumTemperature)

	if err := row.Err(); err != nil {
		return domain.Warehouse{}, err
	}

	if w.Id == 0 {
		return domain.Warehouse{}, &usecases.NoElementFoundError{Err: errors.New("can't find element with this code")}
	}

	return w, nil
}

func (r *mySQLRepositoryAdapter) GetByWarehouseCode(code string) (domain.Warehouse, error) {
	const query = `SELECT * FROM warehouse WHERE warehouse_code=?`

	row := r.db.QueryRow(query, code)

	w := domain.Warehouse{}
	row.Scan(&w.Id, &w.WarehouseCode, &w.Address, &w.Telephone, &w.MinimumCapacity, &w.MinimumTemperature)

	if err := row.Err(); err != nil {
		return domain.Warehouse{}, err
	}

	if w.Id == 0 {
		return domain.Warehouse{}, &usecases.NoElementFoundError{Err: errors.New("can't find element with this code")}
	}

	return w, nil
}

func (r *mySQLRepositoryAdapter) UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error) {
	const query = `UPDATE warehouses SET warehouse_code=?, address=?, telephone=?, minimum_capacity=?, minimum_temperature=? WHERE id=?`

	res, err := r.db.Exec(query, warehouseCode, address, telephone, minimumCapacity, minimumTemperature, id)

	if err != nil {
		return domain.Warehouse{}, err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return domain.Warehouse{}, err
	}

	if rows == 0 {
		if w, _ := r.GetById(id); w.Id == 0 {
			return domain.Warehouse{}, &usecases.NoElementFoundError{Err: errors.New("can't find element with this id")}
		}
	}

	return domain.Warehouse{
		Id:                 id,
		WarehouseCode:      warehouseCode,
		Address:            address,
		Telephone:          telephone,
		MinimumCapacity:    minimumCapacity,
		MinimumTemperature: minimumTemperature,
	}, nil
}

func (r *mySQLRepositoryAdapter) DeleteById(id int) error {
	const query = `DELETE FROM warehouse WHERE id=?`

	res, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return &usecases.NoElementFoundError{Err: errors.New("can't find element with this id")}
	}

	return nil
}
