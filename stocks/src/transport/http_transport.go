package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/exceptions"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/services"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpServer(svc services.StockService, log log.Logger) *mux.Router {

	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(log),
		httptransport.ServerErrorEncoder(encodeErrorResp),
	}

	return r
}

func encodeErrorResp(ctx context.Context, err error, w http.ResponseWriter) {
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

	switch err {
	case exceptions.ErrorNotFoundItem:
		return http.StatusNotFound
	case exceptions.ErrorInvalidQty:
	case exceptions.ErrorInvalidSKU:
	case exceptions.ErrorInvalidParam:
	case exceptions.ErrorInvalidTitle:
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError

}
