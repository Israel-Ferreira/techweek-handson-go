package services

import (
	"context"

	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/exceptions"
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
	repo repositories.ProductRepository
}

func (s *productService) GetProducts(ctx context.Context) ([]models.Product, error) {
	products, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) GetProductBySku(ctx context.Context, sku string) (models.Product, error) {
	product, err := s.repo.FindBySku(sku)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *productService) DeleteBySku(ctx context.Context, sku string) error {

	if sku == "" {
		return exceptions.ErrorInvalidParam
	}

	if err := s.repo.Delete(sku); err != nil {
		return err
	}

	return nil
}

func (s *productService) UpdateProduct(ctx context.Context, sku string, dto data.UpdateProduct) error {
	return nil
}

func (s *productService) CreateProduct(ctx context.Context, dto data.CreateProduct) (string, error) {
	product := models.Product{
		Sku:         dto.Sku,
		Title:       dto.Title,
		Description: dto.Description,
		Brand:       dto.Brand,
	}

	if err := product.IsValid(); err != nil {
		return "", err
	}

	productSku, err := s.repo.Create(product)

	if err != nil {
		return "", err
	}

	return productSku, nil
}

func NewProductService(repo repositories.ProductRepository) *productService {
	return &productService{
		repo: repo,
	}
}
