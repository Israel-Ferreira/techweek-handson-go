package data

import "github.com/Israel-Ferreira/techweek-hands-on/prices/src/exceptions"

type UpdatePrice struct {
	CurrentPrice float64 `json:"currentPrice"`
	SalePrice    float64 `json:"wholeSalePrice"`
	OfferPrice   float64 `json:"offerPrice"`
	Discount     float64 `json:"discount"`
}

func (u *UpdatePrice) Validate() error {
	if u.OfferPrice <= 0.00 {
		return exceptions.ErrorInvalidOfferPrice
	}

	if u.CurrentPrice <= 0.00 {
		return exceptions.ErrorInvalidCurrentPrice
	}

	if u.SalePrice <= 0.00 {
		return exceptions.ErrorInvalidSalePrice
	}

	return nil
}
