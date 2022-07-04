package buyers

type Service interface {
	Create(firstName string, lastName string, address string, document string) (Buyer, error)
	GetAll() ([]Buyer, error)
	GetBuyerById(id int) (Buyer, error)
	UpdateBuyerById(id int, firstName string, lastName string, address string, document string) (Buyer, error)
	DeleteBuyerById(id int) error
}

type service struct {
	repository Repository
}

func CreateBuyerService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(firstName string, lastName string, address string, document string) (Buyer, error) {
	buyer, err := s.repository.Create(firstName, lastName, address, document)

	if err != nil {
		return Buyer{}, err
	}

	return buyer, nil
}

func (s *service) GetAll() ([]Buyer, error) {
	ws, err := s.repository.GetAll()

	if err != nil {
		return []Buyer{}, err
	}

	return ws, nil
}

func (s *service) GetBuyerById(id int) (Buyer, error) {
	w, err := s.repository.GetBuyerById(id)

	if err != nil {
		return Buyer{}, err
	}

	return w, nil
}

func (s *service) UpdateBuyerById(id int, firstName string, lastName string, address string, document string) (Buyer, error) {
	w, err := s.repository.UpdateBuyerById(id, firstName, lastName, address, document)

	if err != nil {
		return Buyer{}, err
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