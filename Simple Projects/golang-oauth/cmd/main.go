package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
)


func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = ":3000"
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("hello to root endpoint"))})
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {fmt.Fprintf(w, "Welcome to home endpoint")})
	log.Fatal(http.ListenAndServe(port, r))
}
