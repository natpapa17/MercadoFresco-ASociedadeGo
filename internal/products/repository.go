package products

import (
	"fmt"
)

var ps []Product = []Product{}

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (*Product, error)
	Create(Id int, ProductCode int, Description string, Width float64, Height float64, Length float64, NetWeight float64, ExpirationRate string, RecommendedFreezingTemperature float64, FreezingRate float64, ProductTypeId string, SellerId string) (Product, error)
	LastID() (int, error)
	Update(Id int, ProductCode int, Description string, Width float64, Height float64, Length float64, netWeight float64, ExpirationRate string, RecommendedFreezingTemperature float64, FreezingRate float64, ProductTypeId string, SellerId string) (Product, error)
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

func (r *repository) GetById() (*Product, error) {
	var ps *Product
	if err := r.db.Read(&ps); err != nil {
		return &Product{}, nil
	}
	return ps, nil
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

func (r *repository) Create(Id int, ProductCode int, Description string, Width float64, Height float64, Length float64, NetWeight float64, ExpirationRate string, RecommendedFreezingTemperature float64, FreezingRate float64, ProductTypeId string, SellerId string) (Product, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return Product{}, err
	}
	p := Product{Id, ProductCode, Description, Width, Height, Length, NetWeight, ExpirationRate, RecommendedFreezingTemperature, FreezingRate, ProductTypeId, SellerId}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (repository) Update(Id int, ProductCode int, Description string, Width float64, Height float64, Length float64, netWeight float64, ExpirationRate string, RecommendedFreezingTemperature float64, FreezingRate float64, ProductTypeId string, SellerId string) (Product, error) {
	p := Product{Id: Id, ProductCode: ProductCode, Description: Description, Width: Width, Height: Height, Length: Length, NetWeight: netWeight, ExpirationRate: ExpirationRate, RecommendedFreezingTemperature: RecommendedFreezingTemperature, FreezingRate: FreezingRate, ProductTypeId: ProductTypeId, SellerId: SellerId}
	updated := false
	for i := range ps {
		if ps[i].Id == Id {
			p.Id = Id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", Id)
	}
	return p, nil
}

func (repository) UpdateId(Id int) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].Id == Id {
			ps[i].Id = Id
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d no encontrado", id)
	}
	return p, nil
}

func (repository) Delete(id int) error {
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
	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
