package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	addr string
	staticDir string
}

var cfg config

func main() {
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network addresss")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path ti static assets")
	
	flag.Parse()
	// router or servemux
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("starting server on %s", cfg.addr)

	// server
	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
