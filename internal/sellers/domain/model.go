package domain

type Seller struct{
	Id	int `json:"id"`
	Cid int `json:"cid"`
	CompanyName string `json:"company_name"`
	Address string `json:"address"`
	Telephone string `json:"telephone"`
	LocalityId  int    `json:"locality_id"`
}

type Sellers []Seller