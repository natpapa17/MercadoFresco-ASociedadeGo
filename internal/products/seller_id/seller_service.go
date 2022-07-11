package seller_id

type Service interface {
	GetById(id int) (Seller, error)
}

type service struct {
	repository Repository
}

func NewSellerService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetById(id int) (Seller, error) {
	seller, err := s.repository.GetById(id)

	if err != nil {
		return Seller{}, err
	}

	return seller, nil
}
