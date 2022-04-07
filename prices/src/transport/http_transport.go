package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/middlewares"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/utils"
)

func NewServer() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.JsonMiddleware)

	r.Handle("/prices", nil).Methods(http.MethodGet)
	r.Handle("/prices/{id}", nil).Methods(http.MethodGet)
	r.Handle("/prices/{id}", nil).Methods(http.MethodPut)

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
