package models

import (
	"github.com/Israel-Ferreira/techweek-hands-on/products/exceptions"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Sku         string `gorm:"unique" json:"sku"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Brand       string `json:"brand"`
	Active      bool   `gorm:"default:true" json:"active"`
}

func (p *Product) IsValid() error {
	if p.Title == "" {
		return exceptions.ErrorBodyIsNotValid
	}

	if p.Sku == "" {
		return exceptions.ErrorBodyIsNotValid
	}

	if p.Brand == "" {
		return exceptions.ErrorBodyIsNotValid
	}

	return nil
}

func NewProduct() Product {
	return Product{}
}
