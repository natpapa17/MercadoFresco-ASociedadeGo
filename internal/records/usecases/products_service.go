package usecases

import "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records/record_domain"

type Service interface {
	GetById(id int) (record_domain.Product, error)
}

type service struct {
	repository ProductRepository
}

func NewProductService(p ProductRepository) Service {
	return &service{
		repository: p,
	}
}

func (s *service) GetById(id int) (record_domain.Product, error) {
	ps, err := s.repository.GetById(id)

	if err != nil {
		return record_domain.Product{}, err
	}

	return ps, nil
}
