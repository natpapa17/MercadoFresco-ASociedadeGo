package buyers

import (
	"database/sql"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases"

)

type buyerRepository struct{
	db *sql.DB
}

func CreateBuyerRepository(db *sql.DB) Repository {
	return &buyerRepository{
		db: db,
	}
}

func (r *buyerRepository) Create(firstName string, lastName string, address string, document string) (domain.Buyer, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return domain.Buyer{}, err
	}

	const query = `INSERT INTO buyer (first_name, last_name, address, document_number) VALUES (?, ?, ?, ?, ?)`

	res, err := tx.Exec(query, firstName, lastName, address, document)


	if err != nil {
		_ = tx.Rollback()
		return domain.Buyer{}, err
	}

	if err = tx.Commit(); err != nil {
		return domain.Buyer{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Buyer{}, err
	}

	return domain.Buyer{
		ID: int(id),
		FirstName: firstName,
		LastName: lastName,
		Address: address,
		DocumentNumber: document,
	}, nil
}

func (r *buyerRepository) GetAll() (domain.Buyers, error) {
	const query = `SELECT first_name, last_name, address, document_number FROM buyer`

	rows, err := r.db.Query(query)
	if err != nil {
		return domain.Buyers{}, err
	}

	defer rows.Close()

	bs := domain.Buyers{}

	for rows.Next() {
		b := domain.Buyer{}
		rows.Scan(&b.ID, &b.FirstName, &b.LastName, &b.Address, &b.DocumentNumber)
		bs = append(bs, b)
	}

	if err = rows.Err(); err != nil {
		return domain.Buyers{}, err
	}

	return bs, nil
}

func (r *buyerRepository) GetBuyerById(id int) (domain.Buyer, error) {
	const query = `SELECT id, first_name, last_name, address, document_number FROM buyer WHERE id=?`

	b := domain.Buyer{}
	err := r.db.QueryRow(query, id).Scan(&b.ID, &b.FirstName, &b.LastName, &b.Address, &b.DocumentNumber)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Buyer{}, usecases.ErrNoElementFound
	}

	if err != nil {
		return domain.Buyer{}, err
	}

	return b, nil

}

func (r *buyerRepository) UpdateBuyerById(id int, firstName string, lastName string, address string, document string) (Buyer, error) {
	var bs []Buyer
	if err := r.file.Read(&bs); err != nil {
		return Buyer{}, nil
	}
	result, found := Buyer{}, false
	for i, b := range bs {
		if b.ID == id {
			bs[i], found = Buyer{
				ID:                 id,
				FirstName:      firstName,
				LastName:          lastName,
				Address:            address,
				DocumentNumber:    document,
			}, true
			result = bs[i]
			break
		}
	}

	if !found {
		return Buyer{}, &NoElementInFileError{errors.New("can't find element with this id")}
	}
	
	if err := r.file.Write(bs); err != nil {
		return Buyer{}, err
	}

	return result, nil
}

func (r *buyerRepository) DeleteBuyerById(id int) error {
	const query = `DELETE FROM buyer WHERE id=?`

	res, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return usecases.ErrNoElementFound
	}

	return nil

}

func (r *buyerRepository) lastId() (int, error) {
	var bs []Buyer
	if err := r.file.Read(&bs); err != nil {
		return 0, err
	}

	if len(bs) == 0 {
		return 0, nil
	}

	return bs[len(bs)-1].ID, nil
}