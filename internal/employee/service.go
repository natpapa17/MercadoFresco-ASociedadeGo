package employee

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
)

type EmployeeServiceInterface interface {
	Create(cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error)
	GetAll() ([]Employee, error)
	GetById(id int) (Employee, error)
	UpdateById(id int, cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error)
	DeleteById(id int) error
}

type service struct {
	repository    employeeInterface
	wareHouseRepo usecases.WarehouseRepository
}

func CreateService(r employeeInterface, w usecases.WarehouseRepository) EmployeeServiceInterface {
	return &service{
		repository:    r,
		wareHouseRepo: w,
	}
}

func (s *service) Create(cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error) {

	_, errCardNumberId := s.repository.GetByCardNumberId(cardNumberId)
	_, wareHouseErr := s.wareHouseRepo.GetById(wareHouseId)

	if errCardNumberId == nil {
		return Employee{}, errors.New("This Card Number id is already In Use!")
	}
	if wareHouseErr != nil {
		return Employee{}, errors.New("This WareHouse Id does not Exist, check the wareHuse list to get an id, or create a new WareHouse!")
	}

	employee, err := s.repository.Create(cardNumberId, firstName, lastName, wareHouseId)

	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (s *service) GetAll() ([]Employee, error) {
	es, err := s.repository.GetAll()

	if err != nil {
		return []Employee{}, err
	}

	return es, nil
}

func (s *service) GetById(id int) (Employee, error) {
	e, err := s.repository.GetById(id)

	if err != nil {
		return Employee{}, err
	}

	return e, nil
}

func (s *service) UpdateById(id int, cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error) {
	isCardNumberInUse, err := s.repository.GetByCardNumberId(cardNumberId)

	if err == nil {
		if isCardNumberInUse.Id != id {
			return Employee{}, &BusinessRuleError{errors.New("this CardNumberId is already in use")}
		}
	}

	e, err := s.repository.UpdateById(id, cardNumberId, firstName, lastName, wareHouseId)

	if err != nil {
		return Employee{}, err
	}

	return e, nil
}

func (s *service) DeleteById(id int) error {
	err := s.repository.DeleteById(id)

	if err != nil {
		return err
	}

	return nil
}
