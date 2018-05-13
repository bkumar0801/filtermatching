package main

import (
	"log"
	"net/http"
	"os"

	h "github.com/filtermatching/handlers"
)

func main() {
	var filename string
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/filter", h.HandleFilter(filename))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
