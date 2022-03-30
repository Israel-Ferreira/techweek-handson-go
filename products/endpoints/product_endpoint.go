package endpoints

import (
	"context"
	"fmt"

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

		fmt.Println(sku)

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

func DeleteProduct(svc services.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		sku := request.(string)

		if err := svc.DeleteBySku(ctx, sku); err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func UpdateProduct(svc services.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(data.UpdateProductReq)

		if err := svc.UpdateProduct(ctx, req.Sku, req.UpdateProduct); err != nil {
			return nil, err
		}

		return nil, nil
	}
}
