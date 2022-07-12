package usecases_test

import (
	"errors"
	"testing"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeCreateParams() (string, string, string, int, int, int) {
	return "123", "01-01-2022", "123", 1, 1, 1
}

func makePurchaseOrder() purchaseOrders.PurchaseOrder {
	return purchaseOrders.PurchaseOrder{
		ID:             1,
		OrderNumber: "123",
		OrderDate: "01-01-2022",
		TrackingCode: "123",
		BuyerId: 1,
		ProductRecordId: 1,
		OrderStatusId: 1,
	}
}

func TestCreate(t *testing.T) {
	mockPurchaseOrderRepository := mocks.NewRepository(t)
	service := purchaseOrder.CreatePurchaseOrderService(mockPurchaseOrderRepository)

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