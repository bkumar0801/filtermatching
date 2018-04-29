package main

import (
	"log"
	"net/http"

	h "github.com/filtermatching/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/filter", h.HandleFilter)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
