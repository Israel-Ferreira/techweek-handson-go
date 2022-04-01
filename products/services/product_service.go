package services

import (
	"context"
	"fmt"

	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/exceptions"
	"github.com/Israel-Ferreira/techweek-hands-on/products/models"
	"github.com/Israel-Ferreira/techweek-hands-on/products/producers"
	"github.com/Israel-Ferreira/techweek-hands-on/products/repositories"
	"github.com/segmentio/kafka-go"
)

type ProductService interface {
	CreateProduct(context.Context, data.CreateProduct) (string, error)
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProductBySku(context.Context, string) (models.Product, error)
	DeleteBySku(context.Context, string) error
	UpdateProduct(context.Context, string, data.UpdateProduct) error
}

type productService struct {
	repo  repositories.ProductRepository
	kafka producers.ProductProducer
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

	if err := s.kafka.SendDeleteEventMsg(sku); err != nil {
		return err
	}

	return nil
}

func (s *productService) UpdateProduct(ctx context.Context, sku string, dto data.UpdateProduct) error {

	if sku == "" {
		return exceptions.ErrorInvalidParam
	}

	if dto.Brand == "" || dto.Title == "" {
		return exceptions.ErrorBodyIsNotValid
	}

	if err := s.repo.Update(sku, dto); err != nil {
		return err
	}

	if err := s.kafka.SendUpdateProductMsg(models.Product{Sku: sku, Title: dto.Title}); err != nil {
		return err
	}

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

	if err := s.kafka.SendNewProductEventMsg(product); err != nil {
		return "", err
	}

	return productSku, nil
}

func NewProductService(repo repositories.ProductRepository, producer *kafka.Writer) *productService {
	
	fmt.Println(producer != nil)

	return &productService{
		repo:  repo,
		kafka: producers.ProductProducer{Kafka: producer},
	}
}
