package repository

import (
	"database/sql"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/domain"
)

type countryMysqlRepository struct {
	db *sql.DB
}

type CountryRepository interface {
	GetById(id int) (domain.Country, error)
}


func CreateCountryRepository(db *sql.DB) CountryRepository {
	return &countryMysqlRepository{
		db: db,
	}
}

func (r *countryMysqlRepository) GetById(id int) (domain.Country, error) {
	const query = `SELECT id, name FROM country WHERE id=?`

	country:= domain.Country{}

	err := r.db.QueryRow(query, id).Scan(&country.Id, &country.Name)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Country{}, errors.New("not found")
	}

	if err != nil {
		return domain.Country{}, err
	}

	return country, nil
}