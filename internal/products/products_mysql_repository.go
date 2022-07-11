package products

import (
	"database/sql"
	"errors"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) Repository {
	return &mysqlRepository{
		db: db,
	}
}

func (r *mysqlRepository) GetAll() ([]Product, error) {
	const query = `SELECT * FROM product`

	rows, err := r.db.Query(query)

	if err != nil {
		return []Product{}, err
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		p := Product{}

		rows.Scan(&p.Id, &p.Product_Code, &p.Description, &p.Width, &p.Height, &p.Length, &p.Net_Weight, &p.Expiration_Rate, &p.Recommended_Freezing_Temperature, &p.Freezing_Rate, &p.Product_Type_Id, &p.Seller_Id)

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *mysqlRepository) GetById(id int) (Product, error) {
	const query = `SELECT * FROM product WHERE id=?`

	rows, err := r.db.Query(query)

	if err != nil {
		return Product{}, err
	}

	defer rows.Close()

	product := Product{}

	if err = rows.Err(); err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *mysqlRepository) GetByCode(product_code string) (Product, error) {
	const query = `SELECT * FROM product WHERE product_code=?`

	row := r.db.QueryRow(query, product_code)

	p := Product{}
	row.Scan(&p.Id, &p.Product_Code, &p.Description, &p.Width, &p.Height, &p.Length, &p.Net_Weight, &p.Expiration_Rate, &p.Recommended_Freezing_Temperature, &p.Freezing_Rate, &p.Product_Type_Id, &p.Seller_Id)

	if err := row.Err(); err != nil {
		return Product{}, err
	}

	if p.Id == 0 {
		return Product{}, errors.New("can't find element with this code")
	}

	return p, nil
}

func (r *mysqlRepository) Create(product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return Product{}, err
	}

	const query = `INSERT INTO product (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := tx.Exec(query, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id)

	if err != nil {
		_ = tx.Rollback()
		return Product{}, err
	}

	if err = tx.Commit(); err != nil {
		return Product{}, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return Product{}, err
	}

	return Product{
		Id:                               int(id),
		Product_Code:                     product_code,
		Description:                      description,
		Width:                            width,
		Height:                           height,
		Length:                           length,
		Net_Weight:                       net_weight,
		Expiration_Rate:                  expiration_rate,
		Recommended_Freezing_Temperature: recommended_freezing_temperature,
		Freezing_Rate:                    freezing_rate,
		Product_Type_Id:                  product_type_id,
		Seller_Id:                        seller_id,
	}, nil
}

func (r *mysqlRepository) Update(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error) {
	const query = `UPDATE products SET product_code=?, description=?, width=?, height=?, length, net_weight=?, expiration_rate=?, recommended_freezing_temperature=?, freezing_rate=?, product_type_id=?, seller_id=? WHERE id=?`

	res, err := r.db.Exec(query, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_code, seller_id)

	if err != nil {
		return Product{}, err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return Product{}, err
	}

	if rows == 0 {
		if w, _ := r.GetById(id); w.Id == 0 {
			return Product{}, errors.New("can't find element with this id")
		}
	}

	return Product{
		Id:                               int(id),
		Product_Code:                     product_code,
		Description:                      description,
		Width:                            width,
		Height:                           height,
		Length:                           length,
		Net_Weight:                       net_weight,
		Expiration_Rate:                  expiration_rate,
		Recommended_Freezing_Temperature: recommended_freezing_temperature,
		Freezing_Rate:                    freezing_rate,
		Product_Type_Id:                  product_type_id,
		Seller_Id:                        seller_id,
	}, nil
}

func (r *mysqlRepository) Delete(id int) error {
	const query = `DELETE FROM product WHERE id=?`

	res, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {

		return errors.New("can't find element with this id")
	}

	return nil
}
