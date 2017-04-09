package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/zmetcalf/sportsbook/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	var dir string

	flag.StringVar(&dir, "dir", "./app/dist", "the directory to serve files from. Set to dist directory")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/home", handlers.HomeHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
