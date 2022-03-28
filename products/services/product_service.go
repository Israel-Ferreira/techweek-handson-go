package services

import (
	"context"

	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/models"
)

type ProductService interface {
	CreateProduct(context.Context, data.CreateProduct) (string, error)
	GetProducts(ctx context.Context) []models.Product
	GetProductBySku(context.Context, string) models.Product
	DeleteBySku(context.Context, string) error
	UpdateProduct(context.Context, string, data.UpdateProduct) error
}

type productService struct{}

func (s *productService) GetProducts(ctx context.Context) []models.Product {
	return nil
}

func (s *productService) GetProductBySku(ctx context.Context, sku string) models.Product {
	return models.NewProduct()
}

func (s *productService) DeleteBySku(ctx context.Context, sku string) error {
	return nil
}

func (s *productService) UpdateProduct(ctx context.Context, sku string, dto data.UpdateProduct) error {
	return nil
}

func (s *productService) CreateProduct(ctx context.Context, dto data.CreateProduct) (string, error) {
	return "", nil
}

func NewProductService() *productService {
	return &productService{}
}
