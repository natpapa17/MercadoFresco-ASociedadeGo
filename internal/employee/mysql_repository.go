package employee

import (
	"database/sql"
	"errors"
)

type employeeInterface interface {
	Create(cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error)
	GetAll() ([]Employee, error)
	GetById(id int) (Employee, error)
	UpdateById(id int, cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error)
	DeleteById(id int) error
	GetByCardNumberId(cardNumberId int) (Employee, error)
}

type employeeMYSQLRepository struct {
	db *sql.DB
}

func CreateEmployeeMysqlRepository(db *sql.DB) employeeInterface {
	return &employeeMYSQLRepository{
		db: db,
	}
}

func (r *employeeMYSQLRepository) Create(cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return Employee{}, err
	}

	const query = `INSERT INTO employee (id_card_number, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)`
	res, err := tx.Exec(query, cardNumberId, firstName, lastName, wareHouseId)

	if err != nil {
		_ = tx.Rollback()
		return Employee{}, err
	}

	if err = tx.Commit(); err != nil {
		return Employee{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Employee{}, err
	}

	return Employee{
		Id:             int(id),
		Card_number_id: cardNumberId,
		First_name:     firstName,
		Last_name:      lastName,
		Warehouse_id:   wareHouseId,
	}, nil
}

func (r *employeeMYSQLRepository) GetAll() ([]Employee, error) {
	const query = `SELECT id, id_card_number, first_name, last_name, warehouse_id, FROM employee`

	rows, err := r.db.Query(query)

	if err != nil {
		return Employees{}, err
	}

	defer rows.Close()

	es := Employees{}

	for rows.Next() {
		e := Employee{}
		rows.Scan(&e.Id, &e.Card_number_id, &e.First_name, &e.Last_name, &e.Warehouse_id)
		es = append(es, e)
	}

	if err = rows.Err(); err != nil {
		return []Employee{}, err
	}

	return es, nil
}

func (r *employeeMYSQLRepository) GetById(id int) (Employee, error) {
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

func (r *employeeMYSQLRepository) GetByCardNumberId(cardNumberId int) (Employee, error) {
	const query = `SELECT id, id_card_number, first_name, last_name, warehouse_id FROM employee WHERE id_card_number=?`

	e := Employee{}
	err := r.db.QueryRow(query, cardNumberId).Scan(&e.Id, &e.Card_number_id, &e.First_name, &e.Last_name, &e.Warehouse_id)

	if errors.Is(err, sql.ErrNoRows) {
		return Employee{}, &NoElementInFileError{errors.New("can't find element with this card number id")}
	}

	if err != nil {
		return Employee{}, err
	}

	return e, nil
}

func (r *employeeMYSQLRepository) UpdateById(id int, cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error) {
	const query = `UPDATE employee SET id_card_number=?, first_name=?, last_name=?, warehouse_id=? WHERE id=?`

	res, err := r.db.Exec(query, cardNumberId, firstName, lastName, wareHouseId, id)

	if err != nil {
		return Employee{}, err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return Employee{}, err
	}

	if rows == 0 {
		if e, _ := r.GetById(id); e.Id == 0 {
			return Employee{}, &NoElementInFileError{errors.New("can't find element with this id")}
		}
	}

	return Employee{
		Id:             int(id),
		Card_number_id: cardNumberId,
		First_name:     firstName,
		Last_name:      lastName,
		Warehouse_id:   wareHouseId,
	}, nil
}

func (r *employeeMYSQLRepository) DeleteById(id int) error {
	const query = `DELETE FROM employee WHERE id=?`

	res, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return &NoElementInFileError{errors.New("can't find element with this id")}
	}

	return nil
}
