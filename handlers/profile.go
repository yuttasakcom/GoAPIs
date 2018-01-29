package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func profile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	w.Write([]byte(username))
}
