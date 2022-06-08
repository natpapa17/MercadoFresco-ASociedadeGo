package employee

import (
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/pkg/store"
)

type employeeInterface interface {
	Create(cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error)
	GetAll() ([]Employee, error)
	GetById(id int) (Employee, error)
	UpdateById(id int, cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error)
	DeleteById(id int) error
	GetByCardNumberId(cardNumberId int) (Employee, error)
}

type repository struct {
	file store.Store
}

func CreateRepository(file store.Store) employeeInterface {
	return &repository{
		file: file,
	}
}

func (r *repository) GetByCardNumberId(cardNumberId int) (Employee, error) {
	var es []Employee
	if err := r.file.Read(&es); err != nil {
		return Employee{}, nil
	}

	result, found := Employee{}, false
	for _, e := range es {
		if e.Card_number_id == cardNumberId {
			result, found = e, true
			break
		}
	}

	if !found {
		return Employee{}, &NoElementInFileError{errors.New("can't find element with this cardNumberId")}
	}

	return result, nil
}

func (r *repository) Create(cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error) {

	var es []Employee
	if err := r.file.Read(&es); err != nil {
		return Employee{}, err
	}

	lastId, _ := r.lastId()

	e := Employee{
		Id:             lastId + 1,
		Card_number_id: cardNumberId,
		First_name:     firstName,
		Last_name:      lastName,
		Warehouse_id:   wareHouseId,
	}
	es = append(es, e)

	if err := r.file.Write(es); err != nil {
		return Employee{}, err
	}
	return e, nil
}

func (r *repository) GetAll() ([]Employee, error) {
	var es []Employee
	if err := r.file.Read(&es); err != nil {
		return []Employee{}, nil
	}
	return es, nil
}

func (r *repository) GetById(id int) (Employee, error) {
	var ws []Employee
	if err := r.file.Read(&ws); err != nil {
		return Employee{}, nil
	}

	result, found := Employee{}, false
	for _, w := range ws {
		if w.Id == id {
			result, found = w, true
			break
		}
	}

	if !found {
		return Employee{}, &NoElementInFileError{errors.New("can't find element with this id")}
	}

	return result, nil
}

func (r *repository) UpdateById(id int, cardNumberId int, firstName string, lastName string, wareHouseId int) (Employee, error) {
	var es []Employee
	if err := r.file.Read(&es); err != nil {
		return Employee{}, nil
	}

	result, updated := Employee{}, false
	for i, e := range es {
		if e.Id == id {
			es[i], updated = Employee{
				Id:             id,
				Card_number_id: cardNumberId,
				First_name:     firstName,
				Last_name:      lastName,
				Warehouse_id:   wareHouseId,
			}, true
			result = es[i]
			break
		}
	}

	if !updated {
		return Employee{}, &NoElementInFileError{errors.New("can't find element with this id")}
	}

	if err := r.file.Write(es); err != nil {
		return Employee{}, err
	}

	return result, nil
}

func (r *repository) DeleteById(id int) error {
	var es []Employee
	if err := r.file.Read(&es); err != nil {
		return nil
	}

	deleted := false
	for i, w := range es {
		if w.Id == id {
			newEs := []Employee{}
			newEs = append(newEs, es[:i]...)
			newEs = append(newEs, es[i+1:]...)
			es = newEs
			deleted = true
			break
		}
	}

	if !deleted {
		return &NoElementInFileError{errors.New("can't find element with this id")}
	}

	if err := r.file.Write(es); err != nil {
		return err
	}

	return nil
}

func (r *repository) lastId() (int, error) {
	var es []Employee
	if err := r.file.Read(&es); err != nil {
		return 0, err
	}

	if len(es) == 0 {
		return 0, nil
	}

	return es[len(es)-1].Id, nil
}
