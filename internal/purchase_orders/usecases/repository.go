package usecases

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/domain"
)

type PurchaseOrderRepository interface {
	Create(OrderNumber string, OrderDate string, TrackingCode string, BuyerId int, ProductRecordId int, OrderStatusId int) (domain.Purchase_Order, error)
}