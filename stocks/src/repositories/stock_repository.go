package repositories

import (
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/models"
	"gorm.io/gorm"
)

type StockRepository interface {
	CreateStockItem(models.Stock) error
	DeleteStockItemBySku(string) error
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

func NewStockRepository(db *gorm.DB) *stockRepository {
	return &stockRepository{db: db}
}
