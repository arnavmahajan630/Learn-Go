package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


type Message struct {
	Satus string `json:"status"`
	Body string `json:"body"`
}

func endPointHander(w  http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message :=  Message{
		Satus: "Successful",
		Body: "Hey You have Hit the Endpoint. How may I help you ?",
	}

	err := json.NewEncoder(w).Encode(&message)
	if err != nil {
		log.Fatal(err)
		return
	}
	
}

func main() {
	http.HandleFunc("/ping", rateLimiter(endPointHander))
	http.HandleFunc("/", func (w http.ResponseWriter, r * http.Request){
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello to the headpoint :) ")
	})
	fmt.Println("Server starting to listen on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}