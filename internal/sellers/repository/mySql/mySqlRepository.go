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
		rows.Scan(&s.Id, &s.Cid, &s.CompanyName, &s.Address, &s.Telephone)
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

	//if s.Id == 0 {
		//return domain.Seller{}, errors.New("can't find element with this code")}
	//}

	return s, nil
}

func (*mySqlRepository) LastID() (int, error) {
	panic("unimplemented")
}

func (*mySqlRepository) Store(id int, cid int, companyName string, address string, telephone string) (domain.Seller, error) {
	panic("unimplemented")
}


func (*mySqlRepository) Update(id int, cid int, companyName string, address string, telephone string) (domain.Seller, error) {
	panic("unimplemented")
}

func CreateMySQLRepository(db *sql.DB) repository.Repository {
	return &mySqlRepository{
		db: db,
	}
}
