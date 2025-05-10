package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"fname"`
	Lastname  string `json:"lname"`
}

var Movies []Movie

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range Movies {
		if item.Id == params["id"] {
			Movies = append(Movies[:idx], Movies[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Movies)
}

func getmovieid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range Movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Movie Movie
	_ = json.NewDecoder(r.Body).Decode(&Movie)
	Movie.Id = strconv.Itoa((rand.Intn(1000000000)))
	Movies = append(Movies, Movie)
	json.NewEncoder(w).Encode(Movie)
}

func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, val := range Movies {
		if val.Id == params["id"] {
			Movies = append(Movies[:i], Movies[i+1:]...)
			break;
		}
	}

	var Movie Movie
	_ = json.NewDecoder(r.Body).Decode(&Movie)
	Movie.Id = strconv.Itoa(rand.Intn(10000000))
	Movies = append(Movies, Movie)
	json.NewEncoder(w).Encode(Movie)
}
func main() {
	Movies = append(Movies, Movie{Id: "1", Isbn: "121324", Title: "Movie1", Director: &Director{Firstname: "Jhon", Lastname: "Doe"}}, Movie{Id: "2", Isbn: "43244", Title: "Movie2", Director: &Director{Firstname: "Sol", Lastname: "Gel"}})
	router := mux.NewRouter()
	// paths
	router.HandleFunc("/movies", getmovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getmovieid).Methods("GET")
	router.HandleFunc("/movies", createmovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	fmt.Println("Starting Server at Port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
