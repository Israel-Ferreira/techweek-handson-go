package repositories

import (
	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/exceptions"
	"github.com/Israel-Ferreira/techweek-hands-on/products/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindBySku(string) (models.Product, error)
	FindAll() ([]models.Product, error)
	Update(string, data.UpdateProduct) error
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
		return product, exceptions.ErrorNotFoundProduct
	}

	return product, nil
}

func (r *repository) FindAll() ([]models.Product, error) {
	var products []models.Product

	txn := r.db.Find(&products, "active = ?", true)

	if txn.Error != nil {
		return nil, txn.Error
	}

	return products, nil
}

func (r *repository) Update(sku string, productDto data.UpdateProduct) error {
	product, err := r.FindBySku(sku)

	if err != nil {
		return err
	}

	txn := r.db.Model(&product).Updates(models.Product{Title: productDto.Title, Description: productDto.Description, Brand: productDto.Brand})

	if txn.Error != nil {
		return txn.Error
	}

	return nil
}

func (r *repository) Delete(sku string) error {
	var product models.Product

	txnFindById := r.db.Find(&product, "sku = ?", sku)

	if txnFindById.Error != nil {
		return txnFindById.Error
	}

	txnUpdate := r.db.Model(&product).Update("active", false)

	if txnUpdate.Error != nil {
		return txnUpdate.Error
	}

	return nil
}

func (r *repository) Create(product models.Product) (string, error) {

	txn := r.db.Create(&product)

	if txn.Error != nil {
		return "", txn.Error
	}

	return product.Sku, nil
}

func NewRepository(db *gorm.DB) ProductRepository {
	return &repository{
		db: db,
	}
}
