package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GuilhermeBatist/Scalabit_Challenge_2/CDLGitreps/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/repos", func(c *gin.Context) {
		handlers.CreateRepoHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest("POST", "/repos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
func TestDeleteRepoHandler(t *testing.T) {
	router := gin.Default()
	router.DELETE("/repos/:owner/:repo", func(c *gin.Context) {
		handlers.DeleteRepoHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest("DELETE", "/repos/octocat/Hello-World", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}
func TestListReposHandler(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", func(c *gin.Context) {
		handlers.ListReposHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest("GET", "/repos/list", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestListReposHandlerWithInvalidPage(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", func(c *gin.Context) {
		handlers.ListReposHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest("GET", "/repos/list?page=invalid&limit=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithInvalidLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", func(c *gin.Context) {
		handlers.ListReposHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest("GET", "/repos/list?page=1&limit=invalid", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithNegativePage(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", func(c *gin.Context) {
		handlers.ListReposHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest("GET", "/repos/list?page=-1&limit=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithNegativeLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", func(c *gin.Context) {
		handlers.ListReposHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest("GET", "/repos/list?page=1&limit=-10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithZeroPage(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", func(c *gin.Context) {
		handlers.ListReposHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest("GET", "/repos/list?page=0&limit=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithZeroLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", func(c *gin.Context) {
		handlers.ListReposHandler(c.Writer, c.Request)
	})

	req, _ := http.NewRequest("GET", "/repos/list?page=1&limit=0", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func adaptHandler(h func(http.ResponseWriter, *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c.Writer, c.Request)
	}
}

func TestListReposHandlerWithEmptyQuery(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", adaptHandler(handlers.ListReposHandler))

	req, _ := http.NewRequest("GET", "/repos/list", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
func TestListReposHandlerWithInvalidQuery(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", adaptHandler(handlers.ListReposHandler))

	req, _ := http.NewRequest("GET", "/repos/list?invalid=query", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithValidQuery(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", adaptHandler(handlers.ListReposHandler))

	req, _ := http.NewRequest("GET", "/repos/list?query=valid", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
func TestListReposHandlerWithValidQueryAndZeroLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", adaptHandler(handlers.ListReposHandler))

	req, _ := http.NewRequest("GET", "/repos/list?query=valid&page=1&limit=0", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithValidQueryAndNegativeLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", adaptHandler(handlers.ListReposHandler))

	req, _ := http.NewRequest("GET", "/repos/list?query=valid&page=1&limit=-10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithValidQueryAndZeroPage(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", adaptHandler(handlers.ListReposHandler))

	req, _ := http.NewRequest("GET", "/repos/list?query=valid&page=0&limit=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithValidQueryAndEmptyPage(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", adaptHandler(handlers.ListReposHandler))

	req, _ := http.NewRequest("GET", "/repos/list?query=valid&page=&limit=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
func TestListReposHandlerWithValidQueryAndEmptyLimit(t *testing.T) {
	router := gin.Default()
	router.GET("/repos/list", adaptHandler(handlers.ListReposHandler))

	req, _ := http.NewRequest("GET", "/repos/list?query=valid&page=1&limit=", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
