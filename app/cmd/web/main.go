package main

import (
    "log"
    "net/http"
)

func main() {

    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/user", showUser)
    mux.HandleFunc("/user/create", createUser)

    fileServer := http.FileServer(http.Dir("../../ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}