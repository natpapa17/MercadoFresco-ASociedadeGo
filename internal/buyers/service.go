package buyers

type Service interface {
	Create(firstName string, lastName string, address string, document string) (Buyer, error)
	GetAll() ([]Buyer, error)
	GetById(id int) (Buyer, error)
	UpdateById(id int, firstName string, lastName string, address string, document string) (Buyer, error)
	DeleteById(id int) error
}

type service struct {
	repository Repository
}

func CreateService(r Repository) Service {
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

func (s *service) GetById(id int) (Buyer, error) {
	w, err := s.repository.GetById(id)

	if err != nil {
		return Buyer{}, err
	}

	return w, nil
}

func (s *service) UpdateById(id int, firstName string, lastName string, address string, document string) (Buyer, error) {
	w, err := s.repository.UpdateById(id, firstName, lastName, address, document)

	if err != nil {
		return Buyer{}, err
	}

	return w, nil
}

func (s *service) DeleteById(id int) error {
	err := s.repository.DeleteById(id)

	if err != nil {
		return err
	}

	return nil
}