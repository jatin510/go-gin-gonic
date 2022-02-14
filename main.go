package main

import "gin-gonic-api/server"

func main() {
	// r := gin.Default()
	// r.GET("/", controllers.Get)

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	server.Initialize()
}
