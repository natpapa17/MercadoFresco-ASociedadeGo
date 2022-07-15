package inbound_orders

import (
	"database/sql"
	"errors"
)

type ProductBatchRepositoryInterface interface {
	CheckProductBatchIfExistById(id int) (bool, error)
}

type ProductBatchmySQLRepositoryStruct struct {
	db *sql.DB
}

func CreateNewMySQLRepositoryBatchProduct(db *sql.DB) ProductBatchRepositoryInterface {
	return &ProductBatchmySQLRepositoryStruct{
		db: db,
	}
}

func (r *ProductBatchmySQLRepositoryStruct) CheckProductBatchIfExistById(id int) (bool, error) {
	const query = `SELECT id FROM product_batch WHERE id=?`

	p := ProductBatch{}

	err := r.db.QueryRow(query, id).Scan(&p.Id)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
