package transport

import "github.com/gorilla/mux"

func NewServer() *mux.Router {
	r := mux.NewRouter()

	return r
}
