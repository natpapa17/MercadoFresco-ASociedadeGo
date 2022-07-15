package usecases_test

import (
	"errors"
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/carriers/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeCarrier() domain.Carrier {
	return domain.Carrier{
		Id:          1,
		Cid:         "valid_cid",
		CompanyName: "valid_name",
		Address:     "valid_address",
		Telephone:   "valid_phone",
		LocalityId:  1,
	}
}

func makeLocality() domain.Locality {
	return domain.Locality{
		Id:         1,
		Name:       "valid_name",
		ProvinceId: 1,
	}
}

func makeLocalities() domain.Localities {
	return domain.Localities{
		domain.Locality{
			Id:         1,
			Name:       "valid_name_1",
			ProvinceId: 1,
		},
		domain.Locality{
			Id:         2,
			Name:       "valid_name_2",
			ProvinceId: 1,
		},
		domain.Locality{
			Id:         3,
			Name:       "valid_name_3",
			ProvinceId: 1,
		},
	}
}

func TestCreate(t *testing.T) {
	makeSut := func() (usecases.CarrierService, *mocks.CarrierRepository, *mocks.LocalityRepository) {
		mockCarrierRepository := mocks.NewCarrierRepository(t)
		mockLocalityRepository := mocks.NewLocalityRepository(t)
		sut := usecases.CreateCarrierService(mockCarrierRepository, mockLocalityRepository)
		return sut, mockCarrierRepository, mockLocalityRepository
	}

	makeCreateParams := func() (string, string, string, string, int) {
		return "valid_cid", "valid_name", "valid_address", "valid_phone", 1
	}

	t.Run("Should call GetByCid from Carrier Repository with correct cid", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockCarrierRepository.
			On("GetByCid", mock.AnythingOfType("string")).
			Return(domain.Carrier{}, usecases.ErrNoElementFound).
			Once()
		mockLocalityRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(domain.Locality{}, nil).
			Once()
		mockCarrierRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeCarrier(), nil).
			Once()

		sut.Create(makeCreateParams())

		mockCarrierRepository.AssertCalled(t, "GetByCid", "valid_cid")
	})

	t.Run("Should return error if cid provided is in use", func(t *testing.T) {
		sut, mockCarrierRepository, _ := makeSut()
		mockCarrierRepository.On("GetByCid", mock.AnythingOfType("string")).Return(makeCarrier(), nil)

		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "this cid is in use")
	})

	t.Run("Should return error if GetByCid from Carrier Repository returns an error other than ErrNoElementFound", func(t *testing.T) {
		sut, mockCarrierRepository, _ := makeSut()
		mockCarrierRepository.On("GetByCid", mock.AnythingOfType("string")).Return(domain.Carrier{}, errors.New("any_error"))

		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should call GetById from Locality Repository with correct id", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockCarrierRepository.On("GetByCid", mock.AnythingOfType("string")).Return(domain.Carrier{}, nil)
		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(makeLocality(), nil)
		mockCarrierRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeCarrier(), nil)
		sut.Create(makeCreateParams())

		mockLocalityRepository.AssertCalled(t, "GetById", 1)
	})

	t.Run("Should return ErrInvalidLocalityId if locality_id provided is in invalid", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockCarrierRepository.On("GetByCid", mock.AnythingOfType("string")).Return(domain.Carrier{}, nil)
		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(domain.Locality{}, usecases.ErrNoElementFound)

		_, err := sut.Create(makeCreateParams())

		assert.Equal(t, usecases.ErrInvalidLocalityId, err)
	})

	t.Run("Should return error if GetById from Locality Repository returns an error other than ErrNoElementFound", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockCarrierRepository.On("GetByCid", mock.AnythingOfType("string")).Return(domain.Carrier{}, nil)
		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(domain.Locality{}, errors.New("any_error"))

		_, err := sut.Create(makeCreateParams())

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should call Create from Carrier Repository with correct values", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockCarrierRepository.
			On("GetByCid", mock.AnythingOfType("string")).
			Return(domain.Carrier{}, usecases.ErrNoElementFound).
			Once()
		mockLocalityRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(domain.Locality{}, nil).
			Once()
		mockCarrierRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeCarrier(), nil).
			Once()

		sut.Create(makeCreateParams())

		mockCarrierRepository.AssertCalled(t, "Create", "valid_cid", "valid_name", "valid_address", "valid_phone", 1)
	})

	t.Run("Should return error if Create from Carrier Repository returns an error", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockCarrierRepository.
			On("GetByCid", mock.AnythingOfType("string")).
			Return(domain.Carrier{}, usecases.ErrNoElementFound)
		mockLocalityRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(domain.Locality{}, nil)
		mockCarrierRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(domain.Carrier{}, errors.New("any_error"))

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, domain.Carrier{}, result)
		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an Carrier on success", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockCarrierRepository.
			On("GetByCid", mock.AnythingOfType("string")).
			Return(domain.Carrier{}, usecases.ErrNoElementFound).
			Once()
		mockLocalityRepository.
			On("GetById", mock.AnythingOfType("int")).
			Return(domain.Locality{}, nil).
			Once()
		mockCarrierRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).
			Return(makeCarrier(), nil).
			Once()

		result, err := sut.Create(makeCreateParams())

		assert.Equal(t, makeCarrier(), result)
		assert.Nil(t, err)
	})
}

func TestGetNumberOfCarriersPerLocalities(t *testing.T) {
	makeSut := func() (usecases.CarrierService, *mocks.CarrierRepository, *mocks.LocalityRepository) {
		mockCarrierRepository := mocks.NewCarrierRepository(t)
		mockLocalityRepository := mocks.NewLocalityRepository(t)
		sut := usecases.CreateCarrierService(mockCarrierRepository, mockLocalityRepository)
		return sut, mockCarrierRepository, mockLocalityRepository
	}

	t.Run("Should call GetById from Locality Repository for each id received", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(makeLocality(), nil)
		mockCarrierRepository.On("GetNumberOfCarriersPerLocality", mock.AnythingOfType("int")).Return(2, nil)

		sut.GetNumberOfCarriersPerLocalities([]int{1, 2, 3})

		mockLocalityRepository.AssertNumberOfCalls(t, "GetById", 3)
		mockLocalityRepository.AssertCalled(t, "GetById", 1)
		mockLocalityRepository.AssertCalled(t, "GetById", 2)
		mockLocalityRepository.AssertCalled(t, "GetById", 3)
	})

	t.Run("Should return an error if GetById from Locality Repository returns an error different than ErrNoElementFound", func(t *testing.T) {
		sut, _, mockLocalityRepository := makeSut()

		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(domain.Locality{}, errors.New("any_error"))

		result, err := sut.GetNumberOfCarriersPerLocalities([]int{1, 2, 3})

		assert.Equal(t, domain.ReportsNumberOfCarriersPerLocality{}, result)
		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should call GetNumberOfCarriersPerLocality from Carrier Repository  for each id received", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(makeLocality(), nil)
		mockCarrierRepository.On("GetNumberOfCarriersPerLocality", mock.AnythingOfType("int")).Return(2, nil)

		sut.GetNumberOfCarriersPerLocalities([]int{1, 2, 3})

		mockCarrierRepository.AssertNumberOfCalls(t, "GetNumberOfCarriersPerLocality", 3)
		mockCarrierRepository.AssertCalled(t, "GetNumberOfCarriersPerLocality", 1)
		mockCarrierRepository.AssertCalled(t, "GetNumberOfCarriersPerLocality", 2)
		mockCarrierRepository.AssertCalled(t, "GetNumberOfCarriersPerLocality", 3)
	})

	t.Run("Should return an error if GetNumberOfCarriersPerLocality from Carrier Repository returns an error", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()

		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(makeLocality(), nil)
		mockCarrierRepository.On("GetNumberOfCarriersPerLocality", mock.AnythingOfType("int")).Return(0, errors.New("any_error"))

		result, err := sut.GetNumberOfCarriersPerLocalities([]int{1, 2, 3})

		assert.Equal(t, domain.ReportsNumberOfCarriersPerLocality{}, result)
		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an report list on success", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()

		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(makeLocality(), nil)
		mockCarrierRepository.On("GetNumberOfCarriersPerLocality", mock.AnythingOfType("int")).Return(3, nil)

		result, err := sut.GetNumberOfCarriersPerLocalities([]int{1, 2})

		expected := domain.ReportsNumberOfCarriersPerLocality{
			domain.ReportNumberOfCarriersPerLocality{
				LocalityId:    1,
				LocalityName:  "valid_name",
				CarriersCount: 3,
			},
			domain.ReportNumberOfCarriersPerLocality{
				LocalityId:    1,
				LocalityName:  "valid_name",
				CarriersCount: 3,
			},
		}

		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})

	t.Run("Should return an report list if GetById from Locality Repository returns ErrNoElementFound", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(makeLocality(), nil).Times(1)
		mockLocalityRepository.On("GetById", mock.AnythingOfType("int")).Return(domain.Locality{}, usecases.ErrNoElementFound).Times(2)
		mockCarrierRepository.On("GetNumberOfCarriersPerLocality", mock.AnythingOfType("int")).Return(2, nil)

		result, err := sut.GetNumberOfCarriersPerLocalities([]int{1, 2, 3})

		mockLocalityRepository.AssertNumberOfCalls(t, "GetById", 3)
		mockCarrierRepository.AssertNumberOfCalls(t, "GetNumberOfCarriersPerLocality", 1)

		assert.Equal(t, 1, len(result))
		assert.Nil(t, err)
	})
}

func TestGetAllNumberOfCarriersPerLocalities(t *testing.T) {
	makeSut := func() (usecases.CarrierService, *mocks.CarrierRepository, *mocks.LocalityRepository) {
		mockCarrierRepository := mocks.NewCarrierRepository(t)
		mockLocalityRepository := mocks.NewLocalityRepository(t)
		sut := usecases.CreateCarrierService(mockCarrierRepository, mockLocalityRepository)
		return sut, mockCarrierRepository, mockLocalityRepository
	}

	t.Run("Should call GetAll from Locality Repository", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockLocalityRepository.On("GetAll").Return(makeLocalities(), nil)
		mockCarrierRepository.On("GetNumberOfCarriersPerLocality", mock.AnythingOfType("int")).Return(2, nil)

		sut.GetAllNumberOfCarriersPerLocality()

		mockLocalityRepository.AssertCalled(t, "GetAll")
		mockLocalityRepository.AssertNumberOfCalls(t, "GetAll", 1)

	})

	t.Run("Should return an error if GetAll from Locality Repository returns an error", func(t *testing.T) {
		sut, _, mockLocalityRepository := makeSut()
		mockLocalityRepository.On("GetAll").Return(domain.Localities{}, errors.New("any_error"))

		result, err := sut.GetAllNumberOfCarriersPerLocality()

		assert.Equal(t, domain.ReportsNumberOfCarriersPerLocality{}, result)
		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should call GetNumberOfCarriersPerLocality from Carrier Repository  for each id received", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockLocalityRepository.On("GetAll").Return(makeLocalities(), nil)
		mockCarrierRepository.On("GetNumberOfCarriersPerLocality", mock.AnythingOfType("int")).Return(2, nil)

		sut.GetAllNumberOfCarriersPerLocality()

		mockCarrierRepository.AssertNumberOfCalls(t, "GetNumberOfCarriersPerLocality", 3)
		mockCarrierRepository.AssertCalled(t, "GetNumberOfCarriersPerLocality", 1)
		mockCarrierRepository.AssertCalled(t, "GetNumberOfCarriersPerLocality", 2)
		mockCarrierRepository.AssertCalled(t, "GetNumberOfCarriersPerLocality", 3)
	})

	t.Run("Should return an error if GetNumberOfCarriersPerLocality from Carrier Repository returns an error", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockLocalityRepository.On("GetAll").Return(makeLocalities(), nil)
		mockCarrierRepository.On("GetNumberOfCarriersPerLocality", mock.AnythingOfType("int")).Return(0, errors.New("any_error"))

		result, err := sut.GetAllNumberOfCarriersPerLocality()

		assert.Equal(t, domain.ReportsNumberOfCarriersPerLocality{}, result)
		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return an report list on success", func(t *testing.T) {
		sut, mockCarrierRepository, mockLocalityRepository := makeSut()
		mockLocalityRepository.On("GetAll").Return(makeLocalities(), nil)
		mockCarrierRepository.On("GetNumberOfCarriersPerLocality", mock.AnythingOfType("int")).Return(3, nil)

		result, err := sut.GetAllNumberOfCarriersPerLocality()

		expected := domain.ReportsNumberOfCarriersPerLocality{
			domain.ReportNumberOfCarriersPerLocality{
				LocalityId:    1,
				LocalityName:  "valid_name_1",
				CarriersCount: 3,
			},
			domain.ReportNumberOfCarriersPerLocality{
				LocalityId:    2,
				LocalityName:  "valid_name_2",
				CarriersCount: 3,
			},
			domain.ReportNumberOfCarriersPerLocality{
				LocalityId:    3,
				LocalityName:  "valid_name_3",
				CarriersCount: 3,
			},
		}

		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})
}
