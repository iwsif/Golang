package main

import (
    //"os"
    "fmt"
    "time"
    "mime"
    "log"
    "net/http"
    //"strconv"
    "encoding/json"
    "io/ioutil" 
)

var movies []Movie

type Movie struct {
    Name string "json:{name}"
    Year string "json:{year}"
    ID string "json{id}"
    Director *Director "json:{director}"
}

type Director struct {
    Name string
    Surname string 
}

func ifempty(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
        w.Header().Set("Content-Type","application/json")
        content_type := r.Header.Get("Content-Type")
        if content_type == " " {
            http.Error(w,"Error no content-type found",http.StatusBadRequest)
            return
        }
        handler.ServeHTTP(w,r)
    })
}

func ifnotempty(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
        w.Header().Set("Content-Type","application/json")
        content_type := r.Header.Get("Content-Type")
        content,_,_ := mime.ParseMediaType(content_type)
        if content != "application/json" {
            http.Error(w,"Error not supported media type",http.StatusUnsupportedMediaType)
            return
        }
        handler.ServeHTTP(w,r)
    })
}

func homepage(w http.ResponseWriter,r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    fmt.Fprintf(w,"<h1>Welcome This is a test application</h1>")
}

func get_movies(w http.ResponseWriter,r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    if r.Method == "GET" {
        for _,movie := range movies{
            json.NewEncoder(w).Encode(movie)
        }
    }
}

func create_movie(w http.ResponseWriter,r *http.Request) {
    w.Header().Set("Content-Type","application/json") 
    if r.Method == "POST" {
        content,_ :=  ioutil.ReadAll(r.Body)
        var new_movie Movie
        json.Unmarshal(content,&movies)
        movies = append(movies,new_movie)
        for _,movie:=range movies {
            json.NewEncoder(w).Encode(movie)
        }
    }
}

func main() {

    movies = append(movies,Movie{Name:"movie1",Year:"2010",ID:"1",Director:&Director{Name:"Directors1_name",Surname:"Directors1_surname"}})
    movies = append(movies,Movie{Name:"movie2",Year:"2011",ID:"2",Director:&Director{Name:"Directors2_name",Surname:"Directors2_surname"}})
    movies = append(movies,Movie{Name:"movie3",Year:"2012",ID:"3",Director:&Director{Name:"Directors3_name",Surname:"Directors3_surname"}})
    
    fmt.Println("Welcome this is a test api-server build with http package")
    router := http.NewServeMux()
    router.Handle("/",ifempty(ifnotempty(http.HandlerFunc(homepage))))
    router.HandleFunc("/api/movies",get_movies)
    router.HandleFunc("/api/movie",create_movie)
    log.Printf("Starting server..")
    time.Sleep(time.Second*2)
    log.Printf("Server is listening on port 20000 ")
    log.Fatal(http.ListenAndServe(":20000",router))
}