package adapters

import (
	"database/sql"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases"
)

type carrierMySQLRepositoryAdapter struct {
	db *sql.DB
}

func CreateCarrierMySQLRepository(db *sql.DB) usecases.CarrierRepository {
	return &carrierMySQLRepositoryAdapter{
		db: db,
	}
}

func (r *carrierMySQLRepositoryAdapter) Create(cid string, companyName string, address string, telephone string, localityId int) (domain.Carrier, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return domain.Carrier{}, err
	}

	const query = `INSERT INTO carrier (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)`

	res, err := tx.Exec(query, cid, companyName, address, telephone, localityId)

	if err != nil {
		_ = tx.Rollback()
		return domain.Carrier{}, err
	}

	if err = tx.Commit(); err != nil {
		return domain.Carrier{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Carrier{}, err
	}

	return domain.Carrier{
		Id:          int(id),
		Cid:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
		LocalityId:  localityId,
	}, nil
}

func (r *carrierMySQLRepositoryAdapter) GetNumberOfCarriersPerLocality(localityId int) (int, error) {
	const query = `SELECT COUNT(*) FROM carrier WHERE locality_id=?;`

	quantity := 0
	err := r.db.QueryRow(query, localityId).Scan(&quantity)

	if err != nil {
		return 0, err
	}

	return quantity, nil
}

func (r *carrierMySQLRepositoryAdapter) GetByCid(cid string) (domain.Carrier, error) {
	const query = `SELECT id, cid, company_name, address, telephone, locality_id FROM carrier WHERE cid=?`

	c := domain.Carrier{}
	err := r.db.QueryRow(query, cid).Scan(&c.Id, &c.Cid, &c.CompanyName, &c.Address, &c.Telephone, &c.LocalityId)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Carrier{}, usecases.ErrNoElementFound
	}

	if err != nil {
		return domain.Carrier{}, err
	}

	return c, nil
}
