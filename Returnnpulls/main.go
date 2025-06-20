package main

import (
	"github.com/GuilhermeBatist/Scalabit_Challenge_2/Returnnpulls/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.GET("/repos/:owner/:repo/pulls", handlers.GetOpenPullRequests)

	route.Run(":8080")

}
