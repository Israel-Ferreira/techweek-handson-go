package endpoints

import (
	"context"

	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/models"
	"github.com/Israel-Ferreira/techweek-hands-on/products/services"
	"github.com/go-kit/kit/endpoint"
)

func GetProductsEndpoint(svc services.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := svc.GetProducts(ctx)

		if err != nil {
			return nil, err
		}

		return response, nil
	}
}

func GetProductEndpoint(svc services.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		sku := request.(string)

		response, err := svc.GetProductBySku(ctx, sku)

		if err != nil {
			return models.Product{}, err
		}

		return response, nil
	}
}

func CreateProduct(svc services.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(data.CreateProduct)

		productId, err := svc.CreateProduct(ctx, req)

		if err != nil {
			return "", err
		}

		return productId, nil

	}
}
