package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello net/http!\n"))
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	w.Write([]byte(fmt.Sprintf("Hi, %s!", name)))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloGoHandler)
	r.HandleFunc("/{name}", greetingHandler)

	/*
	   Matchers:
	   r.HandleFunc("/products", ProductsHandler).
	       Host("www.example.com").                // Only match a specific domain
	       Methods("GET", "PUT").                  // Only match GET+PUT methods
	       Schemes("http")                         // Only match the http scheme
	*/

	log.Fatal(http.ListenAndServe(":8080", r))
}
