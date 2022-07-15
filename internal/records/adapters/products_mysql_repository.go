package adapters

import (
	"database/sql"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/record_domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/usecases"
)

type ProductMysqlRepository struct {
	db *sql.DB
}

func NewMysqlProductRepository(db *sql.DB) usecases.ProductRepository {
	return &ProductMysqlRepository{
		db: db,
	}
}

func (r *ProductMysqlRepository) GetById(id int) (record_domain.Product, error) {
	const query = `SELECT id, description FROM product WHERE id=?`
	product := record_domain.Product{}

	err := r.db.QueryRow(query, id).Scan(&product.Id, &product.Description)

	if err != nil {
		return record_domain.Product{}, err
	}

	return product, nil
}
