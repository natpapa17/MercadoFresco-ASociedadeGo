package sellers_test

import (
	"errors"
	"fmt"

	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestGetAll(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	s := sellers.Seller{
		Id:    1,
		Cid:  1,
		CompanyName:  "None",
		Address: "none",
		Telephone: "00000",
	}

	sList := make([]sellers.Seller, 0)
	sList = append(sList, s)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetAll").Return(sList, nil).Once()

		sl := sellers.NewService(mockRepo)
		list, err := sl.GetAll()

		assert.NoError(t, err)

		assert.Equal(t,"None", list[0].CompanyName)

		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetAll").
			Return(nil, errors.New("failed to retrieve products")).
			Once()

		s := sellers.NewService(mockRepo)
		_, err := s.GetAll()

		assert.NotNil(t, err)

		mockRepo.AssertExpectations(t)
	})
}


func TestDelete(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	s := sellers.NewService(mockRepo)
	t.Run("Should call Delete from seller repository with correct id", func(t *testing.T) {
		mockRepo.
			On("Delete", mock.AnythingOfType("int")).
			Return(nil).
			Once()

		s.Delete(1)

		mockRepo.AssertCalled(t, "Delete", 1)
	})

	t.Run("Should return an error if Delete from seller repository returns an error", func(t *testing.T) {
		mockRepo.
			On("Delete", mock.AnythingOfType("int")).
			Return(errors.New("any_error")).
			Once()

		err := s.Delete(1)

		assert.EqualError(t, err, "any_error")
	})

	t.Run("Should return nil on success", func(t *testing.T) {
		mockRepo.
			On("Delete", mock.AnythingOfType("int")).
			Return(nil).
			Once()

		err := s.Delete(1)

		assert.Nil(t, err)
	})
}


func TestStore(t *testing.T){
	mockRepo := mocks.NewRepository(t)
	expectSeller := sellers.Seller{
		Id:    1,
		Cid:  1,
		CompanyName:  "None",
		Address: "none",
		Telephone: "00000",
	}

	t.Run("crete_ok: Se contiver os campos necessários, será criado", func(t *testing.T) {
		
		mockRepo.On("Store", int(219), "Name", "Addres", "telephone").Return(expectSeller, nil)

		service := sellers.NewService(mockRepo)

		result, _ := service.Store(int(219), "Name", "Addres", "telephone")

		assert.Equal(t, expectSeller, result)
	})

	t.Run("create_conflict: Se o Cid já existir, ele não pode ser criado", func(t *testing.T) {

		

		mockRepo.On("Store", int(219), "Name", "Addres", "telephone").Return(expectSeller, fmt.Errorf("Card number id is not unique."))

		service := sellers.NewService(mockRepo)

		seller, err := service.Store(int(219), "Name", "Addres", "telephone")

		assert.NotNil(t, err)
		assert.Empty(t, seller)
		assert.Equal(t, err.Error(), "Cid is not unique.")

	})
	

}