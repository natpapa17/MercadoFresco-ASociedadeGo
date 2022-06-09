package products

import (
	"errors"
	"fmt"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	GetByCode(productCode string) (Product, error)
	Create(id int, productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate int, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error)
	LastID() (int, error)
	Update(id int, productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate int, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
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
	return Product{}, errors.New("no product was found")
}

func (r *repository) GetByCode(productCode string) (Product, error) {
	var ps []Product

	if err := r.db.Read(&ps); err != nil {
		return Product{}, nil
	}
	for _, p := range ps {
		if p.ProductCode == productCode {
			return p, errors.New("product code already in use")
		}

	}
	return Product{}, errors.New("element not found")
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func (r *repository) Create(id int, productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate int, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return Product{}, err
	}
	p := Product{id, productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r repository) Update(id int, productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate int, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return Product{}, err
	}
	p := Product{Id: id, ProductCode: productCode, Description: description, Width: width, Height: height, Length: length, NetWeight: netWeight, ExpirationRate: expirationRate, RecommendedFreezingTemperature: recommendedFreezingTemperature, FreezingRate: freezingRate, ProductTypeId: productTypeId, SellerId: sellerId}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	if err := r.db.Write(ps); err != nil {
		return Product{}, errors.New("can't write the product file")
	}
	return p, nil
}

func (r repository) Delete(id int) error {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return errors.New("can't read the product")
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
		return fmt.Errorf("produto %d nao encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)
	if err := r.db.Write(ps); err != nil {
		return errors.New("can't write the product file")
	}

	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
