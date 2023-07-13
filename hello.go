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
	Id       string    `json:"Id"`
	Isbn     string    `json:"Isbn"`
	Title    string    `json:"Title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"fn"`
	Lastname  string `json:"ln"`
}

var movies []Movie

func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie( w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content type ", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content type", "application /json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode((&movie))
	movie.Id = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode((movie))
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "43244", Title: "movie rrr", Director: &Director{Firstname: " james", Lastname: "bond"}})
	movies = append(movies, Movie{Id: "2", Isbn: "43222", Title: "movie xyz", Director: &Director{Firstname: " selena", Lastname: "gomez"}})
	movies = append(movies, Movie{Id: "3", Isbn: "43233", Title: "movie muv", Director: &Director{Firstname: " rakesh", Lastname: "bhatt"}})
	movies = append(movies, Movie{Id: "4", Isbn: "43266", Title: "movie tyz", Director: &Director{Firstname: " james", Lastname: "bond"}})
	movies = append(movies, Movie{Id: "5", Isbn: "43255", Title: "movie lmn", Director: &Director{Firstname: " john", Lastname: "doe"}})

	r.HandleFunc("/movies", getmovie).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("Delete")

	fmt.Printf("starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
