package models

type (
	AppleResponse struct {
		Etas []Eta `json:"etas"`
	}

	Eta struct {
		Destination               Address `json:"destination"`
		TransportType             string  `json:"transportType"`
		DistanceMeters            int     `json:"distanceMeters"`
		ExpectedTravelTimeSeconds int     `json:"expectedTravelTimeSeconds"`
		StaticTravelTimeSeconds   int     `json:"staticTravelTimeSeconds"`
	}
)
