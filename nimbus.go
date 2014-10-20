package main

import (
    //go
    "fmt"
    "log"
    "net/http"
    //pkgs
    "github.com/gorilla/mux"
    //me
    "nimbus/app/home"
    "nimbus/app/postal"
)

func main() {
    fmt.Println("Nimbus forming...")
    //router
    router := mux.NewRouter()

    //route paths
    router.HandleFunc("/", home.Root).Methods("GET")
    api := router.PathPrefix("/api").Subrouter()

    //posts
    posts_router := api.PathPrefix("/posts").Subrouter()
    posts_router.HandleFunc("/", postal.GetPosts)
    
    http.Handle("/",router)
    //end of route management
    
    //static file server
    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    //listen
    fmt.Println("Listening...")
    err := http.ListenAndServe(":5000", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}