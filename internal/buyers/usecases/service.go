package usecases

import "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/domain"

type Service interface {
	Create(firstName string, lastName string, address string, document string) (domain.Buyer, error)
	GetAll() (domain.Buyers, error)
	GetBuyerById(id int) (domain.Buyer, error)
	UpdateBuyerById(id int, firstName string, lastName string, address string, document string) (domain.Buyer, error)
	DeleteBuyerById(id int) error
}

type service struct {
	repository BuyerRepository
}

func CreateBuyerService(r BuyerRepository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(firstName string, lastName string, address string, document string) (domain.Buyer, error) {
	buyer, err := s.repository.Create(firstName, lastName, address, document)

	if err != nil {
		return domain.Buyer{}, err
	}

	return buyer, nil
}

func (s *service) GetAll() (domain.Buyers, error) {
	ws, err := s.repository.GetAll()

	if err != nil {
		return domain.Buyers{}, err
	}

	return ws, nil
}

func (s *service) GetBuyerById(id int) (domain.Buyer, error) {
	w, err := s.repository.GetBuyerById(id)

	if err != nil {
		return domain.Buyer{}, err
	}

	return w, nil
}

func (s *service) UpdateBuyerById(id int, firstName string, lastName string, address string, document string) (domain.Buyer, error) {
	w, err := s.repository.UpdateBuyerById(id, firstName, lastName, address, document)

	if err != nil {
		return domain.Buyer{}, err
	}

	return w, nil
}

func (s *service) DeleteBuyerById(id int) error {
	err := s.repository.DeleteBuyerById(id)

	if err != nil {
		return err
	}

	return nil
}