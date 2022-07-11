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

type Repository interface{
	GetAll() ([]domain.Seller, error)
	GetById(id int) (domain.Seller, error)
	Store(id int, cid int, companyName string, address string , telephone string , localityId int) (domain.Seller , error)
	LastID() (int, error)
	Update(id , cid int, companyName, address, telephone string, localityId int) (domain.Seller, error)
	Delete(id int) error
}

func (r *mySqlRepository) Delete(id int) error {
	const query = `DELETE FROM seller WHERE id=?`

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
	const query = `SELECT * FROM seller`

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
	stmt, err := r.db.Prepare("SELECT * FROM seller WHERE id = ?")
	if err != nil {
		return domain.Seller{}, err
	}
	defer stmt.Close()

	seller := domain.Seller{}

	err = stmt.QueryRow(id).Scan(
		&seller.Id,
		&seller.Cid,
		&seller.CompanyName,
		&seller.Address,
		&seller.Telephone,
		&seller.LocalityId,
	)
	if err != nil {
		return seller, fmt.Errorf("Seller %d not found", id)
	}
	return seller, nil
}

func (*mySqlRepository) LastID() (int, error) {
	panic("unimplemented")
}

func (r *mySqlRepository) Store( id, cid int, companyName string, address string, telephone string, localityId int) (domain.Seller, error) {
	stmt, err := r.db.Prepare(`INSERT INTO seller
	(cid,
	company_name,
	address,
   	telephone,
	locality_id) 
   	VALUES(?,?,?,?,?)`)

	if err != nil {
		return domain.Seller{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(
		cid,
		companyName,
		address,
		telephone,
		localityId,
	)
	if err != nil {
		return domain.Seller{}, err
	}
	lastID, err := rows.LastInsertId()
	if err != nil {
		return domain.Seller{}, err
	}
	newSeller := domain.Seller{int(lastID), cid, companyName, address, telephone, localityId}
	return newSeller, nil
}


func (r *mySqlRepository) Update(id int, cid int, companyName string, address string, telephone string, locality_Id int) (domain.Seller, error) {
	
	updatedSeller := domain.Seller{id, cid, companyName, address, telephone, locality_Id}
	stmt, err := r.db.Prepare(`UPDATE seller SET 
	 	cid=?,
	  	company_name=?,
		address=?,
		telephone=?,
		locality_id=? WHERE id=?`)
	if err != nil {
		return domain.Seller{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Exec(
		cid,
		companyName,
		address,
		telephone,
		locality_Id,
		id)
	if err != nil {
		
		return updatedSeller, err
	}

	_, err = rows.RowsAffected()
	if err != nil {
		return domain.Seller{}, err
	}
	return updatedSeller, nil
}


func CreateMySQLRepository(db *sql.DB) repository.Repository {
	return &mySqlRepository{
		db: db,
	}
}
