package repository

import (
	"database/sql"
	
	"fmt"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/repository"
)

type mySqlRepository struct {
	db *sql.DB
}

func (r *mySqlRepository) Delete(id int) error {
	const query = `DELETE FROM sellers WHERE id=?`

	res, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("seller with %d not found", id)
	}

	return nil
}

func (r *mySqlRepository) GetAll() ([]domain.Seller, error) {
	const query = `SELECT * FROM sellers`

	rows, err := r.db.Query(query)

	if err != nil {
		return domain.Sellers{}, err
	}

	defer rows.Close()

	sl := domain.Sellers{}

	for rows.Next() {
		s := domain.Seller{}
		rows.Scan(&s.Id, &s.Cid, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityId)
		sl = append(sl, s)
	}

	if err = rows.Err(); err != nil {
		return domain.Sellers{}, err
	}

	return sl, nil
}


func (r *mySqlRepository) GetById(id int) (domain.Seller, error) {
	const query = `SELECT * FROM sellers WHERE id=?`

	row := r.db.QueryRow(query, id)

	s := domain.Seller{}
	row.Scan(&s.Id, &s.Cid, &s.CompanyName, &s.Address, &s.Telephone)

	if err := row.Err(); err != nil {
		return domain.Seller{}, err
	}



	return s, nil
}

func (*mySqlRepository) LastID() (int, error) {
	panic("unimplemented")
}

func (r *mySqlRepository) Store( id, cid int, companyName string, address string, telephone string, localityId int) (domain.Seller, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return domain.Seller{}, err
	}

	const query = `INSERT INTO seller (cid, companyName, address, telephone, localityId) VALUES (?, ?, ?, ?, ?, ?)`

	res, err := tx.Exec(query, cid, companyName, address, telephone, localityId)

	if err != nil {
		_ = tx.Rollback()
		return domain.Seller{}, err
	}

	if err = tx.Commit(); err != nil {
		return domain.Seller{}, err
	}

	a, err := res.LastInsertId()
		if err != nil {
			return domain.Seller{}, err
		}

	return domain.Seller{
		Id:                 int(a),
		Cid:      cid,
		CompanyName:            companyName,
		Address:          address,
		Telephone:    telephone,
		LocalityId: localityId,
		
	}, nil
}


func (r *mySqlRepository) Update(id int, cid int, companyName string, address string, telephone string, localityId int) (domain.Seller, error) {
	
	panic("unimplemented")
}


func CreateMySQLRepository(db *sql.DB) repository.Repository {
	return &mySqlRepository{
		db: db,
	}
}
