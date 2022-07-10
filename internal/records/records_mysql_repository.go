package records

import (
	"database/sql"
	"errors"
)

type Repository interface {
	GetRecordsPerProduct(product_id int) (int, error)
	GetByProductId(id int) (Records, error)
	Create(last_update_date string, purchase_price int, sale_price int, product_id int) (Records, error)
}

type mysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) Repository {
	return &mysqlRepository{
		db: db,
	}
}

func (r *mysqlRepository) GetRecordsPerProduct(product_id int) (int, error) {
	const query = `SELECT COUNT(*) FROM records WHERE product_id=?`

	quantity := 0
	err := r.db.QueryRow(query, product_id).Scan(&quantity)

	if err != nil {
		return 0, err
	}

	return quantity, nil
}

func (r *mysqlRepository) GetByProductId(product_id int) (Records, error) {
	const query = `SELECT id, last_update_date, purchase_price, sale_price, product_id FROM records WHERE id=?`

	rec := Records{}
	err := r.db.QueryRow(query, product_id).Scan(&rec.Id, &rec.Last_Update_Date, &rec.Purchase_Price, &rec.Sale_Price, &rec.Product_Id)

	if errors.Is(err, sql.ErrNoRows) {
		return Records{}, errors.New("product not found")
	}

	if err != nil {
		return Records{}, err
	}

	return rec, nil
}

func (r *mysqlRepository) Create(last_update_date string, purchase_price int, sale_price int, product_id int) (Records, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return Records{}, err
	}

	const query = `INSERT INTO records (last_update_date, purchase_price, sale_price, product_id) values (?, ?, ?, ?)`

	res, err := tx.Exec(query, last_update_date, purchase_price, sale_price, product_id)

	if err != nil {
		_ = tx.Rollback()
		return Records{}, err
	}

	if err = tx.Commit(); err != nil {
		return Records{}, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return Records{}, err
	}

	return Records{
		Id:               int(id),
		Last_Update_Date: last_update_date,
		Purchase_Price:   purchase_price,
		Sale_Price:       sale_price,
		Product_Id:       product_id,
	}, nil
}
