package handlers

import (
	"net/http"
	"strconv"

	"github.com/GuilhermeBatist/Scalabit_Challenge/Returnnpulls/services"
	"github.com/gin-gonic/gin"
)

func GetOpenPullRequests(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")
	limitStr := c.DefaultQuery("limit", "20")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	prs, err := services.FetchOpenPRs(owner, repo, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prs)
}
