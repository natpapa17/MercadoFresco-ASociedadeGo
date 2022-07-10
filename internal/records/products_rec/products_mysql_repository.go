package products_rec

import (
	"database/sql"
)

type ProductMysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) Repository {
	return &ProductMysqlRepository{
		db: db,
	}
}

func (r *ProductMysqlRepository) GetById(id int) (Product, error) {
	const query = `SELECT * FROM product WHERE id=?`

	rows, err := r.db.Query(query)

	if err != nil {
		return Product{}, err
	}

	defer rows.Close()

	product := Product{}

	if err = rows.Err(); err != nil {
		return Product{}, err
	}

	return product, nil
}
