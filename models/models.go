package models

type CostResponse struct {
	Rajaongkir struct {
		Results Province `json:"results"`
	} `json:"rajaongkir"`
}

type eResults struct {
	ID          int    `gorm:"primaryKey"`
	City_id     string `json:"city_id"`
	Province_id string `json:"province_id"`
	Province    string `json:"province"`
	Tipe        string `json:"type"`
	City_name   string `json:"city_name"`
	Postal_code string `json:"postal_code"`
}

type City struct {
	City_id     string `json:"city_id"`
	Province_id string `json:"province_id"`
	Province    string `json:"province"`
	Tipe        string `json:"type"`
	City_name   string `json:"city_name"`
	Postal_code string `json:"postal_code"`
}

type Province struct {
	ID          int    `gorm:"primaryKey"`
	Province_id string `json:"province_id"`
	Province    string `json:"province"`
}
