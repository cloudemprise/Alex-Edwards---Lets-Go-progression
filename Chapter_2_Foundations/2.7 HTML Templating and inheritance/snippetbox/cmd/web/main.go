package main

/* Alex Edwards : Let's Go

Chapter 2 Foundations

Section 2.7 HTML Templating and inheritance */

import (
	"log"
	"net/http"
)

///

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
