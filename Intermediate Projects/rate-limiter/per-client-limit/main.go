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

func endPointHander(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := Message {
		Satus: "Successful",
		Body: "Hi! Reached the endpoint :)",
	}
	if err := json.NewEncoder(w).Encode(&message); err != nil {
		return
	}
}

func main() {

	http.HandleFunc("/endpoint", perClientRateLimit(endPointHander))
	fmt.Println("Server started on Port 3000 :| ")
	log.Fatal(http.ListenAndServe(":3000", nil))

}