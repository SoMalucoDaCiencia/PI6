package models

type GeoPos struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func NewGeoPos(lat, long float64) GeoPos {
	return GeoPos{
		Lat:  lat,
		Long: long,
	}
}
