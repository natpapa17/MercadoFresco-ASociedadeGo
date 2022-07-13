package usecases

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/domain"
)

type WarehouseService interface {
	Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error)
	GetAll() (domain.Warehouses, error)
	GetById(id int) (domain.Warehouse, error)
	UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error)
	DeleteById(id int) error
}

type warehouseService struct {
	warehouseRepository WarehouseRepository
}

func CreateWarehouseService(r WarehouseRepository) WarehouseService {
	return &warehouseService{
		warehouseRepository: r,
	}
}

func (s *warehouseService) Create(warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error) {
	hasWarehouseWithThisCode, err := s.warehouseRepository.GetByWarehouseCode(warehouseCode)

	if hasWarehouseWithThisCode.Id != 0 {
		return domain.Warehouse{}, ErrWarehouseCodeInUse
	}

	if err != nil && !errors.Is(err, ErrNoElementFound) {
		return domain.Warehouse{}, err
	}

	warehouse, err := s.warehouseRepository.Create(warehouseCode, address, telephone, minimumCapacity, minimumTemperature)

	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, nil
}

func (s *warehouseService) GetAll() (domain.Warehouses, error) {
	warehouses, err := s.warehouseRepository.GetAll()

	if err != nil {
		return domain.Warehouses{}, err
	}

	return warehouses, nil
}

func (s *warehouseService) GetById(id int) (domain.Warehouse, error) {
	warehouse, err := s.warehouseRepository.GetById(id)

	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, nil
}

func (s *warehouseService) UpdateById(id int, warehouseCode string, address string, telephone string, minimumCapacity int, minimumTemperature float64) (domain.Warehouse, error) {
	hasWarehouseWithThisCode, err := s.warehouseRepository.GetByWarehouseCode(warehouseCode)

	if hasWarehouseWithThisCode.Id != 0 && hasWarehouseWithThisCode.Id != id {
		return domain.Warehouse{}, ErrWarehouseCodeInUse
	}

	if err != nil && !errors.Is(err, ErrNoElementFound) {
		return domain.Warehouse{}, err
	}

	warehouse, err := s.warehouseRepository.UpdateById(id, warehouseCode, address, telephone, minimumCapacity, minimumTemperature)

	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, nil
}

func (s *warehouseService) DeleteById(id int) error {
	err := s.warehouseRepository.DeleteById(id)

	if err != nil {
		return err
	}

	return nil
}
