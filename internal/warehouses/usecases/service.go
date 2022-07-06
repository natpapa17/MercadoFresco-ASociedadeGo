package usecases

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/domain"
)

type Service interface {
	Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error)
	GetAll() (domain.Warehouses, error)
	GetById(id int) (domain.Warehouse, error)
	UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error)
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

func (s *service) Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error) {
	w, err := s.repository.GetByWarehouseCode(warehouseCode)

	if w.Id != 0 {
		return domain.Warehouse{}, ErrWarehouseCodeInUse
	}

	if !errors.Is(err, ErrNoElementFound) {
		return domain.Warehouse{}, err
	}

	warehouse, err := s.repository.Create(warehouseCode, address, telephone, minimumCapacity, minimumTemperature)

	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, nil
}

func (s *service) GetAll() (domain.Warehouses, error) {
	ws, err := s.repository.GetAll()

	if err != nil {
		return domain.Warehouses{}, err
	}

	return ws, nil
}

func (s *service) GetById(id int) (domain.Warehouse, error) {
	w, err := s.repository.GetById(id)

	if err != nil {
		return domain.Warehouse{}, err
	}

	return w, nil
}

func (s *service) UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error) {
	isWareHouseCodeInUse, err := s.repository.GetByWarehouseCode(warehouseCode)

	if isWareHouseCodeInUse.Id != 0 && isWareHouseCodeInUse.Id != id {
		return domain.Warehouse{}, ErrWarehouseCodeInUse
	}

	if !errors.Is(err, ErrNoElementFound) {
		return domain.Warehouse{}, err
	}

	w, err := s.repository.UpdateById(id, warehouseCode, address, telephone, minimumCapacity, minimumTemperature)

	if err != nil {
		return domain.Warehouse{}, err
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
