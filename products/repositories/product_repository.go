package repositories

import (
	"github.com/Israel-Ferreira/techweek-hands-on/products/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindBySku(string) (models.Product, error)
	FindAll() ([]models.Product, error)
	Update(string, models.Product) error
	Create(models.Product) (string, error)
	Delete(string) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindBySku(sku string) (models.Product, error) {
	var product models.Product

	txn := r.db.First(&product, "sku = ?", sku)

	if txn.Error != nil {
		return product, txn.Error
	}

	return product, nil
}

func (r *repository) FindAll() ([]models.Product, error) {
	var products []models.Product

	txn := r.db.Find(&products)

	if txn.Error != nil {
		return nil, txn.Error
	}

	return products, nil
}

func (r *repository) Update(sku string, product models.Product) error {
	return nil
}

func (r *repository) Delete(sku string) error {
	return nil
}

func (r *repository) Create(product models.Product) (string, error) {
	return "", nil
}

func NewRepository(db *gorm.DB) ProductRepository {
	return &repository{
		db: db,
	}
}
