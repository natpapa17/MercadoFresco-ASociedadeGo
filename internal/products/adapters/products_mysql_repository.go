package adapters

import (
	"database/sql"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/usecases"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewProductMysqlRepository(db *sql.DB) usecases.RepositoryProduct {
	return &mysqlRepository{
		db: db,
	}
}

func (r *mysqlRepository) GetAll() (domain.Products, error) {
	const query = `SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product`

	rows, err := r.db.Query(query)

	if err != nil {
		return domain.Products{}, err
	}

	defer rows.Close()

	products := domain.Products{}

	for rows.Next() {
		p := domain.Product{}

		rows.Scan(&p.Id, &p.Product_Code, &p.Description, &p.Width, &p.Height, &p.Length, &p.Net_Weight, &p.Expiration_Rate, &p.Recommended_Freezing_Temperature, &p.Freezing_Rate, &p.Product_Type_Id, &p.Seller_Id)

		products = append(products, p)
	}

	return products, nil
}

func (r *mysqlRepository) GetById(id int) (domain.Product, error) {
	const query = `SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product WHERE id=?`

	p := domain.Product{}

	err := r.db.QueryRow(query, id).Scan(&p.Id, &p.Product_Code, &p.Description, &p.Width, &p.Height, &p.Length, &p.Net_Weight, &p.Expiration_Rate, &p.Recommended_Freezing_Temperature, &p.Freezing_Rate, &p.Product_Type_Id, &p.Seller_Id)

	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *mysqlRepository) GetByCode(product_code string) (domain.Product, error) {
	const query = `SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM product WHERE product_code=?`

	row := r.db.QueryRow(query, product_code)

	p := domain.Product{}
	row.Scan(&p.Id, &p.Product_Code, &p.Description, &p.Width, &p.Height, &p.Length, &p.Net_Weight, &p.Expiration_Rate, &p.Recommended_Freezing_Temperature, &p.Freezing_Rate, &p.Product_Type_Id, &p.Seller_Id)

	if err := row.Err(); err != nil {
		return domain.Product{}, err
	}

	if p.Id == 0 {
		return domain.Product{}, errors.New("can't find element with this code")
	}

	return p, nil
}

func (r *mysqlRepository) Create(product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (domain.Product, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return domain.Product{}, err
	}

	const query = `INSERT INTO product (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := tx.Exec(query, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id)

	if err != nil {
		_ = tx.Rollback()
		return domain.Product{}, err
	}

	if err = tx.Commit(); err != nil {
		return domain.Product{}, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return domain.Product{}, err
	}

	return domain.Product{
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

func (r *mysqlRepository) Update(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (domain.Product, error) {
	const query = `UPDATE product SET product_code=?, description=?, width=?, height=?, length=?, net_weight=?, expiration_rate=?, recommended_freezing_temperature=?, freezing_rate=?, product_type_id=?, seller_id=? WHERE id=?`

	res, err := r.db.Exec(query, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id, id)

	if err != nil {
		return domain.Product{}, err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return domain.Product{}, err
	}

	if rows == 0 {
		if w, _ := r.GetById(id); w.Id == 0 {
			return domain.Product{}, errors.New("can't find element with this id")
		}
	}

	return domain.Product{
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
