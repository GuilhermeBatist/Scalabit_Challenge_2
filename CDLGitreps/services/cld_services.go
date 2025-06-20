package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

// Repo represents a GitHub repository
type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private,omitempty"`
}

// GitHubClient is the client for interacting with GitHub API
type GitHubClient struct {
	client *github.Client
}

// NewGitHubClient creates a new GitHub client with authentication
func NewGitHubClient() *GitHubClient {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		panic("GITHUB_TOKEN environment variable is not set")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return &GitHubClient{client: client}
}

// CreateRepo creates a new repository
func (gh *GitHubClient) CreateRepo(repo Repo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	githubRepo := &github.Repository{
		Name:        &repo.Name,
		Description: &repo.Description,
		Private:     &repo.Private,
	}

	_, _, err := gh.client.Repositories.Create(ctx, "", githubRepo)
	if err != nil {
		return fmt.Errorf("failed to create repository: %w", err)
	}
	return nil
}

// DeleteRepo deletes a repository by name
func (gh *GitHubClient) DeleteRepo(repoName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	owner := "your-github-username" // Replace with your GitHub username
	_, err := gh.client.Repositories.Delete(ctx, owner, repoName)
	if err != nil {
		if _, ok := err.(*github.ErrorResponse); ok {
			return fmt.Errorf("failed to delete repository: %w", err)
		}
		return errors.New("repository not found or you do not have permission to delete it")
	}
	return nil
}

// ListRepos lists all repositories for the authenticated user
func (gh *GitHubClient) ListRepos() ([]Repo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	repos, _, err := gh.client.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list repositories: %w", err)
	}

	var repoList []Repo
	for _, r := range repos {
		repoList = append(repoList, Repo{
			Name:        *r.Name,
			Description: r.GetDescription(),
			Private:     r.GetPrivate(),
		})
	}
	return repoList, nil
}

// CreateRepo creates a new repository
func CreateRepo(repo Repo) error {
	client := NewGitHubClient()
	return client.CreateRepo(repo)
}

// DeleteRepo deletes a repository by name
func DeleteRepo(repoName string) error {
	client := NewGitHubClient()
	return client.DeleteRepo(repoName)
}

// ListRepos lists all repositories for the authenticated user
func ListRepos() ([]Repo, error) {
	client := NewGitHubClient()
	return client.ListRepos()
}
