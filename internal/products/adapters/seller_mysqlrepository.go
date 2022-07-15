package adapters

import (
	"database/sql"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/usecases"
)

type SellerMysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) usecases.Repository {
	return &SellerMysqlRepository{
		db: db,
	}
}

func (r *SellerMysqlRepository) GetById(id int) (domain.Seller, error) {
	const query = `SELECT * FROM sellers WHERE id=?`

	row := r.db.QueryRow(query, id)

	s := domain.Seller{}
	row.Scan(&s.Id)

	if err := row.Err(); err != nil {
		return domain.Seller{}, err
	}

	return s, nil
}
