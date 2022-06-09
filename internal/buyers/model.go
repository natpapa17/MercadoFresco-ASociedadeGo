package buyers

type Buyer struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string `json:"address"`
	DocumentNumber string `json:"document_number"`
}