package adapters

import (
	"database/sql"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases"

)

type buyerMySQLRepository struct{
	db *sql.DB
}

func CreateBuyerMySQLRepository(db *sql.DB) usecases.BuyerRepository {
	return &buyerMySQLRepository{
		db: db,
	}
}

func (r *buyerMySQLRepository) Create(firstName string, lastName string, address string, document string) (domain.Buyer, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return domain.Buyer{}, err
	}

	const query = `INSERT INTO buyer (first_name, last_name, address, document_number) VALUES (?, ?, ?, ?)`

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

func (r *buyerMySQLRepository) GetAll() (domain.Buyers, error) {
	const query = `SELECT id, first_name, last_name, address, document_number FROM buyer`

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

func (r *buyerMySQLRepository) GetBuyerById(id int) (domain.Buyer, error) {
	const query = `SELECT id, first_name, last_name, address, document_number FROM buyer WHERE id=?`

	b := domain.Buyer{}
	err := r.db.QueryRow(query, id).Scan(&b.ID, &b.FirstName, &b.LastName, &b.Address, &b.DocumentNumber)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Buyer{}, &usecases.ErrNoElementFound{Err: errors.New("Id não encontrado")}
	}

	if err != nil {
		return domain.Buyer{}, err
	}

	return b, nil

}

func (r *buyerMySQLRepository) UpdateBuyerById(id int, firstName string, lastName string, address string, document string) (domain.Buyer, error) {
	const query = `UPDATE buyer SET first_name=?, last_name=?, address=?, document_number=? WHERE id=?`

	res, err := r.db.Exec(query, firstName, lastName, address, document, id)

	if err != nil {
		return domain.Buyer{}, nil
	}
	
	rows, err := res.RowsAffected()

	if err != nil {
		return domain.Buyer{}, err
	}

	if rows == 0 {
		if b, _ := r.GetBuyerById(id); b.ID == 0 {
			return domain.Buyer{}, &usecases.ErrNoElementFound{Err: errors.New("elemento para update não encontrado")}
		}
	}

	return domain.Buyer{
		ID: int(id),
		FirstName: firstName,
		LastName: lastName,
		Address: address,
		DocumentNumber: document,
	}, nil
}

func (r *buyerMySQLRepository) DeleteBuyerById(id int) error {
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
		return &usecases.ErrNoElementFound{Err: errors.New("elemento para deletar não encontrado")}
	}

	return nil

}
