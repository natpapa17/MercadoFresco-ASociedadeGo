package products

import (
	"fmt"
)

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Create(product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error)
	Update(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewProductService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()

	if err != nil {
		return []Product{}, err
	}

	return ps, nil
}

func (s *service) GetById(id int) (Product, error) {
	ps, err := s.repository.GetById(id)

	if err != nil {
		return Product{}, err
	}

	return ps, nil
}

func (s *service) Create(product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Product{}, err
	}

	lastID++

	ps, _ := s.repository.GetByCode(product_code)

	if ps.Product_Code != "" {
		return Product{}, fmt.Errorf("product code: %s is already in use", ps.Product_Code)
	}

	product, err := s.repository.Create(lastID, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id)

	if err != nil {
		return Product{}, err
	}

	return product, nil

}

func (s *service) Update(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (Product, error) {
	codeProductInUse, _ := s.repository.GetByCode(product_code)

	if codeProductInUse.Product_Code != "" && codeProductInUse.Id != id {
		return Product{}, fmt.Errorf("product code: %s is already in use", codeProductInUse.Product_Code)
	}

	product, err := s.repository.Update(id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id)

	if err != nil {
		return Product{}, err
	}

	return product, err
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return err
}
