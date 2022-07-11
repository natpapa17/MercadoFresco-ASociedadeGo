package domain

type Locality struct {
	Id         int    `json:"id"`
	Name       string `json:"locality_name"`
	ProvinceId int    `json:"province_name"`
}
