package handlers

import (
	"github.com/gorilla/mux"
)

// Router handlers
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	return r
}
