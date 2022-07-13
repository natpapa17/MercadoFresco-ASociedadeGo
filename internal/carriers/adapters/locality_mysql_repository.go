package adapters

import (
	"database/sql"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases"
)

type localityMySQLRepositoryAdapter struct {
	db *sql.DB
}

func CreateLocalityMySQLRepository(db *sql.DB) usecases.LocalityRepository {
	return &localityMySQLRepositoryAdapter{
		db: db,
	}
}

func (r *localityMySQLRepositoryAdapter) GetById(id int) (domain.Locality, error) {
	const query = `SELECT id, name, province_id FROM locality WHERE id=?`

	locality := domain.Locality{}

	err := r.db.QueryRow(query, id).Scan(&locality.Id, &locality.Name, &locality.ProvinceId)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Locality{}, usecases.ErrNoElementFound
	}

	if err != nil {
		return domain.Locality{}, err
	}

	return locality, nil
}

func (r *localityMySQLRepositoryAdapter) GetAll() (domain.Localities, error) {
	const query = `SELECT id, name, province_id FROM locality`

	rows, err := r.db.Query(query)

	if err != nil {
		return domain.Localities{}, err
	}

	defer rows.Close()

	localities := domain.Localities{}

	for rows.Next() {
		locality := domain.Locality{}
		rows.Scan(&locality.Id, &locality.Name, &locality.ProvinceId)
		localities = append(localities, locality)
	}

	if err = rows.Err(); err != nil {
		return domain.Localities{}, err
	}

	return localities, nil
}
