package main

import (
	"log"
	"net/http"

	"github.com/Chennai-Golang/learning-golang/004-http-application/006-rest-api/blog"
	"github.com/gorilla/mux"
)

var port = "8000"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", blog.GetPostsHandler).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
