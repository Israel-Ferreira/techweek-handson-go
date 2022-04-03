package services

import (
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/exceptions"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/models"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/repositories"
)

type StockService interface {
	AddStockUpdate(string, data.StockQuantity) (models.Stock, error)
	SubstractStockUpdate(string, data.StockQuantity) (models.Stock, error)
	GetStock(sku string) (models.Stock, error)
	GetStocks() ([]models.Stock, error)
}

type stockService struct {
	repo repositories.StockRepository
}

func (s stockService) AddStockUpdate(sku string, stockAmountBody data.StockQuantity) (models.Stock, error) {

	if err := s.checkQuantityIsLessOrEqualZero(stockAmountBody); err != nil {
		return models.NewStock(), err
	}

	stock, err := s.repo.AddStockQty(sku, stockAmountBody)

	if err != nil {
		return models.NewStock(), err
	}

	return stock, nil
}

func (s stockService) SubstractStockUpdate(sku string, stockAmountBody data.StockQuantity) (models.Stock, error) {

	if err := s.checkQuantityIsLessOrEqualZero(stockAmountBody); err != nil {
		return models.NewStock(), err
	}

	stock, err := s.repo.SubstractStockQty(sku, stockAmountBody)

	if err != nil {
		return models.NewStock(), err
	}

	return stock, nil
}

func (s stockService) GetStock(sku string) (models.Stock, error) {
	stockItem, err := s.repo.GetStockItem(sku)

	if err != nil {
		return models.NewStock(), err
	}

	return stockItem, nil
}

func (s stockService) GetStocks() ([]models.Stock, error) {
	stocks, err := s.repo.ListItems()

	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (s stockService) checkQuantityIsLessOrEqualZero(data data.StockQuantity) error {
	if data.Qty <= 0 {
		return exceptions.ErrorInvalidQty
	}

	return nil
}

func NewStockService(repo repositories.StockRepository) *stockService {
	return &stockService{
		repo: repo,
	}
}
