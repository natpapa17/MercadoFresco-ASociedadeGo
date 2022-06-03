package warehouses

type Service interface {
	Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error)
	GetAll() ([]Warehouse, error)
	GetById(id int) (Warehouse, error)
}

type service struct {
	repository Repository
}

func CreateService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error) {
	warehouse, err := s.repository.Create(warehouseCode, address, telephone, minimumCapacity, minimumTemperature)

	if err != nil {
		return Warehouse{}, err
	}

	return warehouse, nil
}

func (s *service) GetAll() ([]Warehouse, error) {
	ws, err := s.repository.GetAll()

	if err != nil {
		return []Warehouse{}, err
	}

	return ws, nil
}

func (s *service) GetById(id int) (Warehouse, error) {
	w, err := s.repository.GetById(id)

	if err != nil {
		return Warehouse{}, err
	}

	return w, nil
}
