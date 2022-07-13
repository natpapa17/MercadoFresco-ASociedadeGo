package usecases_test

import (
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeCreateParams() (string, string, string, int, int, int) {
	return "123", "01-01-2022", "123", 1, 1, 1
}

func makePurchaseOrder() domain.Purchase_Order {
	return domain.Purchase_Order{
		ID:              1,
		OrderNumber:     "123",
		OrderDate:       "01-01-2022",
		TrackingCode:    "123",
		BuyerId:         1,
		ProductRecordId: 1,
		OrderStatusId:   1,
	}
}

func TestCreate(t *testing.T) {
	mockPurchaseOrderRepository := mocks.NewPurchaseOrderRepository(t)
	service := usecases.CreatePurchaseOrderService(mockPurchaseOrderRepository)

	t.Run("create_ok", func(t *testing.T) {
		mockPurchaseOrderRepository.
			On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(makePurchaseOrder(), nil).
			Once()

		p, err := service.Create(makeCreateParams())

		assert.Equal(t, makePurchaseOrder(), p)
		assert.Nil(t, err)
	})
}
