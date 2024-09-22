package models

type Address struct {
	District string  `json:"District"`
	Lat      float64 `json:"Lat"`
	Long     float64 `json:"Long"`
}
