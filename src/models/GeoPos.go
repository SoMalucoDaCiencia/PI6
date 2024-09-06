package models

type GeoPos struct {
	lat, long float32
}

func NewGeoPos(lat, long float32) GeoPos {
	return GeoPos{
		lat:  lat,
		long: long,
	}
}
