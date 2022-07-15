package inbound_orders

import (
	"database/sql"
	"errors"
)

type EmployeeInboundInterface interface {
	GetById(id int) (Employee, error)
}

type EmployeeMYSQLRepositoryInboundStruct struct {
	db *sql.DB
}

func CreateEmployeeMysqlRepositoryInbound(db *sql.DB) EmployeeInboundInterface {
	return &EmployeeMYSQLRepositoryInboundStruct{
		db: db,
	}
}

func (r *EmployeeMYSQLRepositoryInboundStruct) GetById(id int) (Employee, error) {
	const query = `SELECT id, id_card_number, first_name, last_name, warehouse_id FROM employee WHERE id=?`

	e := Employee{}
	err := r.db.QueryRow(query, id).Scan(&e.Id, &e.Card_number_id, &e.First_name, &e.Last_name, &e.Warehouse_id)

	if errors.Is(err, sql.ErrNoRows) {
		return Employee{}, &NoElementInFileError{errors.New("can't find element with this id")}
	}

	if err != nil {
		return Employee{}, err
	}

	return e, nil
}
