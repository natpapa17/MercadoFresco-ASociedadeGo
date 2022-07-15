package adapters

import (
	"database/sql"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases"
)

type purchaseOrderMySQLRepository struct {
	db *sql.DB
}

func CreatePurchaseOrderMySQLRepository(db *sql.DB) usecases.PurchaseOrderRepository {
	return &purchaseOrderMySQLRepository{
		db: db,
	}
}

func (r *purchaseOrderMySQLRepository) Create(orderNumber string, orderDate string, trackingCode string, buyerId int, productRecordId int, orderStatusId int) (domain.Purchase_Order, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return domain.Purchase_Order{}, err
	}

	const query = `INSERT INTO purchase_order (order_number, order_date, tracking_code, buyer_id, product_record_id, order_status_id) VALUES (?, ?, ?, ?, ?, ?, ?)`

	res, err := tx.Exec(query, orderNumber, orderDate, trackingCode, buyerId, productRecordId, orderStatusId)

	if err != nil {
		_ = tx.Rollback()
		return domain.Purchase_Order{}, err
	}

	if err = tx.Commit(); err != nil {
		return domain.Purchase_Order{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Purchase_Order{}, err
	}

	return domain.Purchase_Order{
		ID:              int(id),
		OrderNumber:     orderNumber,
		OrderDate:       orderDate,
		TrackingCode:    trackingCode,
		BuyerId:         buyerId,
		ProductRecordId: productRecordId,
		OrderStatusId:   orderStatusId,
	}, nil
}
