package main

import (
	"log"
	"net/http"

	"github.com/OceanWhisperer/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterStoreRzoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe(":8080", route))
}