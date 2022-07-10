package domain

type Purchase_Order struct {
	ID int `json:"id"`
	OrderNumber string `json:"order_number"`
	OrderDate string `json:"order_date"`
	TrackingCode string `json:"tracking_code"`
	BuyerId string `json:"buyer_id"`
	ProductRecordId string `json:"product_record_id"`
	OrderStatusId string `json:"order_status_id"`
}

type Purchase_Orders []Purchase_Order