package entity

import (
	"PI6/share"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	yaml "gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type (
	Chemical struct {
		ID              *uint64 `json:"id,omitempty"              gorm:"column:id;primary_key;auto_increment;notnull"`
		ExternalId      string  `json:"productId,omitempty"       gorm:"column:externalId;notnull;unique"`
		Ean             string  `json:"ean,omitempty"             gorm:"column:ean;notnull;unique"`
		ProductTittle   string  `json:"productTitle,omitempty"    gorm:"column:productTittle;notnull"`
		Brand           string  `json:"brand,omitempty"           gorm:"column:brand;notnull"`
		Link            string  `json:"link,omitempty"            gorm:"column:link;notnull"`
		MeasurementUnit string  `json:"measurementUnit,omitempty" gorm:"column:measurementUnit;notnull"`
		UnitMultiplier  float64 `json:"unitMultiplier,omitempty"  gorm:"column:unitMultiplier;notnull"`

		// One-to-Many
		Prices []PriceUnity `json:"prices,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChemicalID"`
	}

	ChemicalJson struct {
		Chemical
		BrandId uint64 `json:"brandId"`
		Items   []struct {
			Ean             string  `json:"ean"`
			MeasurementUnit string  `json:"measurementUnit"`
			UnitMultiplier  float64 `json:"unitMultiplier"`
			Sellers         []struct {
				CommertialOffer struct {
					Price float64 `json:"Price"`
				} `json:"commertialOffer"`
			} `json:"sellers"`
		} `json:"items"`
	}
)

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
	return fmt.Sprintf("Chemical {\n%s}", string(data))
}

func NewChemical(href string, client *http.Client) (*Chemical, error) {
	skuColly := colly.NewCollector()
	var sku string

	// Find and visit all links
	skuColly.OnHTML("div.skuReference", func(e *colly.HTMLElement) {
		sku = e.Text
		//fmt.Println("Processing", sku)
	})

	err := skuColly.Visit(href)
	if err != nil {
		return nil, err
	}

	getRef := fmt.Sprintf("https://www.drogariasaopaulo.com.br/api/catalog_system/pub/products/search?fq=skuId:%s", sku)

	var body []byte
	_, err = share.RestClient(client, http.MethodGet, getRef, &body, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []ChemicalJson
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	che := res[0].Adapt()
	return &che, nil
}

func (this ChemicalJson) Adapt() Chemical {
	return Chemical{
		ExternalId:      this.ExternalId,
		Ean:             this.Items[0].Ean,
		ProductTittle:   this.ProductTittle,
		Brand:           this.Brand,
		Link:            this.Link,
		MeasurementUnit: this.Items[0].MeasurementUnit,
		UnitMultiplier:  this.Items[0].UnitMultiplier,
		Prices: []PriceUnity{{
			Price:     this.Items[0].Sellers[0].CommertialOffer.Price,
			CreatedAt: time.Now(),
		}},
	}
}
