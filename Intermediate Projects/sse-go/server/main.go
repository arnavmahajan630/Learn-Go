package main

import (
	"fmt"
	"log"
	"net/http"
)


func events(w http.ResponseWriter, r * http.Request) {

	// Set necessary headers
	w.Header().Set("Cotnent-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")


}

func home(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Wlecome to home endpoint! ")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}