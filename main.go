package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yuttasakcom/go-apis-sample/handlers"
)

func main() {

	h := handlers.Router()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Print("Server running at port:", port)

	err := http.ListenAndServe(":"+port, h)

	if err != nil {
		log.Fatal(err)
	}
}
