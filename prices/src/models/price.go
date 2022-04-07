package models

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	Title              string  `json:"title"`
	Sku                string  `gorm:"unique" json:"sku"`
	OfferPrice         float64 `json:"offerPrice"`
	SalePrice          float64 `json:"salePrice"`
	DiscountPercentage float64 `json:"discountPercentage"`
}
