package models

import (
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/exceptions"
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	Title      string `json:"title"`
	Sku        string `gorm:"unique" json:"sku"`
	CurrentQty int    `json:"currentQty"`
}

func (st *Stock) StockIsValid() error {
	if st.Title == "" {
		return exceptions.ErrorInvalidTitle
	}

	if st.Sku == "" {
		return exceptions.ErrorInvalidSKU
	}

	return nil
}

func NewStock() Stock {
	return Stock{}
}

