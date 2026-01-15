package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/didip/tollbooth"
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
	message := Message {
		Satus: "Request Failed",
		Body: "The API is hit capacity try again later :( ",
	}

	js , _ := json.Marshal(message)

	// tollbooth initialization
	tollBoothLimiter := tollbooth.NewLimiter(1, nil)
	tollBoothLimiter.SetMessageContentType("application/json")
	tollBoothLimiter.SetMessage(string(js))

	http.Handle("/endpoint", tollbooth.LimitFuncHandler(tollBoothLimiter, endPointHander))
	fmt.Println("Server started on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}