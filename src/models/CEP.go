package models

import "fmt"

type Address struct {
	Cep      string `json:"cep"`
	District string `json:"destiny"`
	Street   string `json:"createdAt"`
}

func (this Address) ToString() string {
	return fmt.Sprintf("%s,%s,%s", this.Street, this.Street, this.Cep)
}
