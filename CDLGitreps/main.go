package main

import (
	//"encoding/json"
	//"log"
	"net/http"
	"fmt"
)



func main() {
	fmt.Println("Starting server on :8080")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"returning all comments\n")
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w,"returning a single comment for comment with id: %s\n", id)
	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"creating a new comment\n")
	})

	if err :=  http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}