package sellers_test

import (
	"errors"
	

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

	t.Run("create_ok: if the fields are correct, the new seller will be stored", func(t *testing.T) {
		mockRepo.On("LastID").Return(1, nil).Once()
		mockRepo.On("Store",  mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(expectSeller, nil).Once()

		service := sellers.NewService(mockRepo)

		result, _ := service.Store(219, "Name", "Addres", "telephone")
		
		assert.Equal(t, expectSeller, result)
	})

}

func TestUpdate(t *testing.T){
	mockRepo := mocks.NewRepository(t)
	expectSeller := sellers.Seller{
		Id:    1,
		Cid:  1,
		CompanyName:  "None",
		Address: "none",
		Telephone: "00000",
	}
	t.Run("update_ok: return the updated information", func(t*testing.T){
	
		mockRepo.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(expectSeller, nil).Once()
		service := sellers.NewService(mockRepo)
		result, _ := service.Update( int(1),int(219), "Name", "Addres", "telephone")
		
		assert.Equal(t, expectSeller, result)

	})
}


func TestGetById(t *testing.T){
	mockRepo := mocks.NewRepository(t)

	expectedSellersList := []sellers.Seller{
		{
			Id:          1,
			Cid:         219,
			CompanyName: "Meta",
			Address:     " SP",
			Telephone:   "00000000",
		},
		{
			Id:          2,
			Cid:         422,
			CompanyName: "Herbalife",
			Address:     "None",
			Telephone:   "0000000",
		},
	}
	t.Run("find_by_id_existent: if the Id exists, it will return the element with its information", func(t *testing.T) {
		mockRepo.On("GetById", int(2)).Return(expectedSellersList[1], nil)
		service := sellers.NewService(mockRepo)
	
		result, err := service.GetById(int(2))
	
		assert.Nil(t, err)
		assert.Equal(t, expectedSellersList[1], result)



	})
	

}


