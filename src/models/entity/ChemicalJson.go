package entity

import "time"

type ChemicalJson struct {
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
