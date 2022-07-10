package products_rec

type Service interface {
	GetById(id int) (Product, error)
}

type service struct {
	repository Repository
}

func NewProductService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetById(id int) (Product, error) {
	ps, err := s.repository.GetById(id)

	if err != nil {
		return Product{}, err
	}

	return ps, nil
}
