package domain

type Seller struct{
	Id	int `json:"Id"`
	Cid int `json:"Cid"`
	CompanyName string `json:"company_name"`
	Address string `json:"Address"`
	Telephone string `json:"Telephone"`
	LocalityId  int    `json:"locality_id"`
}

type Sellers []Seller