package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func movieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, movie := range movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movie)

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

	w.Write([]byte(fmt.Sprintf("Movie with ID %s not found", id)))
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
	movies = append(movies, Movie{ID: "2", Isbn: "67890", Title: "My second movie", Director: &Director{FirstName: "Jane", LastName: "Doe"}})

	r.HandleFunc("/movies", moviesHandler).Methods("GET")
	r.HandleFunc("/movie/{id}", movieHandler).Methods("GET")

	fmt.Printf("Listening on port 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
