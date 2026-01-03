package main

import (
	"fmt"
	"html"
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
}
