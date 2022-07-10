package records

import "errors"

type Service interface {
	GetRecordsPerProduct(product_id int) (int, error)
	GetByProductId(id int) (Records, error)
	Create(last_update_date string, purchase_price int, sale_price int, product_id int) (Records, error)
}

type service struct {
	repository Repository
}

func NewRecordsService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetRecordsPerProduct(product_id int) (int, error) {
	rs, err := s.repository.GetRecordsPerProduct(product_id)

	if err != nil {
		return 0, err
	}

	return rs, nil
}

func (s *service) GetByProductId(id int) (Records, error) {
	rs, err := s.repository.GetByProductId(id)

	if err != nil {
		return Records{}, err
	}

	return rs, nil
}

func (s *service) Create(last_update_date string, purchase_price int, sale_price int, product_id int) (Records, error) {
	hasProductId, err := s.repository.GetByProductId(product_id)

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
