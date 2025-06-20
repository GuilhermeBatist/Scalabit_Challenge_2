package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeBatist/Scalabit_Challenge_2/CDLGitreps/services"
)

// CreateRepoHandler handles the creation of a new repository
func CreateRepoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var repo services.Repo
	if err := json.NewDecoder(r.Body).Decode(&repo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := services.CreateRepo(repo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(repo)
}

// DeleteRepoHandler handles the deletion of a repository by name
func DeleteRepoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	repoName := r.URL.Path[len("/repos/"):]
	if repoName == "" {
		http.Error(w, "Repository name is required", http.StatusBadRequest)
		return
	}

	if err := services.DeleteRepo(repoName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListReposHandler handles listing all repositories
func ListReposHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	repos, err := services.ListRepos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(repos) == 0 {
		http.Error(w, "No repositories found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(repos)
}
