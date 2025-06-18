package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GuilhermeBatist/Scalabit_Challenge/Returnnpulls/handlers"
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
