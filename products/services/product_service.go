package services

import (
	"context"

	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/models"
	"github.com/Israel-Ferreira/techweek-hands-on/products/repositories"
)

type ProductService interface {
	CreateProduct(context.Context, data.CreateProduct) (string, error)
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProductBySku(context.Context, string) (models.Product, error)
	DeleteBySku(context.Context, string) error
	UpdateProduct(context.Context, string, data.UpdateProduct) error
}

type productService struct {
	repo *repositories.ProductRepository
}

func (s *productService) GetProducts(ctx context.Context) ([]models.Product, error) {
	return nil, nil
}

func (s *productService) GetProductBySku(ctx context.Context, sku string) (models.Product, error) {
	return models.NewProduct(), nil
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

func NewProductService(repo *repositories.ProductRepository) *productService {
	return &productService{
		repo: repo,
	}
}
