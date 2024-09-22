package models

import "PI6/share"

type Address struct {
	District string  `json:"District"`
	Lat      float64 `json:"latitude"`
	Long     float64 `json:"longitude"`
}

func (this *Address) GetUuid() string {
	return share.FloatsAsUUID(this.Lat, this.Long)
}
