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
