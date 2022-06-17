package sellers_test

import (
	"errors"
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers/mocks"
	"github.com/stretchr/testify/assert"
)


func TestGetAll(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)

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

		s := sellers.NewService(mockRepo)
		list, err := s.GetAll()

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