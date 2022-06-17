package warehouses_test

import (
	"errors"
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/warehouses/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeCreateParams() (string, string, string, int, float64) {
	return "valid_code", "valid_address", "(99) 99999-9999", 10, 5.0
}

func makeWarehouse() warehouses.Warehouse {
	return warehouses.Warehouse{
		Id:                 1,
		WarehouseCode:      "valid_code",
		Address:            "valid_address",
		Telephone:          "(99) 99999-9999",
		MinimumCapacity:    10,
		MinimumTemperature: 5.0,
	}
}

func TestCreate(t *testing.T) {
	mockWarehouseRepository := mocks.NewRepository(t)
	sut := warehouses.CreateService(mockWarehouseRepository)

	t.Run("Should call GetByWarehouseCode from Warehouse Repository with correct code", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(warehouses.Warehouse{}, &warehouses.NoElementInFileError{errors.New("can't find element with this id")}).
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
		mockWarehouseRepository.On("GetByWarehouseCode", mock.AnythingOfType("string")).Return(warehouses.Warehouse{}, errors.New("any_error")).Once()

		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should call Create from Warehouse Repository with correct values", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(warehouses.Warehouse{}, &warehouses.NoElementInFileError{errors.New("can't find element with this id")}).
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
			Return(warehouses.Warehouse{}, &warehouses.NoElementInFileError{errors.New("can't find element with this id")}).
			Once()
		mockWarehouseRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("float64")).
			Return(warehouses.Warehouse{}, errors.New("any_error")).Once()
		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Warehouse on success", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetByWarehouseCode", mock.AnythingOfType("string")).
			Return(warehouses.Warehouse{}, &warehouses.NoElementInFileError{errors.New("can't find element with this id")}).
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
	mockWarehouseRepository := mocks.NewRepository(t)
	sut := warehouses.CreateService(mockWarehouseRepository)

	t.Run("Should call GetAll from Warehouse Repository", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetAll").
			Return([]warehouses.Warehouse{makeWarehouse()}, nil).
			Once()

		sut.GetAll()

		mockWarehouseRepository.AssertCalled(t, "GetAll")
	})

	t.Run("Should return an error if GetAll from Warehouse Repository returns an error", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetAll").
			Return([]warehouses.Warehouse{}, errors.New("any_error")).
			Once()

		_, err := sut.GetAll()

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an slice of Warehouses on success", func(t *testing.T) {
		mockWarehouseRepository.
			On("GetAll").
			Return([]warehouses.Warehouse{makeWarehouse()}, nil).
			Once()

		ws, err := sut.GetAll()

		assert.Equal(t, []warehouses.Warehouse{makeWarehouse()}, ws)
		assert.Nil(t, err)
	})
}
