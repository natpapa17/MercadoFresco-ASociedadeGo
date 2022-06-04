package warehouses

import "errors"

type Service interface {
	Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error)
	GetAll() ([]Warehouse, error)
	GetById(id int) (Warehouse, error)
	UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error)
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

func (s *service) Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error) {
	_, err := s.repository.GetByWarehouseCode(warehouseCode)

	if err == nil {
		return Warehouse{}, &BusinessRuleError{errors.New("this warehouse_code is already in use")}
	}

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

func (s *service) UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (Warehouse, error) {
	isWareHouseCodeInUse, err := s.repository.GetByWarehouseCode(warehouseCode)

	if err == nil {
		if isWareHouseCodeInUse.Id != id {
			return Warehouse{}, &BusinessRuleError{errors.New("this warehouse_code is already in use")}
		}
	}

	w, err := s.repository.UpdateById(id, warehouseCode, address, telephone, minimumCapacity, minimumTemperature)

	if err != nil {
		return Warehouse{}, err
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
