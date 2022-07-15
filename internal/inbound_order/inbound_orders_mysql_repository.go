package inbound_order

import (
	"database/sql"
)

type Inbound_Orders_RepositoryInterface interface {
	GetNumberOfOdersByEmployeeId(id int) (int, error)
	Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (Inbound_orders, error)
}

type Inbound_Orders_RepositoryStruct struct {
	db *sql.DB
}

func Create_Inbound_Orders_MySQLRepository(db *sql.DB) Inbound_Orders_RepositoryInterface {
	return &Inbound_Orders_RepositoryStruct{
		db: db,
	}
}

func (r *Inbound_Orders_RepositoryStruct) Create(orderDate string, orderNumber string, employeeId int, productBatchId int, warehouseId int) (Inbound_orders, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return Inbound_orders{}, err
	}

	const query = `INSERT INTO inbound_order (order_date, order_number, product_batch_id, warehouse_id, employee_id) VALUES (?, ?, ?, ?, ?)`

	res, err := tx.Exec(query, orderDate, orderNumber, productBatchId, warehouseId, employeeId)

	if err != nil {
		_ = tx.Rollback()
		return Inbound_orders{}, err
	}

	if err = tx.Commit(); err != nil {
		return Inbound_orders{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Inbound_orders{}, err
	}
	return Inbound_orders{
		Id:               int(id),
		Order_date:       orderDate,
		Order_number:     orderNumber,
		Employee_id:      employeeId,
		Product_batch_id: productBatchId,
		Warehouse_id:     warehouseId,
	}, nil
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
