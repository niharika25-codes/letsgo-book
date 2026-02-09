package main

import (
	"log"
	"net/http"
)

func main() {

	// router or servemux
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	log.Print("starting server on :4000")

	// server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
