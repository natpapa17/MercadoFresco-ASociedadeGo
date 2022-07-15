package inbound_orders

import (
	"database/sql"
)

type Inbound_Orders_RepositoryInterface interface {
	GetNumberOfOdersByEmployeeId(id int) (int, error)
}

type Inbound_Orders_RepositoryStruct struct {
	db *sql.DB
}

func Create_Inbound_Orders_MySQLRepository(db *sql.DB) Inbound_Orders_RepositoryInterface {
	return &Inbound_Orders_RepositoryStruct{
		db: db,
	}
}

func (r *Inbound_Orders_RepositoryStruct) GetNumberOfOdersByEmployeeId(id int) (int, error) {

	const query = `SELECT COUNT(*) FROM inbound_order WHERE employee_id=?`
	quantity := 0
	err := r.db.QueryRow(query, id).Scan(&quantity)
	if err != nil {
		return 0, err
	}

	return quantity, nil
}
