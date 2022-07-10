package records

type Records struct {
	Id               int    `json:"id"`
	Last_Update_Date string `json:"last_update_date"`
	Purchase_Price   int    `json:"purchase_price"`
	Sale_Price       int    `json:"sale_price"`
	Product_Id       int    `json:"product_id"`
}
