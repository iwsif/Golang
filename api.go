package main


import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    //"math/rand"
    "io/ioutil"
    "strconv"
    "time"
    "github.com/gorilla/mux"
    //"io"
)

var movies []Movie

type Movie struct {
    Name string "json:{name}"
    Year string "json:{year}"
    ID string "json:{id}"
    Director *Director "json:{director}"
}

type Director struct {
    Name string "json:{Name}"
    Surname string "json:{Surname}"
}

//Handling each route with functions,writing with writer and read request with reader

func get_movies(w http.ResponseWriter,r *http.Request) {
    w.Header().Set("Content-Type","application/json")

    for _, movie := range movies {
        json.NewEncoder(w).Encode(movie)
    }
}

func get_movie(w http.ResponseWriter,r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    id := mux.Vars(r)
    for _,movie:=range movies {
        if movie.ID == id["id"] {
            json.NewEncoder(w).Encode(movie)
            break
        }
    }
}

func add_movie(w http.ResponseWriter,r *http.Request) {

    w.Header().Set("Content-Type","application/json")
    content,_:= ioutil.ReadAll(r.Body)
    var users_movie Movie
    json.Unmarshal(content,&users_movie)
    movies = append(movies,users_movie)
    for _,movie := range movies {
        json.NewEncoder(w).Encode(movie)
    }
}

func delete_movie(w http.ResponseWriter,r *http.Request) {    
    w.Header().Set("Content-Type","application/json")
    user_choice := mux.Vars(r)
    
    for pointer,movie := range movies {
        if movie.ID == user_choice["id"] {
            movies = append(movies[:pointer],movies[pointer+1:]...)
            break
        }
    }
}

func update_movie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    id := mux.Vars(r)
    for i,movie := range movies {
        if movie.ID == id["id"] {
            movies = append(movies[:i],movies[i+1:]...)
            break
        }
    }
    new_movie,_ := ioutil.ReadAll(r.Body)
    json.Unmarshal(new_movie,&movies)
    for _,movie := range movies {
        json.NewEncoder(w).Encode(movie)
    }
}

func main() {
    //adding some data for fun

    movies = append(movies,Movie{Name:"Movie1",Year:"2000",ID:"1",Director:&Director{Name:"Director1_name",Surname:"Director1_surname"}})
    movies = append(movies,Movie{Name:"Movie2",Year:"2010",ID:"2",Director:&Director{Name:"Director2_name",Surname:"Director2_surname"}})
    movies = append(movies,Movie{Name:"Movie3",Year:"2011",ID:"3",Director:&Director{Name:"Director3_name",Surname:"Director3_surname"}})
    movies = append(movies,Movie{Name:"Movie4",Year:"2012",ID:"4",Director:&Director{Name:"Director4_name",Surname:"Director4_surname"}})

    fmt.Println("This is a simple rest api based on gorilla-mux router")
    fmt.Println("Starting the router..")
    fmt.Printf("Choose port to start the server:")
    var port int 
    fmt.Scanf("%d",&port)
    time.Sleep(time.Second*2)
    fmt.Println("Listening on port " + strconv.Itoa(port))
    //Starting new router
    
    router := mux.NewRouter()
    router.HandleFunc("/api/movies",get_movies).Methods("GET")
    router.HandleFunc("/api/movie/{id}",get_movie).Methods("GET")
    router.HandleFunc("/api/movie",add_movie).Methods("POST")
    router.HandleFunc("/api/movie/{id}",delete_movie).Methods("DELETE")
    router.HandleFunc("/api/movie/{id}",update_movie).Methods("PUT")

    //Serve and listen at port:

    log.Fatal(http.ListenAndServe(":"+ strconv.Itoa(port),router))
}
