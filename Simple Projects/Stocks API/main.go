package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Ocean-Whisperer/Learn-Go/Simple-Projects/Stocks-API/router"
)


func main() {
	r := router.Router()
	fmt.Println("Stating server on Port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}