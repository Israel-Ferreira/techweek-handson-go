package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Sku         string `gorm:"unique" json:"sku"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Active      bool   `gorm:"default:true" json:"active"`
}

func (p *Product) IsValid() error {
	return nil
}

func NewProduct() Product {
	return Product{}
}
