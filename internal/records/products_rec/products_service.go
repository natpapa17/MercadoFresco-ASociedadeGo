package products_rec

type Service interface {
	GetById(id int) (Product, error)
}

type service struct {
	repository ProductRepository
}

func NewProductService(p ProductRepository) Service {
	return &service{
		repository: p,
	}
}

func (s *service) GetById(id int) (Product, error) {
	ps, err := s.repository.GetById(id)

	if err != nil {
		return Product{}, err
	}

	return ps, nil
}
