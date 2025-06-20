package main

import (
	"fmt"
	"net/http"

	"github.com/GuilhermeBatist/Scalabit_Challenge_2/CDLGitreps/handlers"
)

func main() {
	fmt.Println("Starting server on :8080")

	CreateRepoHandler := handlers.CreateRepoHandler
	DeleteRepoHandler := handlers.DeleteRepoHandler
	ListReposHandler := handlers.ListReposHandler
	http.HandleFunc("/repos", CreateRepoHandler)
	http.HandleFunc("/repos/", DeleteRepoHandler)
	http.HandleFunc("/repos/list", ListReposHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
