package models

import (
	share "PI6/share"
	"encoding/json"
	"fmt"
	"time"

	yaml "gopkg.in/yaml.v2"

	mine "github.com/artking28/myGoUtils"
)

type Register struct {
	Id          *uint64        `json:"id"            gorm:"column:id;primary_key;auto_increment;notnull"`
	Origin      *string        `json:"origin"        gorm:"type:varchar(9);primaryKey;autoIncrement:false;notnull"`
	Destiny     *string        `json:"destiny"       gorm:"type:varchar(9);primaryKey;autoIncrement:false;notnull"`
	CreatedAt   *time.Time     `json:"createdAt"     gorm:"column:createdAt;notnull"`
	Transport   *TransportEnum `json:"transport"   gorm:"column:weather;type:varchar(10);notnull"`
	Distance    *uint32        `json:"distance"      gorm:"column:distance;notnull"`
	Duration    *uint32        `json:"duration"      gorm:"column:duration;notnull"`
	Temperature *uint8         `json:"temperature"   gorm:"column:temperature;notnull"`
	Weather     *WeatherEnum   `json:"weather"       gorm:"column:weather;type:varchar(10);notnull"`
}

func (r *Register) AsString() string {
	data, err := yaml.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func ExtractRegister(token string, origin GeoPos, destiny []GeoPos) (ret []Register, err error) {

	response := []byte{}
	headers := map[string]string{
		"Host":          "go-app",
		"Authorization": "Bearer " + share.AppleToken,
	}

	url := fmt.Sprintf(token, fmt.Sprintf("%.6f,%.6f", origin.Lat, origin.Long))
	for _, gp := range destiny {
		url += fmt.Sprintf("%g,%g|", gp.Lat, gp.Long)
	}

	_, err = share.Rest("GET", url[:len(url)-2], &response, headers, nil, nil)
	if err != nil {
		return nil, err
	}

	f := AppleResponse{}
	err = json.Unmarshal(response, &f)
	if err != nil {
		return nil, err
	}

	originUUID := share.FloatsAsUUID(origin.Lat, origin.Long)
	for _, eta := range f.Etas {
		ret = append(ret, Register{
			Origin:    &originUUID,
			Destiny:   mine.Ptr(share.FloatsAsUUID(eta.Destination.Lat, eta.Destination.Long)),
			Transport: (*TransportEnum)(&eta.TransportType),
			CreatedAt: mine.Ptr(time.Now()),
			Distance:  mine.Ptr[uint32](uint32(eta.DistanceMeters)),
			Duration:  mine.Ptr[uint32](uint32(eta.ExpectedTravelTimeSeconds)),
		})
	}
	return ret, nil
}
