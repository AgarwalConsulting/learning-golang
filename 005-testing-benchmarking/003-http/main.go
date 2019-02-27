package main

import (
	"log"
	"net/http"

	"github.com/AgarwalConsulting/learning-golang/004-http-application/006-rest-api/blog"
	"github.com/gorilla/mux"

	_ "net/http/pprof"
)

var port = "8000"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", blog.GetPostsHandler).Methods("GET")

	go func() {
		log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
	}()

	log.Fatal(http.ListenAndServe(":"+port, r))
}
