package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/exceptions"
	"github.com/gorilla/mux"
)

func decode(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeProductJsonBody(ctx context.Context, req *http.Request) (interface{}, error) {
	var productDTO data.CreateProduct

	if err := json.NewDecoder(req.Body).Decode(&productDTO); err != nil {
		return nil, err
	}

	return productDTO, nil
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
