package repositories

import "github.com/Israel-Ferreira/techweek-hands-on/products/models"

type ProductRepository interface {
	FindBySku(string) (models.Product, error)
	FindAll() ([]models.Product, error)
	Update(string, models.Product) error
	Create(models.Product) (string, error)
	Delete(string) error
}

type repository struct{}

func (r *repository) FindBySku(sku string) (models.Product, error) {
	return models.NewProduct(), nil
}

func (r *repository) FindAll() ([]models.Product, error) {
	return nil, nil
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

func NewRepository() ProductRepository {
	return &repository{}
}
