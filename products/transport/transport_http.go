package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/exceptions"
	"github.com/Israel-Ferreira/techweek-hands-on/products/middlewares"
	"github.com/Israel-Ferreira/techweek-hands-on/products/services"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc services.ProductService, log log.Logger) *mux.Router {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(log),
		httptransport.ServerErrorEncoder(encodeServerError),
	}

	fmt.Println(options)

	r := mux.NewRouter()

	r.Use(middlewares.JsonMiddleware)

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
