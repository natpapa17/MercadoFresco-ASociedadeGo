package seller_id

import (
	"database/sql"
)

type SellerMysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) Repository {
	return &SellerMysqlRepository{
		db: db,
	}
}

func (r *SellerMysqlRepository) GetById(id int) (Seller, error) {
	const query = `SELECT * FROM sellers WHERE id=?`

	row := r.db.QueryRow(query, id)

	s := Seller{}
	row.Scan(&s.Id)

	if err := row.Err(); err != nil {
		return Seller{}, err
	}

	return s, nil
}
