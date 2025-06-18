package main

func main() {

	route := gin.Default()

	route.GET("/repos/:owner/:repo/pulls", handler.GetOpenPullRequests)

	router.run(":8080")

}
