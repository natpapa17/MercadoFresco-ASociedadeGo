package usecases

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/record_domain"
)

type RecordsService interface {
	GetRecordsPerProduct(product_id int) (record_domain.ReportRecord, error)
	Create(last_update_date string, purchase_price int, sale_price int, product_id int) (record_domain.Records, error)
}

type recordsService struct {
	recordsRepository RecordsRepository
	productRepository ProductRepository
}

func NewRecordsService(r RecordsRepository, p ProductRepository) RecordsService {
	return &recordsService{
		recordsRepository: r,
		productRepository: p,
	}
}

func (s *recordsService) GetRecordsPerProduct(product_id int) (record_domain.ReportRecord, error) {
	product, err := s.productRepository.GetById(product_id)

	if err != nil {
		return record_domain.ReportRecord{}, err
	}

	rp, err := s.recordsRepository.GetRecordsPerProduct(product_id)

	if err != nil {
		return record_domain.ReportRecord{}, err
	}

	result := record_domain.ReportRecord{
		Product_Id:    product_id,
		Description:   product.Description,
		Records_Count: rp,
	}

	return result, nil
}

func (s *recordsService) Create(last_update_date string, purchase_price int, sale_price int, product_id int) (record_domain.Records, error) {
	_, err := s.productRepository.GetById(product_id)

	if err != nil {
		return record_domain.Records{}, errors.New("element not found")
	}

	record, err := s.recordsRepository.Create(last_update_date, purchase_price, sale_price, product_id)

	if err != nil {
		return record_domain.Records{}, err
	}

	return record, nil
}
