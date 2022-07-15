package usecases

import "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/products/domain"

type Service interface {
	GetById(id int) (domain.Seller, error)
}

type service struct {
	repository Repository
}

func NewSellerService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetById(id int) (domain.Seller, error) {
	seller, err := s.repository.GetById(id)

	if err != nil {
		return domain.Seller{}, err
	}

	return seller, nil
}
