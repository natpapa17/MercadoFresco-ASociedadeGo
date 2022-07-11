package domain 

type Provincy struct {
	Id         int    `json:"id"`
	Name       string `json:"locality_name"`
	Country_id int    `json:"country_id"`
}