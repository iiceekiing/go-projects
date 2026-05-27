package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gprilla/mux"
)



	type Movie struct{
		ID string `json: "id"`
		Isbn string `json: "isbn"`
		Title string `json: "title"`
		Director *Director `json: "director"`
	}


	//Note: Struct in Golang is: Object in Javascript, || Dictionary in Python.
	type Director struct{
		FirstName string `json: "firstname"`
		LastName string `json: "lastname"`
	}


	
	//Slice "[]" of type Movie Object/Dictionary (Note: Slice is array (in Js) || list(in Py) 😂).
var movies []Movie



func getMovies(w hhtp.ResponseWriter, r *hhtp.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Var(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEnconder(w).Encode(item)
			return  
		}
	}
}



func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strConv.Itoa(rand.Intn(100000000))

	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	return
}



func updateMovie(w http.ResponseWriter, r http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEnconder(w).Encode(movie)
			return
		}
	}

}



	func main(){
		// r := mux.NewRouter()

		movies = append(movies, Movie{ID: "1", Isbn: "778224", Title: "Forge Technologies", Director: &Director{FirstName: "iiceekiing", LastName: "Miracle"} });

		movies = append(movies, Movie{ID: "2", Isbn: "2244778", Title: "BlockFuse Labs", Director: &Director{FirstName: "Scarface", LastName: "Joseph"}})

		fmt.Println(movies[10])

		r.HandleFunc("/movies", getMovies).Methods("GET")
		r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
		r.HandleFunc("/movies", createMovie).Methods("POST")
		r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
		r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

		fmt.Printf("Starting server on port 8000...\n")
		log.Fatal(http.ListenAndServe(":8000",r))
	}
