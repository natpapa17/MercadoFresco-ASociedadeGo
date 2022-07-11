package domain

type Locality struct {
	Id	int `json:"Id"`
	Name string `json:"name"`
	Province_id int `json:"province_id"`
	
}

type FullLocality struct{
	Id	int `json:"Id"`
	Name string `json:"name"`
	ProvinceName string `json:"province_name"`
	CountryName string `json:"country_name"`

}


type Localities []Locality
