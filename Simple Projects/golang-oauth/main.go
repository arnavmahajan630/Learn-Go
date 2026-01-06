package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func setupRoutes(m *http.ServeMux) {
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to home point")
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not set")
	}

	m := http.NewServeMux()
	setupRoutes(m)

	log.Fatal(http.ListenAndServe(port, nil))
}
