package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GuilhermeBatist/Scalabit_Challenge_2/Returnnpulls/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/*
	func TestGetOpenPullRequests(t *testing.T) {
		router := gin.Default()
		router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

		req, _ := http.NewRequest("GET", "/repos/octocat/Hello-World/pullrequests?limit=3", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	}
*/
func TestGetOpenPullRequests(t *testing.T) {
	limit := 21 // 👈 change this to any number you want

	router := gin.Default()
	router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/repos/octocat/Hello-World/pullrequests?limit=%d", limit), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetOpenPullRequestsInvalidLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

	req, _ := http.NewRequest("GET", "/repos/octocat/Hello-World/pullrequests?limit=invalid", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestGetOpenPullRequestsInvalidOwner(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

	req, _ := http.NewRequest("GET", "/repos/invalid_owner/Hello-World/pullrequests?limit=3", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetOpenPullRequestsInvalidRepo(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

	req, _ := http.NewRequest("GET", "/repos/octocat/invalid_repo/pullrequests?limit=3", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetOpenPullRequestsWithoutLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

	req, _ := http.NewRequest("GET", "/repos/octocat/Hello-World/pullrequests", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetOpenPullRequestsWithNegativeLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

	req, _ := http.NewRequest("GET", "/repos/octocat/Hello-World/pullrequests?limit=-5", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestGetOpenPullRequestsWithZeroLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

	req, _ := http.NewRequest("GET", "/repos/octocat/Hello-World/pullrequests?limit=0", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestGetOpenPullRequestsWithNegativePage(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

	req, _ := http.NewRequest("GET", "/repos/octocat/Hello-World/pullrequests?page=-1&limit=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestGetOpenPullRequestsWithInvalidPage(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/:owner/:repo/pullrequests", handlers.GetOpenPullRequests)

	req, _ := http.NewRequest("GET", "/repos/octocat/Hello-World/pullrequests?page=invalid&limit=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
