package sellers

type Seller struct{
	Id	int `json:"Id"`
	Cid int `json:"Cid"`
	CompanyName string `json:"CompanyName"`
	Address string `json:"Address"`
	Telephone string `json:"Telephone"`
}