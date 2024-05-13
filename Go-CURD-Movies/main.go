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

type Movie struct {
	ID       string    `json:"id"`
	Isdn     string    `json:"isdn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	parms := mux.Vars(r)

	for index, item := range movies {
		if item.ID == parms["id"] {
			// slice movie start from 0 to index
			// append slice movie from index+1 to end
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	parms := mux.Vars(r)

	for _, item := range movies {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	parms := mux.Vars(r)

	for index, item := range movies {
		if item.ID == parms["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			movie.ID = parms["id"]
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}

}
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isdn:  "123",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isdn:  "456",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Steve",
			Lastname:  "Smith",
		},
	})

	movies = append(movies, Movie{
		ID:    "3",
		Isdn:  "789",
		Title: "Movie Three",
		Director: &Director{
			Firstname: "Jane",
			Lastname:  "Doe",
		},
	})

	movies = append(movies, Movie{
		ID:    "4",
		Isdn:  "101",
		Title: "Movie Four",
		Director: &Director{
			Firstname: "Mike",
			Lastname:  "Smith",
		},
	})

	movies = append(movies, Movie{
		ID:    "5",
		Isdn:  "112",
		Title: "Movie Five",
		Director: &Director{
			Firstname: "Sara",
			Lastname:  "Doe",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server is running on port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
