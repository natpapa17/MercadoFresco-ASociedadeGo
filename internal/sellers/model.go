package sellers

type Seller struct{
	Id	int `json:"id"`
	Cid int `json:"cid"`
	CompanyName string `json:"CompanyName"`
	Addres string `json:"Addres"`
	Telephone string `json:"telephone"`
}