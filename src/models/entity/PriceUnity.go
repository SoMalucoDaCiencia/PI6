package entity

import (
	"time"

	yaml "gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

type PriceUnity struct {
	ID         *uint64   `json:"id"         gorm:"column:id;primary_key;auto_increment;notnull"`
	Price      float64   `json:"price"      gorm:"column:price;notnull"`
	CreatedAt  time.Time `json:"createdAt"  gorm:"column:createdAt;notnull"`
	ChemicalID uint64    `json:"chemicalID" gorm:"column:chemicalID;"`

	// Many-to-One
	Chemical *Chemical `json:"chemical" gorm:"foreignKey:ChemicalID;references:ID"`
}

func (this *PriceUnity) TableName() string {
	return "DB_PRICE_UNITY"
}

func (this *PriceUnity) BeforeCreate(db *gorm.DB) error {
	return nil
}

func (this *PriceUnity) BeforeUpdate(db *gorm.DB) error {
	return nil
}

func (this *PriceUnity) BeforeDelete(db *gorm.DB) error {
	return nil
}

func (this *PriceUnity) AsString() string {
	data, err := yaml.Marshal(this)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
