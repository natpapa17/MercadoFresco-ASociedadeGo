package repository

import (
	"database/sql"
	"fmt"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/domain"
)

type LocalityRequestCreate struct {
	Name string `json:"name" binding:"required"`
	Province_id int `json:"province_id" binding:"required"`
	Country_id int  `json:"country_id" binding:"required"`
	
	
}

type LocalityReport struct {
	Locality_id  int    `json:"locality_id"`
	Name string `json:"name"`
	SellersCount int    `json:"sellers_count"`
}


type LocalityRepository interface {
	Create(name string, province_id int) (domain.Locality, error)
	ReportAll() ([]LocalityReport, error)
	ReportById(id int) (LocalityReport, error)
	GetAll() ([]domain.Locality, error)
	
}

type mySqlRepository struct {
	db *sql.DB
}


func (r *mySqlRepository) Create(name string, province_id int) (domain.Locality , error) {
	stmt, err := r.db.Prepare(`INSERT INTO locality (name,province_id) VALUES(?,?)`)

	if err != nil {
		return domain.Locality{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(
		name,
		province_id,
		
	)
	if err != nil {
		return domain.Locality{}, fmt.Errorf("essa localidade já existe")
	}
	lastId, err := rows.LastInsertId()
	if err != nil {
		return domain.Locality{}, err
	}
	
	
	return domain.Locality{
		Id:           int(lastId),
		Name:         name,
		Province_id:  province_id,
		
	}, nil
	
}
//getby id country e provincy no repository


func (r *mySqlRepository) GetAll() ([]domain.Locality, error) {
	const query = `SELECT * FROM locality`

	rows, err := r.db.Query(query)

	if err != nil {
		return domain.Localities{}, err
	}

	defer rows.Close()

	ll := domain.Localities{}

	for rows.Next() {
		l := domain.Locality{}
		rows.Scan(&l.Id, &l.Name, &l.Province_id)
		ll = append(ll, l)
	}

	if err = rows.Err(); err != nil {
		return domain.Localities{}, err
	}

	return ll, nil

}


func (r *mySqlRepository) ReportAll() ([]LocalityReport, error) {
	var localityList []LocalityReport
	rows, err := r.db.Query(`SELECT locality.id, locality.name, COUNT(seller.locality_id) 
	FROM fresh_market.locality
	INNER JOIN seller ON locality.Id = seller.locality_id
	GROUP BY locality_Id;`)
	if err != nil {
		return localityList, err
	}
	defer rows.Close()

	for rows.Next() {
		locality := LocalityReport{}

		err := rows.Scan(
			&locality.Locality_id,
			&locality.Name,
			&locality.SellersCount,
		)
		if err != nil {
			return localityList, err
		}
		localityList = append(localityList, locality)
	}

	return localityList, nil
}


func (r *mySqlRepository) ReportById(id int) (LocalityReport, error) {
	stmt, err := r.db.Prepare(`SELECT locality.id, locality.name, COUNT(seller.locality_id) 
	FROM fresh_market.locality
	INNER JOIN seller ON locality.id = seller.locality_id
	WHERE locality.id = ?
	GROUP BY locality_id;`)
	if err != nil {
		return LocalityReport{}, err
	}

	defer stmt.Close()

	locality := LocalityReport{}

	err = stmt.QueryRow(id).Scan(
		&locality.Locality_id,
		&locality.Name,
		&locality.SellersCount,

		
	)
	if err != nil {
		return locality, fmt.Errorf("Locality %d not found", id)
	}

	return locality, nil
}

func CreateMySQLRepository(db *sql.DB) LocalityRepository {
	return &mySqlRepository{
		db: db,
	}
}
