package adapters

import (
	"database/sql"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/record_domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/usecases"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewRecordMysqlRepository(db *sql.DB) usecases.RecordsRepository {
	return &mysqlRepository{
		db: db,
	}
}

func (r *mysqlRepository) GetRecordsPerProduct(product_id int) (int, error) {
	const query = `SELECT COUNT(*) FROM product_record WHERE product_id=?`

	quantity := 0
	err := r.db.QueryRow(query, product_id).Scan(&quantity)

	if err != nil {
		return 0, err
	}

	return quantity, nil
}

func (r *mysqlRepository) Create(last_update_date string, purchase_price int, sale_price int, product_id int) (record_domain.Records, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return record_domain.Records{}, err
	}

	const query = `INSERT INTO product_record (last_update_date, purchase_price, sale_price, product_id) values (?, ?, ?, ?)`

	res, err := tx.Exec(query, last_update_date, purchase_price, sale_price, product_id)

	if err != nil {
		_ = tx.Rollback()
		return record_domain.Records{}, err
	}

	if err = tx.Commit(); err != nil {
		return record_domain.Records{}, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return record_domain.Records{}, err
	}

	return record_domain.Records{
		Id:               int(id),
		Last_Update_Date: last_update_date,
		Purchase_Price:   purchase_price,
		Sale_Price:       sale_price,
		Product_Id:       product_id,
	}, nil
}
