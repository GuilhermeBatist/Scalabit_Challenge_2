package services

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/CDLGitreps/services/comment"
	"github.com/yourusername/CDLGitreps/services/repo"
	"github.com/yourusername/CDLGitreps/services/user"
)

// StartServer initializes the HTTP server and routes
func StartServer() {
	fmt.Println("Starting server on :8080")

	router := mux.NewRouter()

	// Register routes for comments
	comment.RegisterRoutes(router)

	// Register routes for repositories
	repo.RegisterRoutes(router)

	// Register routes for users
	user.RegisterRoutes(router)

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// main function to start the server
func main() {
	StartServer()
}
