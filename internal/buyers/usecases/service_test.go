package usecases_test

import (
	"errors"
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeCreateParams() (string, string, string, string) {
	return "valid_first_name", "valid_last_name", "valid_address", "valid_document_number"
}

func makeUpdateParams() (int, string, string, string, string) {
	return 2, "updated_first_name", "updated_last_name", "updated_address", "updated_document_number"
}

func makeBuyer() domain.Buyer {
	return domain.Buyer{
		ID:             1,
		FirstName:      "valid_first_name",
		LastName:       "valid_last_name",
		Address:        "valid_address",
		DocumentNumber: "valid_document_number",
	}
}

func makeUpdateBuyer() domain.Buyer {
	return domain.Buyer{
		ID:             2,
		FirstName:      "updated_first_name",
		LastName:       "updated_last_name",
		Address:        "updated_address",
		DocumentNumber: "updated_document_number",
	}
}

func TestGetAll(t *testing.T) {
	mockBuyerRepository := mocks.NewRepository(t)
	service := usecases.CreateBuyerService(mockBuyerRepository)

	t.Run("find_all", func(t *testing.T) {
		mockBuyerRepository.
			On("GetAll").
			Return(domain.Buyers{makeBuyer()}, nil).
			Once()

		ps, err := service.GetAll()

		assert.Equal(t, domain.Buyers{makeBuyer()}, ps)
		assert.Nil(t, err)
	})
}

func TestGetBuyerById(t *testing.T) {
	mockBuyerRepository := mocks.NewRepository(t)
	service := usecases.CreateBuyerService(mockBuyerRepository)

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		mockBuyerRepository.On("GetBuyerById", mock.AnythingOfType("int")).
			Return(makeBuyer(), errors.New("Error")).Once()

		_, err := service.GetBuyerById(1)

		assert.EqualError(t, err, "Error")
	})

	t.Run("find_by_id_existent", func(t *testing.T) {
		mockBuyerRepository.On("GetBuyerById", mock.AnythingOfType("int")).Return(makeBuyer(), nil).Once()

		p, err := service.GetBuyerById(1)

		assert.Equal(t, makeBuyer(), p)
		assert.Nil(t, err)
	})
}

func TestCreate(t *testing.T) {
	mockBuyerRepository := mocks.NewRepository(t)
	service := usecases.CreateBuyerService(mockBuyerRepository)

	t.Run("create_ok", func(t *testing.T) {
		mockBuyerRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(makeBuyer(), nil).
			Once()

		p, err := service.Create(makeCreateParams())

		assert.Equal(t, makeBuyer(), p)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	mockBuyerRepository := mocks.NewRepository(t)
	service := usecases.CreateBuyerService(mockBuyerRepository)

	t.Run("update_ok", func(t *testing.T) {
		mockBuyerRepository.
			On("UpdateBuyerById", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(makeUpdateBuyer(), nil).
			Once()

		b, err := service.UpdateBuyerById(makeUpdateParams())

		assert.Equal(t, makeUpdateBuyer(), b)
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	mockBuyerRepository := mocks.NewRepository(t)
	service := usecases.CreateBuyerService(mockBuyerRepository)

	t.Run("delete_non_existent", func(t *testing.T) {
		mockBuyerRepository.
			On("DeleteBuyerById", mock.AnythingOfType("int")).Return(errors.New("Error")).Once()

		err := service.DeleteBuyerById(1)

		assert.EqualError(t, err, "Error")
	})

	t.Run("delete_ok", func(t *testing.T) {
		mockBuyerRepository.On("DeleteBuyerById", mock.AnythingOfType("int")).Return(nil)

		p := service.DeleteBuyerById(1)

		assert.Nil(t, p)
	})
}
