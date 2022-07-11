package domain

type Carrier struct {
	Id          int    `json:"id"`
	Cid         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityId  int    `json:"locality_id"`
}

type NumberOfCarriersPerLocality struct {
	LocalityId    int    `json:"locality_id"`
	LocalityName  string `json:"locality_name"`
	CarriersCount int    `json:"carriers_count"`
}
