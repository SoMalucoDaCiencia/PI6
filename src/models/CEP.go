package models

import "PI6/share"

type Address struct {
	District string  `json:"District"`
	Lat      float64 `json:"Lat"`
	Long     float64 `json:"Long"`
}

func (this *Address) GetUuid() string {
	return share.FloatsAsUUID(this.Lat, this.Long)
}
