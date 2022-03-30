package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/endpoints"
	"github.com/Israel-Ferreira/techweek-hands-on/products/exceptions"
	"github.com/Israel-Ferreira/techweek-hands-on/products/middlewares"
	"github.com/Israel-Ferreira/techweek-hands-on/products/services"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc services.ProductService, log log.Logger) *mux.Router {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(log),
		httptransport.ServerErrorEncoder(encodeServerError),
	}

	getProductsHandler := httptransport.NewServer(
		endpoints.GetProductsEndpoint(svc),
		decode,
		encodeResponse,
		options...,
	)

	getProductHandler := httptransport.NewServer(
		endpoints.GetProductEndpoint(svc),
		decodeRequestWithParam,
		encodeResponse,
		options...,
	)

	createProductHandler := httptransport.NewServer(
		endpoints.CreateProduct(svc),
		decodeProductJsonBody,
		encodeResponse,
		options...,
	)

	deleteProductHandler := httptransport.NewServer(
		endpoints.DeleteProduct(svc),
		decodeRequestWithParam,
		encodeResponse,
		options...,
	)

	r := mux.NewRouter()

	r.Use(middlewares.JsonMiddleware)

	r.Handle("/products", getProductsHandler).Methods(http.MethodGet)
	r.Handle("/products", createProductHandler).Methods(http.MethodPost)
	r.Handle("/products/{sku}", getProductHandler).Methods(http.MethodGet)
	r.Handle("/products/{sku}", deleteProductHandler).Methods(http.MethodDelete)

	return r
}

func encodeServerError(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("Encode com Erro vazio")
	}

	w.WriteHeader(getStatusCodeFromError(err))

	errResp := data.ErrorMsg{
		Msg:        err.Error(),
		StatusCode: getStatusCodeFromError(err),
	}

	if err = json.NewEncoder(w).Encode(&errResp); err != nil {
		panic(err)
	}

}

func getStatusCodeFromError(err error) int {
	if err == exceptions.ErrorNotFoundProduct {
		return http.StatusNotFound
	} else if err == exceptions.ErrorBodyIsNotValid {
		return http.StatusBadRequest
	} else {
		return http.StatusInternalServerError
	}
}
