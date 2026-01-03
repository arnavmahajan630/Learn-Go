package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Root of the Project :) %q", html.EscapeString(r.URL.Path))
}
func handleHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Hi of the Project :) %q", html.EscapeString(r.URL.Path))
}
func main() {

	fmt.Println("hello to docker :) ")

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/hi", handleHi)

	log.Println("Server started successfully on Port: 3000")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Error starting the server %v", err)
	}
}
