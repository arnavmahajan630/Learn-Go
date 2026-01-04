package main

import (
	"fmt"
	"log"
	"net/http"
)


func setupApi() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}

func main() {
	setupApi()
	fmt.Println("Server starting on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}