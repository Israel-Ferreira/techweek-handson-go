package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/endpoint"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/middlewares"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/services"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/utils"
)

func NewServer(svc services.PriceService, log log.Logger) *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.JsonMiddleware)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(log),
		httptransport.ServerErrorEncoder(encodeErrorResp),
	}

	getPricesItemsHandler := httptransport.NewServer(
		endpoint.GetAllPrices(svc),
		decode,
		encodeResponse,
		options...,
	)

	getPriceItemHandler := httptransport.NewServer(
		endpoint.GetBySKU(svc),
		decodeRequestWithParam,
		encodeResponse,
		options...,
	)

	updatePriceHandler := httptransport.NewServer(
		endpoint.UpdatePrice(svc),
		decodeRequestBodyWithParam,
		encodeResponse,
		options...,
	)

	r.Handle("/prices", getPricesItemsHandler).Methods(http.MethodGet)
	r.Handle("/prices/{sku}", getPriceItemHandler).Methods(http.MethodGet)
	r.Handle("/prices/{sku}", updatePriceHandler).Methods(http.MethodPut)

	return r
}

func encodeErrorResp(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("Encode com Erro vazio")
	}

	w.WriteHeader(utils.GetStatusCodeFromError(err))

	errResp := data.ErrorMsg{
		Msg:        err.Error(),
		StatusCode: utils.GetStatusCodeFromError(err),
	}

	if err = json.NewEncoder(w).Encode(&errResp); err != nil {
		panic(err)
	}
}
