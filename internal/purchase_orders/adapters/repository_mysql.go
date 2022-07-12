package adapters

import (
	"database/sql"
	"errors"

	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/domain"
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/usecases"

)

type purchaseOrderMySQLRepository struct{
	db *sql.DB
}

func CreatePurchaseOrderMySQLRepository(db *sql.DB) usecases.PurchaseOrderRepository {
	return &purchaseOrderRepository{
		db: db,
	}
}

func (r *purchaseOrderMySQLRepository) Create(orderNumber string, orderDate string, trackingCode string, buyerId int, productRecordId int, orderStatusId int) (domain.PurchaseOrder, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return domain.PurchaseOrder{}, err
	}

	const query = `INSERT INTO purchase_order (order_number, order_date, tracking_code, buyer_id, product_record_id, order_status_id) VALUES (?, ?, ?, ?, ?, ?, ?)`

	res, err := tx.Exec(query, orderNumber, orderDate, trackingCode, buyerId, productRecordId, orderStatusId)


	if err != nil {
		_ = tx.Rollback()
		return domain.PurchaseOrder{}, err
	}

	if err = tx.Commit(); err != nil {
		return domain.PurchaseOrder{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.PurchaseOrder{}, err
	}

	return domain.PurchaseOrder{
		ID: int(id),
		OrderNumber: orderNumber,
		OrderDate: orderDate,
		TrackingCode: trackingCode,
		BuyerId: buyerId,
		ProductRecordId: productRecordId,
		OrderStatusId: orderStatusId,
	}, nil
}

