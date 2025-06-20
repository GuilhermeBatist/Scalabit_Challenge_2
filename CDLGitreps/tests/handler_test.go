package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GuilhermeBatist/Scalabit_Challenge_2/CDLGitreps/handlers"
	"github.com/GuilhermeBatist/Scalabit_Challenge_2/CDLGitreps/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test CreateRepoHandler
func TestCreateRepoHandler(t *testing.T) {
	// Success case
	services.CreateRepo = func(repo services.Repo) error {
		return nil
	}
	router := gin.Default()
	router.POST("/repos", func(c *gin.Context) {
		handlers.CreateRepoHandler(c.Writer, c.Request)
	})

	body := bytes.NewBuffer([]byte(`{"name":"myrepo"}`))
	req, _ := http.NewRequest(http.MethodPost, "/repos", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Invalid JSON
	body = bytes.NewBuffer([]byte(`not-json`))
	req, _ = http.NewRequest(http.MethodPost, "/repos", body)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Service error
	services.CreateRepo = func(repo services.Repo) error {
		return assert.AnError
	}
	body = bytes.NewBuffer([]byte(`{"name":"myrepo"}`))
	req, _ = http.NewRequest(http.MethodPost, "/repos", body)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Test DeleteRepoHandler
func TestDeleteRepoHandler(t *testing.T) {
	// Successful delete
	services.DeleteRepo = func(name string) error {
		return nil
	}
	router := gin.Default()
	router.DELETE("/repos/:name", func(c *gin.Context) {
		handlers.DeleteRepoHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest(http.MethodDelete, "/repos/myrepo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Missing repo name
	req, _ = http.NewRequest(http.MethodDelete, "/repos/", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Service error
	services.DeleteRepo = func(name string) error {
		return assert.AnError
	}
	req, _ = http.NewRequest(http.MethodDelete, "/repos/myrepo", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Test ListReposHandler
func TestListReposHandler(t *testing.T) {
	// Success
	services.ListRepos = func() ([]services.Repo, error) {
		return []services.Repo{{Name: "repo1"}, {Name: "repo2"}}, nil
	}
	router := gin.Default()
	router.GET("/repos", func(c *gin.Context) {
		handlers.ListReposHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest(http.MethodGet, "/repos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var repos []services.Repo
	err := json.Unmarshal(w.Body.Bytes(), &repos)
	assert.NoError(t, err)
	assert.Len(t, repos, 2)

	// No repos found
	services.ListRepos = func() ([]services.Repo, error) {
		return []services.Repo{}, nil
	}
	req, _ = http.NewRequest(http.MethodGet, "/repos", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Service error
	services.ListRepos = func() ([]services.Repo, error) {
		return nil, assert.AnError
	}
	req, _ = http.NewRequest(http.MethodGet, "/repos", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
