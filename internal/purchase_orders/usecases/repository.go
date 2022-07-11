package usecases

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/domain"
)

type PurchaseOrderRepository interface {
	Create(orderNumber string, orderDate string, trackingCode string, buyerId int, productRecordId int, orderStatusId int) (domain.Purchase_Order, error)
}