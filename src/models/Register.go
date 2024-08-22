package models

import "time"

type Register struct {
	Origin      *string      `json:"origin"      gorm:"type:varchar(9);primaryKey;autoIncrement:false;notnull"`
	Destiny     *string      `json:"destiny"     gorm:"type:varchar(9);primaryKey;autoIncrement:false;notnull"`
	CreatedAt   *time.Time   `json:"createdAt"   gorm:"column:createdAt;notnull"`
	Distance    *uint32      `json:"distance"    gorm:"column:distance;notnull"`
	Duration    *uint32      `json:"duration"    gorm:"column:duration;notnull"`
	Temperature *uint8       `json:"temperature" gorm:"column:temperature;notnull"`
	Weather     *WeatherEnum `json:"weather"     gorm:"column:weather;type:varchar(10);notnull"`
}
