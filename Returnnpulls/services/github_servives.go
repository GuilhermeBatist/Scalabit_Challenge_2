package services

import (
	"context"
	"os"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

func getClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: os.Getenv("GITHUB_TOKEN"),
	})
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func FetchOpenPRs(owner, repo string, limit int) ([]*github.PullRequest, error) {
	client := getClient()
	opts := &github.PullRequestListOptions{
		State: "open",
		ListOptions: github.ListOptions{
			PerPage: limit,
		},
	}
	prs, _, err := client.PullRequests.List(context.Background(), owner, repo, opts)
	return prs, err
}
