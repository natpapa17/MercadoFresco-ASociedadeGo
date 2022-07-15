package usecases

import (
	"fmt"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/domain"
)

type ServiceProduct interface {
	GetAll() (domain.Products, error)
	GetById(id int) (domain.Product, error)
	Create(product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (domain.Product, error)
	Update(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (domain.Product, error)
	Delete(id int) error
}

type serviceProduct struct {
	repositoryProduct RepositoryProduct
}

func NewProductService(r RepositoryProduct) ServiceProduct {
	return &serviceProduct{
		repositoryProduct: r,
	}
}

func (s *serviceProduct) GetAll() (domain.Products, error) {
	ps, err := s.repositoryProduct.GetAll()

	if err != nil {
		return domain.Products{}, err
	}

	return ps, nil
}

func (s *serviceProduct) GetById(id int) (domain.Product, error) {
	ps, err := s.repositoryProduct.GetById(id)

	if err != nil {
		return domain.Product{}, err
	}

	return ps, nil
}

func (s *serviceProduct) Create(product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (domain.Product, error) {

	codeProductInUse, _ := s.repositoryProduct.GetByCode(product_code)

	if codeProductInUse.Product_Code != "" {
		return domain.Product{}, fmt.Errorf("product code: %s is already in use", codeProductInUse.Product_Code)
	}

	product, err := s.repositoryProduct.Create(product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil

}

func (s *serviceProduct) Update(id int, product_code string, description string, width float64, height float64, length float64, net_weight float64, expiration_rate int, recommended_freezing_temperature float64, freezing_rate int, product_type_id int, seller_id int) (domain.Product, error) {
	codeProductInUse, _ := s.repositoryProduct.GetByCode(product_code)

	if codeProductInUse.Product_Code != "" && codeProductInUse.Id != id {
		return domain.Product{}, fmt.Errorf("product code: %s is already in use", codeProductInUse.Product_Code)
	}

	product, err := s.repositoryProduct.Update(id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id)

	if err != nil {
		return domain.Product{}, err
	}

	return product, err
}

func (s *serviceProduct) Delete(id int) error {
	err := s.repositoryProduct.Delete(id)

	if err != nil {
		return err
	}

	return err
}
