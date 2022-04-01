package services

import (
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/exceptions"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/repositories"
)

type StockService interface {
	AddStockUpdate(data.StockQuantity) (interface{}, error)
	SubstractStockUpdate(data.StockQuantity) (interface{}, error)
	GetStock(sku string) (interface{}, error)
	GetStocks() ([]any, error)
}

type stockService struct {
	repo *repositories.StockRepository
}

func (s stockService) AddStockUpdate(stockAmountBody data.StockQuantity) (any, error) {

	if err := s.checkQuantityIsLessOrEqualZero(stockAmountBody); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s stockService) SubstractStockUpdate(stockAmountBody data.StockQuantity) (any, error) {
	return nil, nil
}

func (s stockService) GetStock(sku string) (any, error) {
	return nil, nil
}

func (s stockService) GetStocks() ([]any, error) {
	return nil, nil
}

func (s stockService) checkQuantityIsLessOrEqualZero(data data.StockQuantity) error {
	if data.Qty <= 0 {
		return exceptions.ErrorInvalidQty
	}

	return nil
}

func NewStockService() *stockService {
	return &stockService{}
}
