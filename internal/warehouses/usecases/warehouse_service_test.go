package usecases_test

import (
	"errors"
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeCreateParams() (string, string, string, int, float64) {
	return "valid_code", "valid_address", "(99) 99999-9999", 10, 5.0
}

func makeUpdateByIdParams() (int, string, string, string, int, float64) {
	return 1, "valid_code", "updated_address", "(99) 99999-9999", 20, 15.0
}

func makeWarehouse() domain.Warehouse {
	return domain.Warehouse{
		Id:                 1,
		WarehouseCode:      "valid_code",
		Address:            "valid_address",
		Telephone:          "(99) 99999-9999",
		MinimumCapacity:    10,
		MinimumTemperature: 5.0,
	}
}

func makeUpdatedWarehouse() domain.Warehouse {
	return domain.Warehouse{
		Id:                 1,
		WarehouseCode:      "valid_code",
		Address:            "updated_address",
		Telephone:          "(99) 99999-9999",
		MinimumCapacity:    20,
		MinimumTemperature: 15.0,
	}
}

func TestCreate(t *testing.T) {
	mockWarehouseRepository := mocks.NewWarehouseRepository(t)
	sut := usecases.CreateWarehouseService(mockWarehouseRepository)

	t.Run("Should call GetByWarehouseCode from Warehouse Repository with correct code", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(domain.Warehouse{}, usecases.ErrNoElementFound).
			Once()
		mockWarehouseRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).
			Return(makeWarehouse(), nil).
			Once()

		sut.Create(makeCreateParams())

		mockWarehouseRepository.AssertCalled(t, "GetByWarehouseCode", "valid_code")
	})

	t.Run("Should return error if warehouse code provided is in use", func(t *testing.T) {
		mockWarehouseRepository.On("GetByWarehouseCode", mock.AnythingOfType("string")).Return(makeWarehouse(), nil).Once()

		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "this warehouse_code is already in use")
	})

	t.Run("Should return error if GetByWarehouseCode from Warehouse Repository returns an error other than noElementInFile", func(t *testing.T) {
		mockWarehouseRepository.On("GetByWarehouseCode", mock.AnythingOfType("string")).Return(domain.Warehouse{}, errors.New("any_error")).Once()

		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should call Create from Warehouse Repository with correct values", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(domain.Warehouse{}, usecases.ErrNoElementFound).
			Once()
		mockWarehouseRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).
			Return(makeWarehouse(), nil).Once()

		sut.Create(makeCreateParams())

		mockWarehouseRepository.AssertCalled(t, "Create", "valid_code", "valid_address", "(99) 99999-9999", 10, 5.0)
	})

	t.Run("Should return error if Create from Warehouse Repository returns an error", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(domain.Warehouse{}, usecases.ErrNoElementFound).
			Once()
		mockWarehouseRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).
			Return(domain.Warehouse{}, errors.New("any_error")).Once()
		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Warehouse on success", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(domain.Warehouse{}, usecases.ErrNoElementFound).
			Once()
		mockWarehouseRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).
			Return(makeWarehouse(), nil).
			Once()

		w, err := sut.Create(makeCreateParams())

		assert.Equal(t, makeWarehouse(), w)
		assert.Nil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	mockWarehouseRepository := mocks.NewWarehouseRepository(t)
	sut := usecases.CreateWarehouseService(mockWarehouseRepository)

	t.Run("Should call GetAll from Warehouse Repository", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetAll").
			Return(domain.Warehouses{makeWarehouse()}, nil).
			Once()

		sut.GetAll()

		mockWarehouseRepository.AssertCalled(t, "GetAll")
	})

	t.Run("Should return an error if GetAll from Warehouse Repository returns an error", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetAll").
			Return(domain.Warehouses{}, errors.New("any_error")).
			Once()

		_, err := sut.GetAll()

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an slice of Warehouses on success", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetAll").
			Return(domain.Warehouses{makeWarehouse()}, nil).
			Once()

		ws, err := sut.GetAll()

		assert.Equal(t, domain.Warehouses{makeWarehouse()}, ws)
		assert.Nil(t, err)
	})
}

func TestGetById(t *testing.T) {
	mockWarehouseRepository := mocks.NewWarehouseRepository(t)
	sut := usecases.CreateWarehouseService(mockWarehouseRepository)

	t.Run("Should call GetById from Warehouse Repository with correct id", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(makeWarehouse(), nil).
			Once()

		sut.GetById(1)

		mockWarehouseRepository.AssertCalled(t, "GetById", 1)
	})

	t.Run("Should return an error if GetById from Warehouse Repository returns an error", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(domain.Warehouse{}, errors.New("any_error")).
			Once()

		_, err := sut.GetById(1)

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Warehouses on success", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(makeWarehouse(), nil).
			Once()

		w, err := sut.GetById(1)

		assert.Equal(t, makeWarehouse(), w)
		assert.Nil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	mockWarehouseRepository := mocks.NewWarehouseRepository(t)
	sut := usecases.CreateWarehouseService(mockWarehouseRepository)

	t.Run("Should call GetByWarehouseCode from Warehouse Repository with correct code", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(domain.Warehouse{}, usecases.ErrNoElementFound).
			Once()
		mockWarehouseRepository.
			On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).
			Return(makeUpdatedWarehouse(), nil).
			Once()

		sut.UpdateById(makeUpdateByIdParams())

		mockWarehouseRepository.AssertCalled(t, "GetByWarehouseCode", "valid_code")
	})

	t.Run("Should return error if new warehouse code provided is in use", func(t *testing.T) {
		dbWarehouse := makeWarehouse()
		dbWarehouse.Id = 3
		mockWarehouseRepository.On("GetByWarehouseCode", mock.AnythingOfType("string")).Return(dbWarehouse, nil).Once()

		_, err := sut.UpdateById(makeUpdateByIdParams())

		assert.EqualError(t, err, "this warehouse_code is already in use")
	})

	t.Run("Should return error if GetByWarehouseCode from Warehouse Repository returns an error other than noElementInFile", func(t *testing.T) {
		mockWarehouseRepository.On("GetByWarehouseCode", mock.AnythingOfType("string")).Return(domain.Warehouse{}, errors.New("any_error")).Once()

		_, err := sut.UpdateById(makeUpdateByIdParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should call UpdateById from Warehouse Repository with correct values", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(domain.Warehouse{}, usecases.ErrNoElementFound).
			Once()
		mockWarehouseRepository.
			On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).
			Return(makeUpdatedWarehouse(), nil).Once()

		sut.UpdateById(makeUpdateByIdParams())

		mockWarehouseRepository.AssertCalled(t, "UpdateById", 1, "valid_code", "updated_address", "(99) 99999-9999", 20, 15.0)
	})

	t.Run("Should return error if UpdateById from Warehouse Repository returns an error", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(domain.Warehouse{}, usecases.ErrNoElementFound).
			Once()
		mockWarehouseRepository.
			On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).
			Return(domain.Warehouse{}, errors.New("any_error")).Once()
		_, err := sut.UpdateById(makeUpdateByIdParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Updated Warehouse on success", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(domain.Warehouse{}, usecases.ErrNoElementFound).
			Once()
		mockWarehouseRepository.
			On("UpdateById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).
			Return(makeUpdatedWarehouse(), nil).
			Once()

		w, err := sut.UpdateById(makeUpdateByIdParams())

		assert.Equal(t, makeUpdatedWarehouse(), w)
		assert.Nil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	mockWarehouseRepository := mocks.NewWarehouseRepository(t)
	sut := usecases.CreateWarehouseService(mockWarehouseRepository)

	t.Run("Should call DeleteById from Warehouse Repository with correct id", func(t *testing.T) {
		mockWarehouseRepository.
			On("DeleteById", mock.AnythingOfType("int")).
			Return(nil).
			Once()

		sut.DeleteById(1)

		mockWarehouseRepository.AssertCalled(t, "DeleteById", 1)
	})

	t.Run("Should return an error if DeleteById from Warehouse Repository returns an error", func(t *testing.T) {
		mockWarehouseRepository.
			On("DeleteById", mock.AnythingOfType("int")).
			Return(errors.New("any_error")).
			Once()

		err := sut.DeleteById(1)

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return nil on success", func(t *testing.T) {
		mockWarehouseRepository.
			On("DeleteById", mock.AnythingOfType("int")).
			Return(nil).
			Once()

		err := sut.DeleteById(1)

		assert.Nil(t, err)
	})
}
