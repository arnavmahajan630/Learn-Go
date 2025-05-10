package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OceanWhisperer/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Started successfully on Port 8080")
	route := mux.NewRouter()
	routes.RegisterStoreRzoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe(":8080", route))
	fmt.Println("Started successfully on Port 8080")
}