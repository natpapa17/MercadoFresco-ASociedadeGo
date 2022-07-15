package usecases

import "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/domain"

type PurchaseOrderService interface {
	Create(orderNumber string, orderDate string, trackingCode string, buyerId int, productRecordId int, orderStatusId int) (domain.Purchase_Order, error)
}

type purchaseOrderService struct {
	purchaseOrderRepository PurchaseOrderRepository
}

func CreatePurchaseOrderService(r PurchaseOrderRepository) PurchaseOrderService {
	return &purchaseOrderService{
		purchaseOrderRepository: r,
	}
}

func (s *purchaseOrderService) Create(orderNumber string, orderDate string, trackingCode string, buyerId int, productRecordId int, orderStatusId int) (domain.Purchase_Order, error) {
	order, err := s.purchaseOrderRepository.Create(orderNumber, orderDate, trackingCode, buyerId, productRecordId, orderStatusId)

	if err != nil {
		return domain.Purchase_Order{}, err
	}

	return order, nil
}

// func (s *purchaseOrderService) GetPurchaseOrderById()
