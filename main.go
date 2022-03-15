package main

import (
	"as-golangweb/handler"
	"log"
	"net/http"
)

func main() {

	// initiate mux server
	mux := http.NewServeMux()

	// routing
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/form", handler.FormHandler)
	mux.HandleFunc("/result", handler.ProcessHandler)

	// for static file
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// printing server running
	log.Println("server is running on :8080")

	// running server
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))

}
