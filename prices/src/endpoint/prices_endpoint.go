package endpoint

import (
	"context"

	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/services"
	"github.com/go-kit/kit/endpoint"
)

func GetAllPrices(svc services.PriceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		prices, err := svc.GetAll()

		if err != nil {
			return nil, err
		}

		return prices, nil
	}
}

func GetBySKU(svc services.PriceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		sku := request.(string)

		price, err := svc.GetBySku(sku)

		if err != nil {
			return nil, err
		}

		return price, nil
	}
}

func UpdatePrice(svc services.PriceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(data.PriceDecodeDTO)

		price, err := svc.Update(req.Sku, req.UpdatePrice)

		if err != nil {
			return nil, err
		}

		return price, nil
	}
}
