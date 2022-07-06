package domain

type Locality struct{
	Id	int `json:"Id"`
	LocalityName int `json:"Locality:Name"`
	ProvinceName string `json:"ProvinceName"`
	CountryName string `json:"CountryName"`
	
}

type Localities []Locality