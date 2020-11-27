package main

import (
	"github.com/gorilla/mux"
    "net/http"
    "log"
    "encoding/json"
    //"math/rand"
)

type Movie struct {
	Id  string  "json:{id}"
	Title string "json:{title}"
	Director *Director "json:{director}"
}

type Director struct {
	Name string "json:{name}"
	Surname string "json:{surname}"
}

var movies []Movie 

func get_movies(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func get_movie(w http.ResponseWriter ,r *http.Request) {

}

func create_movie(w http.ResponseWriter,r *http.Request) {

}

func update_movie(w http.ResponseWriter,r *http.Request) {

}

func delete_movie(w http.ResponseWriter,r *http.Request) {

}

func main() {

	movies = append(movies,Movie{Id:"0",Title:"This is the first movie",Director:&Director{Name:"Joseph",Surname:"Kiriakopoulos"}})
	movies = append(movies,Movie{Id:"1",Title:"This is the second movie",Director:&Director{Name:"kiriakos",Surname:"kiriakopoulos"}})

	router := mux.NewRouter()
	router.HandleFunc("/api/movies",get_movies).Methods("GET")
	router.HandleFunc("/api/movie/{id}",get_movie).Methods("GET")
	router.HandleFunc("/api/movies",create_movie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",update_movie).Methods("SET")
	router.HandleFunc("/api/movie/{id}",delete_movie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000",router))
}