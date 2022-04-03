package repositories

import (
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/models"
	"gorm.io/gorm"
)

type StockRepository interface {
	CreateStockItem(models.Stock) error
	DeleteStockItemBySku(string) error
	ListItems() ([]models.Stock, error)
	GetStockItem(sku string) (models.Stock, error)
	AddStockQty(sku string, data data.StockQuantity) (models.Stock, error)
	SubstractStockQty(sku string, data data.StockQuantity) (models.Stock, error)
}

type stockRepository struct {
	db *gorm.DB
}

func (sr stockRepository) CreateStockItem(stockItem models.Stock) error {

	if err := stockItem.StockIsValid(); err != nil {
		return err
	}

	txn := sr.db.Create(&stockItem)

	if txn.Error != nil {
		return txn.Error
	}

	return nil
}

func (sr stockRepository) DeleteStockItemBySku(sku string) error {

	txn := sr.db.Delete(&models.Stock{Sku: sku}, "sku = ?", sku)

	if txn.Error != nil {
		return txn.Error
	}

	return nil
}

func (sr stockRepository) ListItems() ([]models.Stock, error) {
	var stockItens []models.Stock

	txn := sr.db.Find(&stockItens)

	if txn.Error != nil {
		return nil, txn.Error
	}

	return stockItens, nil
}

func (sr stockRepository) GetStockItem(sku string) (models.Stock, error) {

	var stockItem models.Stock

	txn := sr.db.First(&stockItem, "sku = ?", sku)

	if txn.Error != nil {
		return models.NewStock(), txn.Error
	}

	return stockItem, nil
}

func (sr stockRepository) AddStockQty(sku string, data data.StockQuantity) (models.Stock, error) {
	stock, err := sr.GetStockItem(sku)

	if err != nil {
		return models.NewStock(), err
	}

	currentQty := stock.CurrentQty + data.Qty

	txn := sr.db.Model(&stock).Update("current_qty", currentQty)

	if txn.Error != nil {
		return models.NewStock(), txn.Error
	}

	return stock, nil
}

func (sr stockRepository) SubstractStockQty(sku string, data data.StockQuantity) (models.Stock, error) {
	stock, err := sr.GetStockItem(sku)

	if err != nil {
		return models.NewStock(), err
	}

	currentQty := stock.CurrentQty - data.Qty

	txn := sr.db.Model(&stock).Update("current_qty", currentQty)

	if txn.Error != nil {
		return models.NewStock(), txn.Error
	}

	return stock, nil
}

func NewStockRepository(db *gorm.DB) *stockRepository {
	return &stockRepository{db: db}
}
