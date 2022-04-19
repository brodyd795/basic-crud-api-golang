package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func movieHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "My first movie", Director: &Director{FirstName: "John", LastName: "Doe"}})

	r.HandleFunc("/movies", movieHandler).Methods("GET")

	fmt.Printf("Listening on port 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
