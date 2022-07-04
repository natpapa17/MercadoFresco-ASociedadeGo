package products

type Product struct {
	Id                               int     `json:"id"`
	Product_Code                     string  `json:"product_code"`
	Description                      string  `json:"description"`
	Width                            float64 `json:"width"`
	Height                           float64 `json:"height"`
	Length                           float64 `json:"length"`
	Net_Weight                       float64 `json:"net_weight"`
	Expiration_Rate                  int     `json:"expiration_rate"`
	Recommended_Freezing_Temperature float64 `json:"recommended_freezing_temperature"`
	Freezing_Rate                    int     `json:"freezing_rate"`
	Product_Type_Id                  int     `json:"product_type_id"`
	Seller_Id                        int     `json:"seller_id"`
}

type ProductRecords struct {
	Id               int    `json:"id"`
	Last_Update_Data string `json:"last_update_data"`
	Purchase_Price   int    `json:"purchase_price"`
	Sale_Price       int    `json:"sale_price"`
	Product_Id       int    `json:"product_id"`
}
