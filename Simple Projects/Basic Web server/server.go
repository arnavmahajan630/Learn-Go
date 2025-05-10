package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Error Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "404 Error Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Hello")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
		return
	}
	if r.Method != "POST" {
		fmt.Fprintf(w, "Access Denied")
		return
	}
	fmt.Fprintln(w, "Post request Successfull")
	name := r.FormValue("Namelabel>")
	address := r.FormValue("Addresslabel>")
	fmt.Fprintln(w, "The name is: ", name)
	fmt.Fprintln(w, "The Address is: ", address)
}

func main() {
	fmt.Println("Hello World")
	server := http.FileServer(http.Dir("./static"))
	http.Handle("/", server)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)
	fmt.Println("Server Started at Port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
