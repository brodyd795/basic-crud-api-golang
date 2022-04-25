package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func moviesHandler(w http.ResponseWriter, _ *http.Request) {
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

func createHandler(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for index, movie := range movies {
		if movie.ID == id {
			movies[index] = movies[len(movies)-1]
			movies = movies[:len(movies)-1]
			break
		}
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updatedMovie Movie
	json.NewDecoder(r.Body).Decode(&updatedMovie)

	for index, movie := range movies {
		if movie.ID == id {
			movies[index] = updatedMovie
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
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
	r.HandleFunc("/movie", createHandler).Methods("POST")
	r.HandleFunc("/movie/{id}", deleteHandler).Methods("DELETE")
	r.HandleFunc("/movie/{id}", updateHandler).Methods("PUT")

	fmt.Printf("Listening on port 8080...\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
