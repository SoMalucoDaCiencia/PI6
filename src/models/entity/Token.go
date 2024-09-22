package entity

import (
	"PI6/models"
	"PI6/share"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID        uint64    `json:"id"        gorm:"column:id;primary_key;auto_increment;notnull"`
	Value     string    `json:"value"     gorm:"column:value;notnull"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;type:datetime;notnull"`
}

func (this *Token) BeforeCreate(db *gorm.DB) error {
	return nil
}

func (this *Token) BeforeDelete(db *gorm.DB) error {
	return nil
}

func (this *Token) BeforeUpdate(db *gorm.DB) error {
	return nil
}

func (this *Token) TableName() string {
	return "DB_TOKEN"
}

func (this *Token) ToString() string {
	return this.Value
}

func ExtractToken(seedToken string) (ret models.AppleTokenGetter, err error) {

	response := []byte{}
	headers := map[string]string{
		"Host":          "go-app",
		"Authorization": "Bearer " + seedToken,
	}

	_, err = share.Rest("GET", "https://maps-api.apple.com/v1/token", &response, headers, nil, nil)
	if err != nil {
		return models.AppleTokenGetter{}, err
	}

	f := models.AppleTokenGetter{
		Seed:      seedToken,
		CreatedAt: time.Now(),
	}
	return f, json.Unmarshal(response, &f)
}
