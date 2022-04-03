package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/exceptions"
	"github.com/gorilla/mux"
)

func decode(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeRequestBodyWithParam(ctx context.Context, req *http.Request) (interface{}, error) {
	sku := mux.Vars(req)["sku"]

	if sku == "" {
		return "", exceptions.ErrorInvalidParam
	}

	var stockQuantityDTO data.StockQuantity

	if err := json.NewDecoder(req.Body).Decode(&stockQuantityDTO); err != nil {
		return nil, err
	}

	return data.StockDecodeDTO{Sku: sku, StockQuantity: stockQuantityDTO}, nil
}

func decodeRequestWithParam(ctx context.Context, req *http.Request) (interface{}, error) {
	sku := mux.Vars(req)["sku"]

	if sku == "" {
		return "", exceptions.ErrorInvalidParam
	}

	return sku, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
