package records

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/products_rec"
)

type Service interface {
	GetRecordsPerProduct(product_id int) (ReportRecords, error)
	Create(last_update_date string, purchase_price int, sale_price int, product_id int) (Records, error)
}

type service struct {
	repository        Repository
	productRepository products_rec.ProductRepository
}

func NewRecordsService(r Repository, p products_rec.ProductRepository) Service {
	return &service{
		repository:        r,
		productRepository: p,
	}
}

func (s *service) GetRecordsPerProduct(product_id int) (ReportRecords, error) {
	product, err := s.productRepository.GetById(product_id)

	if err != nil {
		return ReportRecords{}, err
	}

	rs, err := s.repository.GetRecordsPerProduct(product_id)

	if err != nil {
		return ReportRecords{}, err
	}

	result := ReportRecords{
		Product_Id:    product_id,
		Description:   product.Description,
		Records_Count: rs,
	}

	return result, nil
}

func (s *service) Create(last_update_date string, purchase_price int, sale_price int, product_id int) (Records, error) {
	hasProductId, err := s.productRepository.GetById(product_id)

	if hasProductId.Id != 0 {
		return Records{}, errors.New("product already in use")
	}

	if err != nil {
		return Records{}, errors.New("element not found")
	}

	record, err := s.repository.Create(last_update_date, purchase_price, sale_price, product_id)

	if err != nil {
		return Records{}, err
	}

	return record, nil
}
