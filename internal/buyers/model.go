package buyers

type Buyer struct {
	ID int `json:"id"`
	First_name string `json:"firstName"`
	Last_name string `json:"lastName"`
	Address string `json:"address"`
	Document_number string `json:"document"`
}