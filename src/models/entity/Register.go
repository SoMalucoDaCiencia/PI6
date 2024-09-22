package entity

import (
	"PI6/models"
	"PI6/share"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	yaml "gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

type Register struct {
	ID             uint64             `json:"id"             gorm:"column:id;primary_key;auto_increment;notnull"`
	Origin         string             `json:"origin"         gorm:"type:varchar(36);primaryKey;autoIncrement:false;notnull"`
	Destiny        string             `json:"destiny"        gorm:"type:varchar(36);primaryKey;autoIncrement:false;notnull"`
	CreatedAt      time.Time          `json:"createdAt"      gorm:"column:createdAt;notnull"`
	Distance       int                `json:"distance"       gorm:"column:distance;notnull"`
	TimeAutomobile int                `json:"timeAutomobile" gorm:"column:timeAutomobile;notnull"`
	TimeTransit    int                `json:"timeTransit"    gorm:"column:timeTransit;notnull"`
	TimeWalking    int                `json:"timeWalking"    gorm:"column:timeWalking;notnull"`
	Temperature    int                `json:"temperature"    gorm:"column:temperature;notnull"`
	Weather        models.WeatherEnum `json:"weather"        gorm:"column:weather;type:varchar(10);notnull"`
}

func (this *Register) BeforeCreate(db *gorm.DB) error {
	if this.Origin == this.Destiny {
		return errors.New("origin and Destiny cannot be the same")
	}
	return nil
}

func (this *Register) BeforeDelete(db *gorm.DB) error {
	return nil
}

func (this *Register) BeforeUpdate(db *gorm.DB) error {
	return nil
}

func (this *Register) TableName() string {
	return "DB_REGISTER"
}

func (this *Register) AsString() string {
	data, err := yaml.Marshal(this)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func ExtractRegister(token string, origin *models.Address, destiny []*models.Address) (ret []Register, err error) {

	var response []byte
	headers := map[string]string{
		"Host":          "go-app",
		"Authorization": "Bearer " + token,
	}

	// Ajusta a URL com todas as localizações
	// ######################################################
	url := fmt.Sprintf(share.AppleURI, fmt.Sprintf("%.6f,%.6f", origin.Lat, origin.Long))
	for _, gp := range destiny {
		url += fmt.Sprintf("%g,%g|", gp.Lat, gp.Long)
	}

	// Cria 3 requisições para 3 transportes diferentes
	// ######################################################
	var arr [3]models.AppleResponse
	for i, str := range []models.TransportEnum{models.Automobile, models.Transit, models.Walking} {
		urlFinal := url[:len(url)-2] + fmt.Sprintf("&transportType=%s", str)
		_, err = share.Rest("GET", urlFinal, &response, headers, nil, nil)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(response, &arr[i])
		if err != nil {
			return nil, err
		}
	}

	originUUID := share.FloatsAsUUID(origin.Lat, origin.Long)
	for i, eta := range arr[0].Etas {
		ret = append(ret, Register{
			Origin:         originUUID,
			Destiny:        share.FloatsAsUUID(eta.Destination.Lat, eta.Destination.Long),
			CreatedAt:      time.Now(),
			Distance:       eta.DistanceMeters,
			TimeAutomobile: eta.ExpectedTravelTimeSeconds,
			TimeTransit:    arr[1].Etas[i].ExpectedTravelTimeSeconds,
			TimeWalking:    arr[2].Etas[i].ExpectedTravelTimeSeconds,
		})
	}
	return ret, nil
}
