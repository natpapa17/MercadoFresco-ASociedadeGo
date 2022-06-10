package products

import (
	"fmt"
)

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Create(productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate int, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error)
	Update(id int, productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate int, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		fmt.Println(ps)
		return nil, err
	}
	fmt.Println(ps)
	return ps, nil
}

func (s service) GetById(id int) (Product, error) {
	ps, err := s.repository.GetById(id)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (s service) Create(productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate int, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Product{}, err
	}

	lastID++

	ps, err := s.repository.GetByCode(productCode)

	if ps.ProductCode != "" {
		return Product{}, err
	}

	product, err := s.repository.Create(lastID, productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)

	if err != nil {
		return Product{}, err
	}

	return product, nil

}

func (s service) Update(id int, productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate int, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (Product, error) {
	codeProductInUse, err := s.repository.GetByCode(productCode)

	if codeProductInUse.ProductCode != "" && codeProductInUse.ProductCode != productCode {
		return Product{}, err
	}

	ps, err := s.repository.GetByCode(productCode)

	if ps.Id != id {
		return Product{}, err
	}

	product, err := s.repository.Update(id, productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)

	if err != nil {
		return Product{}, err
	}
	return product, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return err
}

func NewProductService(r Repository) Service {

	return service{
		repository: r,
	}
}
