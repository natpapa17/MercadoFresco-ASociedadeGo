package employee_test

import (
	"errors"
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/employee"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/employee/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeCreateParams() (int, string, string, int) {
	return 123, "valid_name", "valid_last_name", 1
}

func makeUpdateByIdParams() (int, int, string, string, int) {
	return 1, 1234, "valid_name", "valid_last_name", 1
}

func makeEmployee() employee.Employee {
	return employee.Employee{
		Id:             1,
		Card_number_id: 123,
		First_name:     "valid_name",
		Last_name:      "valid_last_name",
		Warehouse_id:   1,
	}
}

func makeUpdatedEmployee() employee.Employee {
	return employee.Employee{
		Id:             1,
		Card_number_id: 1234,
		First_name:     "valid_name",
		Last_name:      "valid_last_name",
		Warehouse_id:   1,
	}
}

func makeEmployeeWareHouse() employee.Warehouse {
	return employee.Warehouse{
		Id:                 1,
		WarehouseCode:      "123",
		Address:            "rua dos testes",
		Telephone:          "(11) 9999-9999",
		MinimumCapacity:    10,
		MinimumTemperature: 15.5,
	}
}

func TestCreate(t *testing.T) {
	mockEmployeeRepository := mocks.NewEmployeeRepositoryInterface(t)
	mockWarehouseRepository := mocks.NewWareHouseRepository(t)
	sut := employee.CreateService(mockEmployeeRepository, mockWarehouseRepository)

	t.Run("Should call GetByCardNumberId from Employee Repository with correct CardNumberId", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetByCardNumberId", mock.AnythingOfType("int")).
			Return(employee.Employee{}, &employee.NoElementInFileError{errors.New("can't find element with this id")}).
			Once()
		mockEmployeeRepository.
			On("Create", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeEmployee(), nil).
			Once()
		mockWarehouseRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(makeEmployeeWareHouse(), nil).
			Once()

		sut.Create(makeCreateParams())

		mockEmployeeRepository.AssertCalled(t, "GetByCardNumberId", 123)
	})

	t.Run("Should return error if CardNumberId provided is in use", func(t *testing.T) {
		mockEmployeeRepository.On("GetByCardNumberId", mock.AnythingOfType("int")).Return(makeEmployee(), nil).Once()
		mockWarehouseRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(makeEmployeeWareHouse(), nil).
			Once()

		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "This Card Number id is already In Use!")
	})

	t.Run("Should call Create from Employee Repository with correct values", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetByCardNumberId", mock.AnythingOfType("int")).
			Return(employee.Employee{}, &employee.NoElementInFileError{errors.New("can't find element with this id")}).
			Once()
		mockEmployeeRepository.
			On("Create", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeEmployee(), nil).Once()
		mockWarehouseRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(makeEmployeeWareHouse(), nil).
			Once()

		sut.Create(makeCreateParams())

		mockEmployeeRepository.AssertCalled(t, "Create", 123, "valid_name", "valid_last_name", 1)
	})

	t.Run("Should return error if Create from Employee Repository returns an error", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetByCardNumberId", mock.AnythingOfType("int")).
			Return(employee.Employee{}, &employee.NoElementInFileError{errors.New("can't find element with this id")}).
			Once()
		mockEmployeeRepository.
			On("Create", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(employee.Employee{}, errors.New("any_error")).Once()
		mockWarehouseRepository.On("GetById", mock.AnythingOfType("int")).
			Return(makeEmployeeWareHouse(), nil).
			Once()
		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Employee on success", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetByCardNumberId", mock.AnythingOfType("int")).
			Return(employee.Employee{}, &employee.NoElementInFileError{errors.New("can't find element with this id")}).
			Once()
		mockEmployeeRepository.
			On("Create", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeEmployee(), nil).
			Once()
		mockWarehouseRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(makeEmployeeWareHouse(), nil).
			Once()

		w, err := sut.Create(makeCreateParams())

		assert.Equal(t, makeEmployee(), w)
		assert.Nil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	mockEmployeeRepository := mocks.NewEmployeeRepositoryInterface(t)
	mockWarehouseRepository := mocks.NewWareHouseRepository(t)
	sut := employee.CreateService(mockEmployeeRepository, mockWarehouseRepository)

	t.Run("Should call GetAll from Employee Repository", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetAll").
			Return([]employee.Employee{makeEmployee()}, nil).
			Once()

		sut.GetAll()

		mockEmployeeRepository.AssertCalled(t, "GetAll")
	})

	t.Run("Should return an error if GetAll from Employee Repository returns an error", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetAll").
			Return([]employee.Employee{}, errors.New("any_error")).
			Once()

		_, err := sut.GetAll()

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an slice of Employee on success", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetAll").
			Return([]employee.Employee{makeEmployee()}, nil).
			Once()

		ws, err := sut.GetAll()

		assert.Equal(t, []employee.Employee{makeEmployee()}, ws)
		assert.Nil(t, err)
	})
}

func TestGetById(t *testing.T) {
	mockEmployeeRepository := mocks.NewEmployeeRepositoryInterface(t)
	mockWarehouseRepository := mocks.NewWareHouseRepository(t)
	sut := employee.CreateService(mockEmployeeRepository, mockWarehouseRepository)

	t.Run("Should call GetById from Employee Repository with correct id", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(makeEmployee(), nil).
			Once()

		sut.GetById(1)

		mockEmployeeRepository.AssertCalled(t, "GetById", 1)
	})

	t.Run("Should return an error if GetById from Employee Repository returns an error", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(employee.Employee{}, errors.New("any_error")).
			Once()

		_, err := sut.GetById(1)

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Employee on success", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(makeEmployee(), nil).
			Once()

		w, err := sut.GetById(1)

		assert.Equal(t, makeEmployee(), w)
		assert.Nil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	mockEmployeeRepository := mocks.NewEmployeeRepositoryInterface(t)
	mockWarehouseRepository := mocks.NewWareHouseRepository(t)
	sut := employee.CreateService(mockEmployeeRepository, mockWarehouseRepository)

	t.Run("Should call GetByCardNumberId from Employee Repository with correct CardNumberId", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetByCardNumberId", mock.AnythingOfType("int")).
			Return(employee.Employee{}, &employee.NoElementInFileError{errors.New("can't find element with this id")}).
			Once()
		mockEmployeeRepository.
			On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeUpdatedEmployee(), nil).
			Once()

		sut.UpdateById(makeUpdateByIdParams())

		mockEmployeeRepository.AssertCalled(t, "GetByCardNumberId", 1234)
	})

	t.Run("Should return error if new employee card number id provided is in use", func(t *testing.T) {
		dbEmployee := makeEmployee()
		dbEmployee.Id = 3
		mockEmployeeRepository.On("GetByCardNumberId", mock.AnythingOfType("int")).Return(dbEmployee, nil).Once()

		_, err := sut.UpdateById(makeUpdateByIdParams())

		assert.EqualError(t, err, "this CardNumberId is already in use")
	})

	t.Run("Should call UpdateById from Employee Repository with correct values", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetByCardNumberId", mock.AnythingOfType("int")).
			Return(employee.Employee{}, &employee.NoElementInFileError{errors.New("can't find element with this id")}).
			Once()
		mockEmployeeRepository.
			On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeUpdatedEmployee(), nil).Once()

		sut.UpdateById(makeUpdateByIdParams())

		mockEmployeeRepository.AssertCalled(t, "UpdateById", 1, 1234, "valid_name", "valid_last_name", 1)
	})

	t.Run("Should return error if UpdateById from Employee Repository returns an error", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetByCardNumberId", mock.AnythingOfType("int")).
			Return(employee.Employee{}, &employee.NoElementInFileError{errors.New("can't find element with this id")}).
			Once()
		mockEmployeeRepository.
			On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(employee.Employee{}, errors.New("any_error")).Once()
		_, err := sut.UpdateById(makeUpdateByIdParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Updated Employee on success", func(t *testing.T) {
		mockEmployeeRepository.
			On("GetByCardNumberId", mock.AnythingOfType("int")).
			Return(employee.Employee{}, &employee.NoElementInFileError{errors.New("can't find element with this id")}).
			Once()
		mockEmployeeRepository.
			On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeUpdatedEmployee(), nil).
			Once()

		w, err := sut.UpdateById(makeUpdateByIdParams())

		assert.Equal(t, makeUpdatedEmployee(), w)
		assert.Nil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	mockEmployeeRepository := mocks.NewEmployeeRepositoryInterface(t)
	mockWarehouseRepository := mocks.NewWareHouseRepository(t)
	sut := employee.CreateService(mockEmployeeRepository, mockWarehouseRepository)

	t.Run("Should call DeleteById from Employee Repository with correct id", func(t *testing.T) {
		mockEmployeeRepository.
			On("DeleteById", mock.AnythingOfType("int")).
			Return(nil).
			Once()

		sut.DeleteById(1)

		mockEmployeeRepository.AssertCalled(t, "DeleteById", 1)
	})

	t.Run("Should return an error if DeleteById from Employee Repository returns an error", func(t *testing.T) {
		mockEmployeeRepository.
			On("DeleteById", mock.AnythingOfType("int")).
			Return(errors.New("any_error")).
			Once()

		err := sut.DeleteById(1)

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return nil on success", func(t *testing.T) {
		mockEmployeeRepository.
			On("DeleteById", mock.AnythingOfType("int")).
			Return(nil).
			Once()

		err := sut.DeleteById(1)

		assert.Nil(t, err)
	})
}
