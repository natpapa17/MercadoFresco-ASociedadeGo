package products

import (
	"fmt"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product

	if err := r.db.Read(&ps); err != nil {
		return []Product{}, nil

	}
	return ps, nil
}

func (r *repository) GetById(id int) (Product, error) {
	var ps []Product

	if err := r.db.Read(&ps); err != nil {
		return Product{}, nil
	}

	for _, p := range ps {
		if p.Id == id {
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("product %d was not found", id)
}

func (r *repository) GetByCode(product_code string) (Product, error) {
	var ps []Product

	if err := r.db.Read(&ps); err != nil {
		return Product{}, nil
	}

	for _, p := range ps {
		if p.Product_Code == product_code {
			return p, fmt.Errorf("product with code: %s is already in use", product_code)
		}

	}
	return Product{}, fmt.Errorf("product with code: %s was not found", product_code)
}

func (r *repository) lastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}

func (r *repository) Create(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error) {
	var ps []Product

	if err := r.db.Read(&ps); err != nil {
		return Product{}, err
	}

	p := Product{
		id,
		product_code,
		description,
		width,
		height,
		length,
		net_weight,
		expiration_rate,
		recommended_freezing_temperature,
		freezing_rate,
		product_type_id,
		seller_id,
	}

	ps = append(ps, p)

	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r repository) Update(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error) {
	var ps []Product

	if err := r.db.Read(&ps); err != nil {
		return Product{}, err
	}

	p := Product{
		Id:                               id,
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
	}

	updated := false

	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("product %d not found", id)
	}

	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r repository) Delete(id int) error {
	var ps []Product

	if err := r.db.Read(&ps); err != nil {
		return err
	}

	deleted := false
	var index int

	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("product %d not found", id)
	}

	ps = append(ps[:index], ps[index+1:]...)

	if err := r.db.Write(ps); err != nil {
		return err
	}
	return nil
}
