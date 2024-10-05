package entity

import (
	yaml "gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

type Chemical struct {
	IEntity

	ID           uint64 `json:"id"           gorm:"column:id;primary_key;auto_increment;notnull"`
	EAN          string `json:"ean"          gorm:"column:ean;notnull"`
	SKU          string `json:"sku"          gorm:"column:sku;notnull"`
	Dose         string `json:"dose"         gorm:"column:dose;notnull"`
	Brand        string `json:"brand"        gorm:"column:brand;notnull"`
	Manufacturer string `json:"manufacturer" gorm:"column:manufacturer;notnull"`

	// One-to-Many
	Prices []PriceUnity `json:"prices" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChemicalID"`
}

func (this *Chemical) TableName() string {
	return "DB_CHEMICAL"
}

func (this *Chemical) BeforeCreate(db *gorm.DB) error {
	return nil
}

func (this *Chemical) BeforeDelete(db *gorm.DB) error {
	return nil
}

func (this *Chemical) BeforeUpdate(db *gorm.DB) error {
	return nil
}

func (this *Chemical) AsString() string {
	data, err := yaml.Marshal(this)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
