package repository

import (
	"database/sql"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/domain"
)

type provincyMysqlRepository struct {
	db *sql.DB
}

type ProvincyRepository interface {
	GetById(id int) (domain.Provincy, error)
}


func CreateProvincyRepository(db *sql.DB) ProvincyRepository {
	return &provincyMysqlRepository{
		db: db,
	}
}

func (r *provincyMysqlRepository) GetById(id int) (domain.Provincy, error) {
	const query = `SELECT id, name, country_id FROM province WHERE id=?`

	provincy:= domain.Provincy{}

	err := r.db.QueryRow(query, id).Scan(&provincy.Id, &provincy.Name, &provincy.Country_id)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Provincy{}, errors.New("not found")
	}

	if err != nil {
		return domain.Provincy{}, err
	}

	return provincy, nil
}