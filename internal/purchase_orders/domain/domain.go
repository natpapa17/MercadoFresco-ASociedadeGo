package domain

type Purchase_Order struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string `json:"address"`
	DocumentNumber string `json:"document_number"`
}

type Purchase_Orders []Purchase_Order