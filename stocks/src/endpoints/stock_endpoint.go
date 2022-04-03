package endpoints

import (
	"context"

	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/models"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/services"
	"github.com/go-kit/kit/endpoint"
)

func GetAllStockItems(svc services.StockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		stocks, err := svc.GetStocks()

		if err != nil {
			return nil, err
		}

		return stocks, nil
	}
}

func GetBySku(svc services.StockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		sku := request.(string)

		stock, err := svc.GetStock(sku)

		if err != nil {
			return models.NewStock(), err
		}

		return stock, nil
	}
}

func AddStockEndpoint(svc services.StockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		decodedBody := request.(data.StockDecodeDTO)

		stock, err := svc.AddStockUpdate(decodedBody.Sku, decodedBody.StockQuantity)

		if err != nil {
			return nil, err
		}

		return stock, nil
	}
}

func SubstractStockEndpoint(svc services.StockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		decodedBody := request.(data.StockDecodeDTO)

		stock, err := svc.SubstractStockUpdate(decodedBody.Sku, decodedBody.StockQuantity)

		if err != nil {
			return nil, err
		}

		return stock, nil
	}
}
