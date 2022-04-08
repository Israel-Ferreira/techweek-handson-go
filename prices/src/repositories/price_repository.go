package repositories

import (
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/models"
	"gorm.io/gorm"
)

type PriceRepository interface {
	AddItem(models.Price) (string, error)
	GetBySku(string) (models.Price, error)
	GetAll() ([]models.Price, error)
	DeleteBySku(string) error
	Update(string, data.UpdatePrice) (models.Price, error)
}

type priceRepo struct {
	Db *gorm.DB
}

func (pr priceRepo) AddItem(price models.Price) (string, error) {
	txn := pr.Db.Save(&price)

	if err := txn.Error; err != nil {
		return "", txn.Error
	}

	return price.Sku, nil
}

func (pr priceRepo) GetBySku(sku string) (models.Price, error) {
	var price models.Price

	txn := pr.Db.Find(&price, "sku = ?", sku)

	if txn.Error != nil {
		return models.Price{}, txn.Error
	}

	return price, nil
}

func (pr priceRepo) DeleteBySku(sku string) error {
	itemPrice, err := pr.GetBySku(sku)

	if err != nil {
		return err
	}

	txn := pr.Db.Delete(&itemPrice)

	if txn.Error != nil {
		return txn.Error
	}

	return nil
}

func (pr priceRepo) Update(sku string, data data.UpdatePrice) (models.Price, error) {
	price, err := pr.GetBySku(sku)

	if err != nil {
		return models.Price{}, err
	}

	txn := pr.Db.Model(&price).UpdateColumns(&models.Price{
		OfferPrice:         data.OfferPrice,
		SalePrice:          data.SalePrice,
		DiscountPercentage: data.Discount,
	})

	if txn.Error != nil {
		return models.Price{}, txn.Error
	}

	return price, nil
}

func (pr priceRepo) GetAll() ([]models.Price, error) {
	var prices []models.Price

	txn := pr.Db.Find(&prices)

	if txn.Error != nil {
		return nil, txn.Error
	}

	return prices, nil
}

func NewPriceRepo(db *gorm.DB) *priceRepo {
	return &priceRepo{Db: db}
}
