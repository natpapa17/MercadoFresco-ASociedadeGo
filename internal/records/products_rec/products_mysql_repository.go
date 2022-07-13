package products_rec

import (
	"database/sql"
)

type ProductMysqlRepository struct {
	db *sql.DB
}

func NewMysqlProductRepository(db *sql.DB) ProductRepository {
	return &ProductMysqlRepository{
		db: db,
	}
}

func (r *ProductMysqlRepository) GetById(id int) (Product, error) {
	const query = `SELECT id, description FROM product WHERE id=?`
	product := Product{}

	err := r.db.QueryRow(query, id).Scan(&product.Id, &product.Description)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}
