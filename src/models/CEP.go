package models

import "fmt"

type Address struct {
	District string  `json:"District"`
	Lat      float64 `json:"Lat"`
	Long     float64 `json:"Long"`
}

func (this Address) ToString() string {
	return fmt.Sprintf("%s,%s,%s", this.Street, this.Street, this.Cep)
}
