package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Router handlers
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	r.HandleFunc("/profile/{username}", profile).Methods("GET")
	return r
}
