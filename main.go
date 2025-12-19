package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"idbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie = []Movie{}

// movies = append(movies, Movie)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", getHealthCheck)

	r.HandleFunc("/movies", getMovies).Methods("GET")

	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")

	r.HandleFunc("/movies", createMovie).Methods("POST")

	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")

	// data, err := json.MarshalIndent(movies, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(data))

	log.Fatal(http.ListenAndServe(":8000", r))

}

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	// Route for checking server health
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is up and running"))
	return
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	// Simple get route for getting all the movies in Database.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	return
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// Get a single movie by ID => Easy one
	params := mux.Vars(r)
	fmt.Println(params)
	w.Header().Set("Content-Type", "application/json")
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "movie not found"})
	return
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// Create route for creating a movie. Getting movie data from request body.
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}
	movie.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	movies = append(movies, movie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Update movie by ID. Getting required fields from request body to update.
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	for i, item := range movies {
		if item.ID == params["id"] {
			var updated Movie
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "invalid request body",
				})
				return
			}

			updated.ID = item.ID
			movies[i] = updated

			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "movie not found",
	})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Delete movie by ID
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "movie not found",
	})
}
